import { readFileSync } from "fs";

function ithSpoken(numbers: number[], limit: number): number {
  const history = new Map<number, [number, number]>();
  let turn = 1;
  let mostRecentSpoken = 0;
  for (let i = 0; i < numbers.length; i++) {
    mostRecentSpoken = numbers[i];
    history.set(numbers[i], [turn, 0]);
    turn++;
  }

  while (turn <= limit) {
    let [last, prevLast] = history.get(mostRecentSpoken);
    if (prevLast === 0) {
      mostRecentSpoken = 0;
    } else {
      mostRecentSpoken = last - prevLast;
    }
    history.set(mostRecentSpoken, [
      turn,
      (history.get(mostRecentSpoken) || [0, 0])[0],
    ]);
    turn++;
  }
  return mostRecentSpoken;
}

const input = readFileSync("./input.txt", "utf8").split("\n");
const numbers = input[0].split(",").map((el) => +el);

console.log(`Part 1: ${ithSpoken(numbers, 2020)}`);
console.log(`Part 2: ${ithSpoken(numbers, 30000000)}`);
