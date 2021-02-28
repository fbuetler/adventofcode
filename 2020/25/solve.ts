import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n").map(Number);

const subjectNumber = 7;
const divider = 20201227;
let loopSizes = new Array<number>(2);

input.forEach((publicKey, index) => {
  let value = 1;
  let loopSize = 0;
  while (value !== publicKey) {
    value *= subjectNumber;
    value %= divider;
    loopSize++;
  }
  console.log(`loopsize: ${loopSize}`);
  loopSizes[index] = loopSize;
});

input.forEach((publicKey, index) => {
  let value = 1;
  const loopSize = loopSizes[(index + 1) % loopSizes.length];
  for (let i = 0; i < loopSize; i++) {
    value *= publicKey;
    value %= divider;
  }
  console.log(`encryption key: ${value}`);
});
