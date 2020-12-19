import { readFileSync } from "fs";

enum Type {
  SEQ = "seq",
  OR = "or",
  CHAR = "char",
}

type Rules = [Type, string | [number[], number[]] | number[]];

function parseRules(rules: string[]): Map<number, Rules> {
  let parsed = new Map<number, Rules>();
  for (let i = 0; i < rules.length; i++) {
    let parts = rules[i].split(":");
    let key = +parts[0];
    let value = parts[1].trim();
    if (value.includes("|")) {
      let options = [new Array<number>(), new Array<number>()];
      let ops = value.split("|");
      options[0].push(
        ...ops[0]
          .trim()
          .split(" ")
          .map((num) => +num)
      );
      options[1].push(
        ...ops[1]
          .trim()
          .split(" ")
          .map((num) => +num)
      );
      parsed.set(key, [Type.OR, options as [number[], number[]]]);
    } else if (value.includes('"')) {
      parsed.set(key, [Type.CHAR, value.trim().split('"').join("")]);
    } else {
      parsed.set(key, [
        Type.SEQ,
        value
          .trim()
          .split(" ")
          .map((el) => +el),
      ]);
    }
  }
  return parsed;
}

function generateMessages(rules: Map<number, Rules>, rule: number): string[] {
  const type = rules.get(rule);
  let messages = new Array<string>();
  switch (type[0]) {
    case Type.SEQ: {
      let msgs = [""];
      for (let i = 0; i < type[1].length; i++) {
        const ends = generateMessages(rules, type[1][i] as number);
        let tmpMsgs = new Array<string>();
        for (let j = 0; j < msgs.length; j++) {
          for (let k = 0; k < ends.length; k++) {
            tmpMsgs.push(msgs[j].concat(ends[k]));
          }
        }
        msgs = tmpMsgs;
      }
      messages = msgs;
      break;
    }
    case Type.OR: {
      for (let i = 0; i < type[1].length; i++) {
        if ((type[1][i] as number[]).length === 1) {
          messages.push(...generateMessages(rules, type[1][i][0]));
        } else {
          let leftMsgs = generateMessages(rules, type[1][i][0]);
          let rightMsgs = generateMessages(rules, type[1][i][1]);
          leftMsgs.forEach((l) =>
            rightMsgs.forEach((r) => messages.push(l.concat(r)))
          );
        }
      }
      break;
    }
    case Type.CHAR: {
      messages.push(type[1] as string);
      break;
    }
  }
  return messages;
}

const input = readFileSync("./example.txt", "utf8").split("\n\n");
const rules = parseRules(input[0].split("\n"));
const messages = input[1].split("\n");

const allMessages = generateMessages(rules, 0);

console.log(
  `Part 1: ${messages.reduce(
    (valids, msg) => (valids += allMessages.includes(msg) ? 1 : 0),
    0
  )}`
);
