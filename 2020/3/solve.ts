import { readFileSync } from "fs";

class Forest {
  lines: string[];
  position: [number, number];
  encounteredTrees: number;
  goRight: number;
  goDown: number;

  constructor(input: string[], goRight: number, goDown: number) {
    this.lines = input;
    this.goRight = goRight;
    this.goDown = goDown;
    this.position = [0, 0];
    this.encounteredTrees = 0;
  }

  isTree(): boolean {
    return this.lines[this.position[0]].charAt(this.position[1]) === "#";
  }

  isTraversed(): boolean {
    return this.position[0] >= this.lines.length;
  }

  doStep() {
    this.position = [
      this.position[0] + this.goDown,
      (this.position[1] + this.goRight) % this.lines[0].length, // all lines are equally long
    ];
  }

  traverse(): number {
    while (!this.isTraversed()) {
      if (this.isTree()) {
        this.encounteredTrees++;
      }
      this.doStep();
    }
    return this.encounteredTrees;
  }
}

const input = readFileSync("./input.txt", "utf8").split("\n");
const slopes = [
  [1, 1],
  [3, 1],
  [5, 1],
  [7, 1],
  [1, 2],
];
let result = 1;
for (let slope of slopes) {
  const f = new Forest(input, slope[0], slope[1]);
  result *= f.traverse();
}
console.log(`product of encountered trees ${result}`);
