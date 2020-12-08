import { readFileSync } from "fs";

enum Instr {
  ACC,
  JMP,
  NOP,
}

function parseInput(input: string[]): [Instr, number][] {
  let instrs = new Array<[Instr, number]>();
  for (let i = 0; i < input.length; i++) {
    const parts = input[i].split(" ");
    const arg = +parts[1];
    switch (parts[0]) {
      case "acc": {
        instrs.push([Instr.ACC, arg]);
        break;
      }
      case "jmp": {
        instrs.push([Instr.JMP, arg]);
        break;
      }
      case "nop": {
        instrs.push([Instr.NOP, arg]);
        break;
      }
    }
  }
  return instrs;
}

function terminates(instrs: [Instr, number][]): [boolean, number] {
  let visited = new Array<boolean>(instrs.length);
  let i = 0;
  let acc = 0;
  while (true) {
    if (visited[i]) {
      break;
    }
    visited[i] = true;
    switch (instrs[i][0]) {
      case Instr.ACC: {
        acc += instrs[i][1];
        i++;
        break;
      }
      case Instr.JMP: {
        i += instrs[i][1];
        break;
      }
      case Instr.NOP: {
        i++;
        break;
      }
    }
    if (i >= instrs.length) {
      return [true, acc];
    }
  }
  return [false, acc];
}

const input = readFileSync("./input.txt", "utf8").split("\n");
const instrs = parseInput(input);

let [t, acc] = terminates(instrs);
console.log(`Part 1 - acc: ${acc}`);

let breakLoop = false;
for (let i = 0; i < instrs.length && !breakLoop; i++) {
  switch (instrs[i][0]) {
    case Instr.ACC: {
      break;
    }
    case Instr.JMP: {
      instrs[i][0] = Instr.NOP;
      [t, acc] = terminates(instrs);
      if (t) {
        breakLoop = true;
      }
      instrs[i][0] = Instr.JMP;
      break;
    }
    case Instr.NOP: {
      instrs[i][0] = Instr.JMP;
      [t, acc] = terminates(instrs);
      if (t) {
        breakLoop = true;
      }
      instrs[i][0] = Instr.NOP;
      break;
    }
  }
}
console.log(`Part 2 - acc: ${acc}`);
