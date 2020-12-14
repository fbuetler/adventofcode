import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");

const reMask = new RegExp(`^mask\\s=\\s((X|\\d){36}?)$`);
const reMem = new RegExp(`^mem\\[(\\d+?)\\]\\s=\\s(\\d+?)$`);

function toBitArray(d: number): string[] {
  let b = new Array<string>();
  while (d > 0) {
    if (d) {
      b.push(String(d % 2));
      d = Math.floor(d / 2);
    }
  }
  for (let i = b.length; i < 36; i++) {
    b.push("0");
  }
  return b.reverse();
}

function toDecimal(b: string[]) {
  let d = 0;
  for (let i = 0; i < b.length; i++) {
    if (b[i] === "1") {
      d += Math.pow(2, b.length - 1 - i);
    }
  }
  return d;
}

function resolveAddr(addr: string[]): string[][] {
  let addrs = new Array<Array<string>>();
  addrs.push(new Array<string>());
  for (let i = 0; i < addr.length; i++) {
    if (addr[i] === "0" || addr[i] === "1") {
      addrs.forEach((a) => a.push(addr[i]));
    } else {
      let dups = new Array<Array<string>>();
      addrs.forEach((a) => {
        let dup = JSON.parse(JSON.stringify(a));
        dup.push("1");
        a.push("0");
        dups.push(dup);
      });
      addrs.push(...dups);
    }
  }
  return addrs;
}

function applyMaskPart1(mask: string[], value: string[]): string[] {
  for (let i = 1; i <= value.length; i++) {
    if (mask[mask.length - i] === "X") {
      continue;
    }
    value[value.length - i] = mask[mask.length - i];
  }
  return value;
}

function applyMaskPart2(mask: string[], value: string[]): string[] {
  for (let i = 1; i <= value.length; i++) {
    let idx = mask.length - i;
    if (mask[idx] === "0") {
      continue;
    } else {
      value[value.length - i] = mask[idx];
    }
  }
  return value;
}

let mask: string[];
let memoryPart1 = new Map<number, number>();
let memoryPart2 = new Map<number, number>();
for (let i = 0; i < input.length; i++) {
  const maskMatch = input[i].match(reMask);
  const memMatch = input[i].match(reMem);
  if (maskMatch !== null) {
    mask = maskMatch[1].split("");
  } else if (memMatch !== null) {
    let addr = +memMatch[1];
    let val = +memMatch[2];

    memoryPart1.set(addr, toDecimal(applyMaskPart1(mask, toBitArray(val))));

    resolveAddr(applyMaskPart2(mask, toBitArray(addr))).forEach((a) =>
      memoryPart2.set(toDecimal(a), val)
    );
  } else {
    console.error(`ERROR: ${input[i]} @ L${i}`);
  }
}

let sum = 0;
memoryPart1.forEach((val) => (sum += val));
console.log(`Part 1: ${sum}`);

sum = 0;
memoryPart2.forEach((val) => (sum += val));
console.log(`Part 2: ${sum}`);
