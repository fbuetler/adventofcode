import { count } from "console";
import { readFileSync } from "fs";

function parseInput(input: string[]): Map<string, Map<string, number>> {
  let outterBags = new Map<string, Map<string, number>>();
  for (let i = 0; i < input.length; i++) {
    let words = input[i].split(" ");
    let innerBags = new Map<string, number>();
    if (words[4] !== "no") {
      let start = 0;
      do {
        start += 4;
        innerBags.set(`${words[start + 1]} ${words[start + 2]}`, +words[start]);
      } while (words[start + 3].endsWith(","));
    }
    outterBags.set(`${words[0]} ${words[1]}`, innerBags);
  }
  return outterBags;
}

function dfs(
  g: Map<string, Map<string, number>>,
  visited: Map<string, [boolean, number]>,
  bag: string
) {
  if (visited.has(bag)) {
    return;
  }

  if (bag === "shiny gold") {
    visited.set(bag, [true, 0]);
    return;
  }

  let s = new Array<string>();
  g.get(bag).forEach((n, b) => {
    for (let i = 0; i < n; i++) {
      s.push(b);
    }
  });

  let c = false;
  let size = 0;
  while (s.length !== 0) {
    let b = s.pop();
    dfs(g, visited, b);
    const [tmpC, tmpSize] = visited.get(b);
    c ||= tmpC;
    size += tmpSize;
  }

  g.get(bag).forEach((val) => (size += val));

  visited.set(bag, [c, size]);
}

const input = readFileSync("./input.txt", "utf8").split("\n");

let bags = parseInput(input);

let colorCounerPartOne = 0;
let colorCounerPartTwo = 0;
let visited = new Map<string, [boolean, number]>();
for (const bag of Array.from(bags.keys())) {
  if (bag === "shiny gold") {
    continue;
  }
  dfs(bags, visited, bag);
  if (visited.get(bag)[0]) {
    colorCounerPartOne++;
  }
}

bags.get("shiny gold").forEach((n, bag) => {
  colorCounerPartTwo += n + n * visited.get(bag)[1];
});

console.log(`Part 1: ${colorCounerPartOne}`);
console.log(`Part 2: ${colorCounerPartTwo}`);
