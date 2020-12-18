import { readFileSync } from "fs";

type treeDim = [number, number, number];
type fourDim = [number, number, number, number];

class Satellite {
  actives: treeDim[];

  constructor(input: string[]) {
    this.actives = new Array<treeDim>();
    let offset = -Math.floor(input.length / 2);
    for (let row = 0; row < input.length; row++) {
      for (let col = 0; col < input[row].length; col++) {
        if (input[row][col] === "#") {
          this.actives.push([row + offset, col + offset, 0]);
        }
      }
    }
    // console.log(`Before any cycles:\n${this.visualize()}`);
  }

  getNeighbours([x, y, z]: treeDim): treeDim[] {
    const neighbours = new Array<treeDim>();
    for (let i = -1; i <= 1; i++) {
      for (let j = -1; j <= 1; j++) {
        for (let k = -1; k <= 1; k++) {
          if (i !== 0 || j !== 0 || k !== 0) {
            neighbours.push([x + i, y + j, z + k]);
          }
        }
      }
    }
    return neighbours;
  }

  equal(left: treeDim, right: treeDim): boolean {
    return JSON.stringify(left) === JSON.stringify(right);
  }

  contains(arr: treeDim[], element: treeDim): boolean {
    let has = arr.reduce((has, el) => (has ||= this.equal(el, element)), false);
    return has;
  }

  getActives(actives: treeDim[], positions: treeDim[]): number {
    return positions.reduce(
      (active, pos) => (active += this.contains(actives, pos) ? 1 : 0),
      0
    );
  }

  simulate(): number {
    const cycles = 6;
    for (let i = 0; i < cycles; i++) {
      let newActives = new Array<treeDim>();
      const checkedPositions = new Array<treeDim>();
      const positionsToCheck = new Array<treeDim>();
      positionsToCheck.push([0, 0, 0]);
      while (positionsToCheck.length !== 0) {
        const pos = positionsToCheck.shift();
        const neighbours = this.getNeighbours(pos);
        const activeNeighbours = this.getActives(this.actives, neighbours);
        if (this.contains(this.actives, pos)) {
          if (activeNeighbours === 2 || activeNeighbours === 3) {
            newActives.push(pos);
          }
        } else {
          if (activeNeighbours === 3) {
            newActives.push(pos);
          }
        }
        checkedPositions.push(pos);
        neighbours.forEach((el) => {
          if (
            !this.contains(positionsToCheck, el) &&
            !this.contains(checkedPositions, el) &&
            this.getActives(this.actives, this.getNeighbours(el)) !== 0
          ) {
            positionsToCheck.push(el);
          }
        });
      }
      this.actives = newActives;
      // console.log(`after ${i + 1} cycles:\n\n${this.visualize()}`);
    }
    return this.actives.length;
  }

  visualize(): string {
    const [minX, maxX] = this.actives.reduce(
      ([min, max], el) => [Math.min(min, el[0]), Math.max(max, el[0])],
      [0, 0]
    );
    const [minY, maxY] = this.actives.reduce(
      ([min, max], el) => [Math.min(min, el[1]), Math.max(max, el[1])],
      [0, 0]
    );
    const [minZ, maxZ] = this.actives.reduce(
      ([min, max], el) => [Math.min(min, el[2]), Math.max(max, el[2])],
      [0, 0]
    );
    let str = `[${minX}, ${maxX}], [${minY}, ${maxY}], [${minZ}, ${maxX}]\n`;
    for (let i = minZ; i <= maxZ; i++) {
      str += `z=${i}\n`;
      for (let j = minX; j <= maxX; j++) {
        for (let k = minY; k <= maxY; k++) {
          str += this.contains(this.actives, [j, k, i]) ? "#" : ".";
        }
        str += "\n";
      }
      str += "\n";
    }
    return str;
  }
}

const input = readFileSync("./example.txt", "utf8").split("\n");

const s = new Satellite(input);
console.log(`Part 1: ${s.simulate()}`);
