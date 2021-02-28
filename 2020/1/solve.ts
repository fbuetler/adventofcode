import { readFileSync } from "fs";

function solvePartOne(input: number[]) {
  console.log("part 1");
  let left = 0;
  let right = input.length - 1;
  while (left < right) {
    const a = input[left];
    const b = input[right];
    if (a + b === 2020) {
      console.log(`${a} * ${b} = ${a * b}`);
      return;
    } else if (a + b > 2020) {
      right--;
    } else {
      left++;
    }
  }
}

function solvePartTwo(input: number[]) {
  console.log("part 2");
  let completed = false;
  for (let i = 0; i < input.length - 1; i++) {
    const a = input[i];
    let left = i + 1;
    let right = input.length - 1;
    while (left < right) {
      const b = input[left];
      const c = input[right];
      if (a + b + c === 2020) {
        completed = true;
        console.log(`${a} * ${b} * ${c} = ${a * b * c}`);
        break;
      } else if (a + b + c > 2020) {
        right--;
      } else {
        left++;
      }
    }
    if (completed) {
      return;
    }
  }
}

const data = readFileSync("./input.txt", "utf8");
const input = data
  .split("\n")
  .map((el) => +el)
  .sort((a, b) => a - b);

solvePartOne(input);
solvePartTwo(input);
