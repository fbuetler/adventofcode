import { readFileSync } from "fs";

class Photo {
  tiles = new Map<number, Tile>();
  reconstructed = new Map<string, Tile>();
  size: number;
  edgeCount = new Map<string, number>();
  photo: Tile;
  monster = [
    "..................#.",
    "#....##....##....###",
    ".#..#..#..#..#..#...",
  ];

  constructor(rawTiles: string[][]) {
    rawTiles.forEach((tile) => {
      const title = tile.shift();
      const id = +title.replace(":", "").split(" ")[1];
      this.tiles.set(id, new Tile(tile));
      // count edge appearances because border edge only appear once
      this.tiles
        .get(id)
        .edges()
        .forEach((edge) => {
          this.edgeCount.set(edge, (this.edgeCount.get(edge) || 0) + 1);
          this.edgeCount.set(
            edge.split("").reverse().join(""),
            (this.edgeCount.get(edge.split("").reverse().join("")) || 0) + 1
          );
        });
    });
    this.size = Math.sqrt(rawTiles.length);
    this.photo = new Tile(
      new Array<string>(this.size * (rawTiles[0].length - 2)).fill("")
    );
  }

  reconstruct() {
    let cornerProduct = 1;
    for (let row = 0; row < this.size; row++) {
      for (let col = 0; col < this.size; col++) {
        let found = false;
        for (let id of Array.from(this.tiles.keys())) {
          const tile = this.tiles.get(id);
          for (let transformation of Tile.transformations(tile)) {
            const topEdge = transformation.rows[0];
            const leftEdge = transformation.col(0);
            /*
                is top row OR top edge fits to upper tile AND
                is leftmost col OR left edge fits to left tile
              */
            if (
              ((row === 0 && this.edgeCount.get(topEdge) === 1) ||
                (row !== 0 &&
                  topEdge ===
                    this.reconstructed.get(`${col},${row - 1}`).rows[
                      tile.size - 1
                    ])) &&
              ((col === 0 && this.edgeCount.get(leftEdge) === 1) ||
                (col !== 0 &&
                  leftEdge ===
                    this.reconstructed
                      .get(`${col - 1},${row}`)
                      .col(tile.size - 1)))
            ) {
              // is corner
              if (
                (row === 0 || row === this.size - 1) &&
                (col === 0 || col === this.size - 1)
              ) {
                cornerProduct *= id;
              }

              // reconstruct the photo
              for (let i = 0; i < tile.size - 2; i++) {
                this.photo.rows[
                  (tile.size - 2) * row + i
                ] += transformation.rows[i + 1].substring(1, tile.size - 1);
              }

              this.reconstructed.set(`${col},${row}`, transformation);
              this.tiles.delete(id);
              found = true;
              break;
            }
          }
          if (found) {
            break;
          }
        }
      }
    }
    console.log(`Part 1: ${cornerProduct}`);
  }

  waterRoughness(): number {
    let monsters = 0;
    for (let transformation of Tile.transformations(this.photo)) {
      for (let row = 0; row < transformation.rows.length; row++) {
        for (let col = 0; col < transformation.rows[0].length; col++) {
          if (this.partContainsMonster(transformation, row, col)) {
            monsters++;
          }
        }
      }
    }
    let hashesInPhoto = 0;
    this.photo.rows.forEach(
      (row) => (hashesInPhoto += (row.match(new RegExp("#", "g")) || []).length)
    );
    let hashesInMonster = 0;
    this.monster.forEach(
      (row) =>
        (hashesInMonster += (row.match(new RegExp("#", "g")) || []).length)
    );
    return hashesInPhoto - monsters * hashesInMonster;
  }

  partContainsMonster(tile: Tile, row: number, col: number): boolean {
    if (
      row + this.monster.length > tile.rows.length ||
      col + this.monster[0].length > tile.rows[0].length
    ) {
      return false;
    }
    for (let i = 0; i < this.monster.length; i++) {
      for (let j = 0; j < this.monster[i].length; j++) {
        if (
          this.monster[i][j] === "#" &&
          tile.rows[row + i][col + j] !== this.monster[i][j]
        ) {
          return false;
        }
      }
    }
    return true;
  }

  toString(): string {
    return this.photo.toString();
  }
}

class Tile {
  rows: string[];
  size: number;

  constructor(rows: string[]) {
    this.rows = rows;
    this.size = rows.length;
  }

  col(index: number): string {
    let col = "";
    this.rows.forEach((row) => (col += row[index]));
    return col;
  }

  edges(): string[] {
    return [
      this.rows[0],
      this.col(this.size - 1),
      this.rows[this.size - 1],
      this.col(0),
    ];
  }

  static transformations(t: Tile): Tile[] {
    const size = t.size;
    const trans = new Array<Tile>();
    for (let i = 0; i < 8; i += 2) {
      let newTile = new Array<string>();
      // flip vertically
      for (let j = 0; j < size; j++) {
        newTile.push(t.rows[j].split("").reverse().join(""));
      }
      trans.push(new Tile(newTile));
      newTile = new Array<string>();
      // rotate
      for (let j = 0; j < size; j++) {
        newTile.push(t.col(j).split("").reverse().join(""));
      }
      trans.push(new Tile(newTile));
      // update
      t = trans[i + 1];
    }
    return trans;
  }

  toString(): string {
    let str = "";
    for (let i = 0; i < this.rows.length; i++) {
      str += `${this.rows[i]}\n`;
    }
    return str;
  }
}

const input = readFileSync("./example.txt", "utf8").split("\n\n");
const tiles = new Array<Array<string>>();
input.forEach((el) => tiles.push(el.split("\n")));

const photo = new Photo(tiles);
photo.reconstruct();
console.log(`Part 2: ${photo.waterRoughness()}`);
