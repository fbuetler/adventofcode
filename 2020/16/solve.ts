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

const possibleRanges = new Map<string, Set<number>>();
fields.forEach((val, key) =>
  possibleRanges.set(
    key,
    new Set<number>(
      Array(validTickets[0].length)
        .fill(0)
        .map((_, i) => i)
    )
  )
);
for (let ticket of validTickets) {
  let index = 0;
  for (let field of ticket) {
    for (let [name, range] of Array.from(fields)) {
      if (
        !isValidField(field, range[0][0], range[0][1], range[1][0], range[1][1])
      ) {
        const pr = possibleRanges.get(name);
        pr.delete(index);
        possibleRanges.set(name, pr);
      }
    }
    index++;
  }
}

const assigned = new Map<string, number>();
for (let i = 0; i < possibleRanges.size; i++) {
  let onePossibility: string;
  possibleRanges.forEach((possibilities, field) => {
    if (possibilities.size === 1) {
      onePossibility = field;
    }
  });
  const category = Array.from(possibleRanges.get(onePossibility))[0];
  assigned.set(onePossibility, category);
  possibleRanges.forEach((possibilities) => {
    possibilities.delete(category);
  });
}

let product = 1;
assigned.forEach((index, el) => {
  if (el.startsWith("departure")) {
    product *= myTicket[index];
  }
});
console.log(`Part 2: ${product}`);
