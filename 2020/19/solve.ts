import { readFileSync } from "fs";

enum Type {
  SEQ = "seq",
  OR = "or",
  CHAR = "char",
}

type Rules = [Type, string | number[][] | number[]]; // char - options - sequence

function parseRules(rules: string[]): Map<number, Rules> {
  let parsed = new Map<number, Rules>();
  for (let i = 0; i < rules.length; i++) {
    let parts = rules[i].split(":");
    let key = +parts[0];
    let value = parts[1].trim();
    if (value.includes("|")) {
      let options = new Array<Array<number>>();
      let ops = value.split("|");
      ops.forEach((op, index) =>
        options.push(
          ops[index]
            .trim()
            .split(" ")
            .map((num) => +num)
        )
      );
      parsed.set(key, [Type.OR, options]);
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

function generateMessages(
  rules: Map<number, Rules>,
  rule: number,
  subSequenceCache: Map<number, string[]>
): string[] {
  if (subSequenceCache.has(rule)) {
    return subSequenceCache.get(rule);
  }
  const type = rules.get(rule);
  let messages = new Array<string>();
  switch (type[0]) {
    case Type.SEQ: {
      let msgs = [""];
      for (let i = 0; i < type[1].length; i++) {
        const ends = generateMessages(
          rules,
          type[1][i] as number,
          subSequenceCache
        );
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
        let accMsgs = [""];
        for (let j = 0; j < (type[1][i] as number[]).length; j++) {
          const tmpMsgs = new Array<string>();
          const msgs = generateMessages(rules, type[1][i][j], subSequenceCache);
          accMsgs.forEach((accMsg) =>
            msgs.forEach((msg) => tmpMsgs.push(accMsg.concat(msg)))
          );
          accMsgs = tmpMsgs;
        }
        messages.push(...accMsgs);
      }
      break;
    }
    case Type.CHAR: {
      messages.push(type[1] as string);
      break;
    }
  }
  subSequenceCache.set(rule, messages);
  return messages;
}

const input = readFileSync("./input.txt", "utf8").split("\n\n");
const rules = parseRules(input[0].split("\n"));
const messages = input[1].split("\n");

const subSequenceCache = new Map<number, string[]>();
const allMessages = generateMessages(rules, 0, subSequenceCache);

console.log(
  `Part 1: ${messages.reduce(
    (valids, msg) => (valids += allMessages.includes(msg) ? 1 : 0),
    0
  )}`
);

/*
0: 8 11
8: 42 | 42 8
11: 42 31 | 42 11 31

Always:
* only contains 42 and 31
* has more 42 than 31
* has at least 2 42
* has at least 1 31

*/
let valids = 0;
const messagesRule31 = subSequenceCache.get(31);
const messagesRule42 = subSequenceCache.get(42);
for (let i = 0; i < messages.length; i++) {
  const parts = new Array<string>();
  for (let j = 0; j < messages[i].length / 8; j++) {
    parts.push(messages[i].substr(j * 8, 8));
  }
  let fitsRule31 = 0;
  let fitsRule42 = 0;
  let fromLeft = 0;
  while (fromLeft < parts.length && messagesRule42.includes(parts[fromLeft])) {
    fitsRule42++;
    fromLeft++;
  }
  let fromRight = parts.length - 1;
  while (fromRight >= 0 && messagesRule31.includes(parts[fromRight])) {
    fitsRule31++;
    fromRight--;
  }

  if (
    fitsRule42 > 1 &&
    fitsRule42 > fitsRule31 &&
    fitsRule31 > 0 &&
    fitsRule42 + fitsRule31 === parts.length
  ) {
    valids++;
  }
}
console.log(`Part 2: ${valids}`);
