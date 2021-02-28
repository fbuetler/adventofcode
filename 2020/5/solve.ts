import { readFileSync } from "fs";

function find(elems: boolean[]): number {
  let counter = 0;
  for (let i = 0; i < elems.length; i++) {
    if (elems[i]) {
      counter += Math.floor(Math.pow(2, elems.length - i - 1));
    }
  }
  return counter;
}

// true means upper half
function findSeat(pass: boolean[]): number {
  const row = find(pass.slice(0, 7));
  const col = find(pass.slice(7, 10));
  return row * 8 + col;
}

const input = readFileSync("./input.txt", "utf8").split("\n");

let maxID = 0;
let checkedSeats = new Set<number>();
for (let i = 0; i < input.length; i++) {
  const id = findSeat(input[i].split("").map((el) => el === "B" || el === "R"));
  maxID = Math.max(maxID, id);
  checkedSeats.add(id);
}
console.log(`highest seat ID: ${maxID}`);

let emptySeats: Array<number> = [];
for (let i = 0; i < maxID; i++) {
  if (!checkedSeats.has(i)) {
    emptySeats.push(i);
  }
}
console.log(
  `my seat ID: ${emptySeats.filter(
    (el) => !emptySeats.includes(el - 1) && !emptySeats.includes(el + 1)
  )}`
);
