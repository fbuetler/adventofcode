import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8")
  .split("\n")
  .map((el) => +el)
  .sort((a, b) => a - b);

const builtInJ = input[input.length - 1] + 3;
input.push(builtInJ);

let counterPerJ = new Map<number, number>();
counterPerJ.set(1, 0);
counterPerJ.set(3, 0);
let countPerAdapter = new Map<number, number>();
countPerAdapter.set(0, 1);

for (let i = 1; i < input.length; i++) {
  let sum = 0;
  let diffFound = false;
  for (let j = 1; j <= 3; j++) {
    if (input.includes(input[i] - j)) {
      sum += countPerAdapter.get(input[i] - j);
    }
    if (!diffFound) {
      const diff = input[i] - input[i - 1];
      counterPerJ.set(diff, counterPerJ.get(diff) + 1);
      diffFound = true;
    }
  }
  countPerAdapter.set(input[i], sum);
}

console.log(
  `Part1: ${counterPerJ.get(1)} * ${counterPerJ.get(3)} = ${
    counterPerJ.get(1) * counterPerJ.get(3)
  }`
);
console.log(`Part2: ${countPerAdapter.get(builtInJ)}`);
