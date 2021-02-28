import { readFileSync } from "fs";

enum Direction {
  NORTHEAST = "NE", // 1, -1
  NORTHWEST = "NW", // -1, -1
  SOUTHEAST = "SE", // 1, 1
  SOUTHWEST = "SW", // -1, 1
  EAST = "E", // 2, 0
  WEST = "W", // -2, 0
}

function parse(input: string[]): Direction[][] {
  const instrs = new Array<Array<Direction>>();
  input.forEach((line) => {
    const instr = new Array<Direction>();
    let i = 0;
    while (i < line.length) {
      if (line.charAt(i) === "n" || line.charAt(i) === "s") {
        switch (line.substr(i, 2)) {
          case "ne": {
            instr.push(Direction.NORTHEAST);
            break;
          }
          case "nw": {
            instr.push(Direction.NORTHWEST);
            break;
          }
          case "se": {
            instr.push(Direction.SOUTHEAST);
            break;
          }
          case "sw": {
            instr.push(Direction.SOUTHWEST);
            break;
          }
        }
        i += 2;
      } else {
        switch (line.substr(i, 1)) {
          case "e": {
            instr.push(Direction.EAST);
            break;
          }
          case "w": {
            instr.push(Direction.WEST);
            break;
          }
        }
        i++;
      }
    }
    instrs.push(instr);
  });
  return instrs;
}

function getNeighbours([x, y]: [number, number]): Array<[number, number]> {
  const neighbours = new Array<[number, number]>();
  [
    [1, -1],
    [-1, -1],
    [1, 1],
    [-1, 1],
    [2, 0],
    [-2, 0],
  ].forEach(([u, v]) => {
    neighbours.push([x + u, y + v]);
  });
  return neighbours;
}

const input = readFileSync("./input.txt", "utf8").split("\n");
const instrs = parse(input);

let blackTiles = new Set<string>();
instrs.forEach((instr) => {
  let x = 0;
  let y = 0;
  instr.forEach((direction) => {
    switch (direction) {
      case Direction.NORTHEAST: {
        x++;
        y--;
        break;
      }
      case Direction.NORTHWEST: {
        x--;
        y--;
        break;
      }
      case Direction.SOUTHEAST: {
        x++;
        y++;
        break;
      }
      case Direction.SOUTHWEST: {
        x--;
        y++;
        break;
      }
      case Direction.EAST: {
        x += 2;
        break;
      }
      case Direction.WEST: {
        x -= 2;
        break;
      }
    }
  });
  if (blackTiles.has(`${x},${y}`)) {
    blackTiles.delete(`${x},${y}`);
  } else {
    blackTiles.add(`${x},${y}`);
  }
});
console.log(`Part 1: ${blackTiles.size}`);

let day = 0;
const maxDay = 100;
while (day < maxDay) {
  const newBlackTiles = new Set<string>();
  for (let el of Array.from(blackTiles.values())) {
    const [x, y] = el.split(",").map(Number);

    const adjacentWhiteTiles = new Array<[number, number]>();
    let blackNeighbours = 0;
    getNeighbours([x, y]).forEach(([u, v]) => {
      if (blackTiles.has(`${u},${v}`)) {
        blackNeighbours++;
      } else {
        adjacentWhiteTiles.push([u, v]);
      }
    });
    if (blackNeighbours === 1 || blackNeighbours === 2) {
      newBlackTiles.add(`${x},${y}`);
    }

    adjacentWhiteTiles.forEach(([u, v]) => {
      blackNeighbours = 0;
      getNeighbours([u, v]).forEach(
        ([a, b]) => (blackNeighbours += blackTiles.has(`${a},${b}`) ? 1 : 0)
      );
      if (blackNeighbours === 2) {
        newBlackTiles.add(`${u},${v}`);
      }
    });
  }
  blackTiles = newBlackTiles;
  day++;
}
console.log(`Part 2: ${blackTiles.size}`);
