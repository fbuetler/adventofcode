import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");

let countPartOne = 0;
let countPartTwo = 0;
let i = 0;
while (i < input.length) {
  let groupAnswers = new Map<string, number>();
  let groupSize = 0;
  while (input[i] != "") {
    input[i]
      .split("")
      .forEach((el) => groupAnswers.set(el, (groupAnswers.get(el) || 0) + 1));
    groupSize++;
    i++;
  }
  countPartOne += groupAnswers.size;
  groupAnswers.forEach((value) => {
    if (value === groupSize) {
      countPartTwo++;
    }
  });
  i++;
}

console.log(
  `Part 1 - number of questions ANYONE answered yes: ${countPartOne}`
);
console.log(
  `Part 2 - number of questions EVERYONE answered yes: ${countPartTwo}`
);
