import { readFileSync } from "fs";

enum Direction {
  LEFT = "LEFT",
  RIGHT = "RIGHT",
  UP = "UP",
  DOWN = "DOWN",
}

enum Transformation {
  ORIGINAL = "ORIGINAL",
  VERTICAL = "VERTICAL",
  HORIZONTAL = "HORIZONTAL",
  ROTATION90 = "ROTATION90",
  ROTATION180 = "ROTATION180",
  ROTATION270 = "ROTATION270",
}

class Photo {
  tiles = new Map<number, Tile>();
  reconstructed = new Array<Array<number>>();
  size: number;

  constructor(rawTiles: string[][]) {
    rawTiles.forEach((tile) => {
      const title = tile.shift();
      this.tiles.set(+title.replace(":", "").split(" ")[1], new Tile(tile));
    });
    this.size = Math.sqrt(rawTiles.length);
    for (let i = 0; i < this.size; i++) {
      this.reconstructed.push(new Array<number>(this.size));
    }
  }

  reconstruct() {
    const neighbours = new Map<number, Set<number>>();
    const tileKeys = Array.from(this.tiles.keys()).sort();
    for (let tileKey of tileKeys) {
      for (let otherTileKey of tileKeys) {
        if (tileKey === otherTileKey) {
          continue;
        }
        const tile = this.tiles.get(tileKey);
        const otherTile = this.tiles.get(otherTileKey);
        const transformations = [
          Transformation.ORIGINAL,
          Transformation.HORIZONTAL,
          Transformation.VERTICAL,
          Transformation.ROTATION90,
          Transformation.ROTATION180,
          Transformation.ROTATION270,
        ];
        [Direction.LEFT, Direction.RIGHT, Direction.UP, Direction.DOWN].forEach(
          (direction: Direction) => {
            for (let i = 0; i < transformations.length; i++) {
              for (let j = 0; j < transformations.length; j++) {
                if (
                  Tile.fits(
                    tile,
                    otherTile,
                    direction,
                    transformations[i],
                    transformations[j]
                  )
                ) {
                  neighbours.set(
                    tileKey,
                    (neighbours.get(tileKey) || new Set<number>()).add(
                      otherTileKey
                    )
                  );
                }
              }
            }
          }
        );
      }
    }
    console.log(neighbours);
    let randomCorner: number;
    let cornerProduct = 1;
    neighbours.forEach((nbrs, tileKey) => {
      cornerProduct *= nbrs.size === 2 ? tileKey : 1;
      randomCorner = tileKey;
    });
    console.log(`Part 1: ${cornerProduct}`);
  }
}

class Tile {
  tile: string[];

  constructor(tile: string[]) {
    this.tile = tile;
  }

  rotate90() {
    let newTile = new Array<string>(this.tile.length);
    for (let i = 0; i < this.tile.length; i++) {
      let row = "";
      for (let j = 0; j < this.tile.length; j++) {
        row = row.concat(this.tile[this.tile.length - 1 - j][i]);
      }
      newTile[i] = row;
    }
    this.tile = newTile;
  }

  flipVertically() {
    let newTile = new Array<string>(this.tile.length);
    for (let i = 0; i < this.tile.length; i++) {
      newTile[i] = this.tile[i].split("").reverse().join("");
    }
    this.tile = newTile;
  }

  flipHorrizontally() {
    let newTile = new Array<string>(this.tile.length);
    for (let i = 0; i < this.tile.length; i++) {
      newTile[this.tile.length - 1 - i] = this.tile[i];
    }
    this.tile = newTile;
  }

  static fits(
    one: Tile,
    other: Tile,
    direction: Direction,
    tf: Transformation,
    otherTf: Transformation
  ): boolean {
    const copy = one.copy();
    const otherCopy = other.copy();
    const config: [Tile, Transformation][] = [
      [copy, tf],
      [otherCopy, otherTf],
    ];

    for (let [t, transformation] of config) {
      switch (transformation) {
        case Transformation.ORIGINAL: {
          break;
        }
        case Transformation.VERTICAL: {
          t.flipVertically();
          break;
        }
        case Transformation.HORIZONTAL: {
          t.flipHorrizontally();
          break;
        }
        case Transformation.ROTATION90: {
          t.rotate90();
          break;
        }
        case Transformation.ROTATION180: {
          t.rotate90();
          t.rotate90();
          break;
        }
        case Transformation.ROTATION270: {
          t.rotate90();
          t.rotate90();
          t.rotate90();
          break;
        }
      }
    }

    if (copy.checkTransformation(otherCopy, direction)) {
      one = copy;
      other = otherCopy;
      return true;
    }
    return false;
  }

  checkTransformation(other: Tile, direction: Direction): boolean {
    switch (direction) {
      case Direction.LEFT: {
        for (let i = 0; i < this.tile.length; i++) {
          if (this.tile[i].substr(0, 1) !== other.tile[i].substr(-1, 1)) {
            return false;
          }
        }
        return true;
      }
      case Direction.RIGHT: {
        for (let i = 0; i < this.tile.length; i++) {
          if (this.tile[i].substr(-1, 1) !== other.tile[i].substr(0, 1)) {
            return false;
          }
        }
        return true;
      }
      case Direction.UP: {
        return this.tile[0] === other.tile[0];
      }
      case Direction.DOWN: {
        const len = this.tile.length - 1;
        return this.tile[len] === other.tile[len];
      }
      default:
        return false;
    }
  }

  copy(): Tile {
    return new Tile(this.tile.slice());
  }

  toString(): string {
    let str = "";
    this.tile.forEach((line) => (str += `${line}\n`));
    return str;
  }
}

const input = readFileSync("./example.txt", "utf8").split("\n\n");
const tiles = new Array<Array<string>>();
input.forEach((el) => tiles.push(el.split("\n")));

const photo = new Photo(tiles);
photo.reconstruct();
