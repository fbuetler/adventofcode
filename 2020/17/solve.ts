import { readFileSync } from "fs";

type dim = [number, number, number, number];

class Satellite {
  actives: Set<dim>;
  rangeXYZ = [-1, 0, 1];
  rangeW = [0];

  constructor(input: string[], fourDim: boolean) {
    if (fourDim) {
      this.rangeW = [-1, 0, 1];
    }
    this.actives = new Set<dim>();
    for (let row = 0; row < input.length; row++) {
      for (let col = 0; col < input[row].length; col++) {
        if (input[row][col] === "#") {
          this.actives.add([row, col, 0, 0]);
        }
      }
    }
    console.log(`Before any cycles:\n${this.visualize()}\n`);
  }

  has(set: Set<dim>, elem: dim): boolean {
    for (let el of Array.from(set)) {
      if (JSON.stringify(el) === JSON.stringify(elem)) {
        return true;
      }
    }
    return false;
  }

  add(set: Set<dim>, elem: dim) {
    if (!this.has(set, elem)) {
      set.add(elem);
    }
  }

  addNeighboursOf(set: Set<dim>, [x, y, z, w]: dim) {
    for (let i of this.rangeXYZ) {
      for (let j of this.rangeXYZ) {
        for (let k of this.rangeXYZ) {
          for (let l of this.rangeW) {
            if (i !== 0 || j !== 0 || k !== 0 || l !== 0) {
              this.add(set, [x + i, y + j, z + k, w + l]);
            }
          }
        }
      }
    }
  }

  countActiveNeighbours([x, y, z, w]: dim): number {
    let activeNeighbours = 0;
    for (let i of this.rangeXYZ) {
      for (let j of this.rangeXYZ) {
        for (let k of this.rangeXYZ) {
          for (let l of this.rangeW) {
            if (i !== 0 || j !== 0 || k !== 0 || l !== 0) {
              if (this.has(this.actives, [x + i, y + j, z + k, w + l])) {
                activeNeighbours++;
              }
            }
          }
        }
      }
    }
    return activeNeighbours;
  }

  simulate(): number {
    const cycles = 6;
    for (let i = 0; i < cycles; i++) {
      console.log(`Simulating cycle: ${i + 1}`);

      let positionsToCheck = new Set<dim>();
      this.actives.forEach((active) =>
        this.addNeighboursOf(positionsToCheck, active)
      );

      console.log(`Positions to check: ${positionsToCheck.size}`);
      let newActives = new Set<dim>();
      positionsToCheck.forEach((pos) => {
        const activeNeighbours = this.countActiveNeighbours(pos);
        // console.log(`pos: ${pos}, actives: ${activeNeighbours}`);
        if (this.has(this.actives, pos)) {
          if (activeNeighbours === 2 || activeNeighbours === 3) {
            this.add(newActives, pos);
          }
        } else {
          if (activeNeighbours === 3) {
            this.add(newActives, pos);
          }
        }
      });
      this.actives = newActives;
      // console.log(`After ${i + 1} cycles:\n\n${this.visualize()}`);
      console.log(`Active cubes: ${this.actives.size}`);
    }
    return this.actives.size;
  }

  visualize(): string {
    const [minX, maxX] = Array.from(this.actives).reduce(
      ([min, max], el) => [Math.min(min, el[0]), Math.max(max, el[0])],
      [0, 0]
    );
    const [minY, maxY] = Array.from(this.actives).reduce(
      ([min, max], el) => [Math.min(min, el[1]), Math.max(max, el[1])],
      [0, 0]
    );
    const [minZ, maxZ] = Array.from(this.actives).reduce(
      ([min, max], el) => [Math.min(min, el[2]), Math.max(max, el[2])],
      [0, 0]
    );
    const [minW, maxW] = Array.from(this.actives).reduce(
      ([min, max], el) => [Math.min(min, el[3]), Math.max(max, el[3])],
      [0, 0]
    );
    let str = `[${minX}, ${maxX}], [${minY}, ${maxY}], [${minZ}, ${maxX}], [${minW}, ${maxW}]\n`;
    for (let l = minW; l <= maxW; l++) {
      for (let i = minZ; i <= maxZ; i++) {
        str += `z=${i}, w=${l}\n`;
        for (let j = minX; j <= maxX; j++) {
          for (let k = minY; k <= maxY; k++) {
            str += this.has(this.actives, [j, k, i, l]) ? "#" : ".";
          }
          str += "\n";
        }
        str += "\n";
      }
    }
    return str;
  }
}

const input = readFileSync("./input.txt", "utf8").split("\n");
console.log(`Part 1: ${new Satellite(input, false).simulate()}`);
console.log(`Part 2: ${new Satellite(input, true).simulate()}`);
