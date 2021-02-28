import { readFileSync } from "fs";

function sumExists(sum: number, values: number[]): boolean {
  values.sort((a, b) => a - b);
  let left = 0;
  let right = values.length - 1;
  while (left < right) {
    const a = values[left];
    const b = values[right];
    if (a + b === sum) {
      return true;
    } else if (a + b > sum) {
      right--;
    } else {
      left++;
    }
  }
  return false;
}

const input = readFileSync("./input.txt", "utf8")
  .split("\n")
  .map((el) => +el);

const preambleLength = 25;
let invalidNumber: number;

for (let i = preambleLength; i < input.length; i++) {
  if (!sumExists(input[i], input.slice(i - preambleLength, i))) {
    console.log(`Part 1 - Sum does not exist for: ${input[i]}`);
    invalidNumber = input[i];
    break;
  }
}

let left = 0;
let right = 2;
while (left < right && right < input.length) {
  if (right - left < 2) {
    right++;
    continue;
  }
  const range = input.slice(left, right);
  const sum = range.reduce((acc, val) => acc + val, 0);
  if (sum === invalidNumber) {
    const min = range.reduce((min, val) => Math.min(min, val));
    const max = range.reduce((max, val) => Math.max(max, val));
    console.log(
      `Part 2 - Encryption weakness found: sum(${input.slice(
        left,
        right
      )}) = ${invalidNumber} => ${min} + ${max} = ${min + max}`
    );
    break;
  } else if (sum > invalidNumber) {
    left++;
  } else {
    right++;
  }
}
