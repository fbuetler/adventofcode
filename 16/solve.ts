import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");
const reField = new RegExp(
  `^((\\w|\\s)+?):\\s(\\d+?)-(\\d+?)\\sor\\s(\\d+?)-(\\d+?)$`
);

let fields = new Map<string, [[number, number], [number, number]]>();
let i = 0;
while (input[i] !== "") {
  const match = input[i].match(reField);
  fields.set(match[1], [
    [+match[3], +match[4]],
    [+match[5], +match[6]],
  ]);
  i++;
}
i += 2;
const myTicket = input[i].split(",").map((el) => +el);
i += 3;
const nearbyTickets = new Array<Array<number>>();
while (i < input.length) {
  nearbyTickets.push(input[i].split(",").map((el) => +el));
  i++;
}

function isValidField(
  field: number,
  a: number,
  b: number,
  c: number,
  d: number
): boolean {
  return (a <= field && field <= b) || (c <= field && field <= d);
}

let errorRate = 0;
const validTickets = new Array<Array<number>>();
for (let i = 0; i < nearbyTickets.length; i++) {
  let validTicket = true;
  for (let j = 0; j < nearbyTickets[i].length; j++) {
    let valid = false;
    fields.forEach((val) => {
      valid ||= isValidField(
        nearbyTickets[i][j],
        val[0][0],
        val[0][1],
        val[1][0],
        val[1][1]
      );
    });
    if (!valid) {
      errorRate += nearbyTickets[i][j];
      validTicket = false;
    }
  }
  if (validTicket) {
    validTickets.push(nearbyTickets[i]);
  }
}
console.log(`Part 1: ${errorRate}`);

