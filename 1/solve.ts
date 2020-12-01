import { readFileSync } from "fs";

function solvePartOne(input: string[]) {
  console.log("part 1");
  for (let i = 0; i < input.length; i++) {
    for (let j = 0; j < input.length; j++) {
      if (+input[i] + +input[j] === 2020) {
        console.log(`${input[i]} * ${input[j]} = ${+input[i] * +input[j]}`);
        return;
      }
    }
  }
}

function solvePartTwo(input: string[]) {
  console.log("part 2");
  for (let i = 0; i < input.length; i++) {
    for (let j = 0; j < input.length; j++) {
      for (let k = 0; k < input.length; k++) {
        if (+input[i] + +input[j] + +input[k] === 2020) {
          console.log(
            `${+input[i]} * ${+input[j]} * ${+input[k]} = ${
              +input[i] * +input[j] * +input[k]
            }`
          );
          return;
        }
      }
    }
  }
  return;
}

const data = readFileSync("./input.txt", "utf8");
const input = data.split("\n");

solvePartOne(input);
solvePartTwo(input);
