import { readFileSync } from "fs";

enum Token {
  NUMBER,
  ADD,
  MULT,
  OPEN,
  CLOSE,
}

enum Op {
  ADD,
  MULT,
}

type Expr = BinOp | Num;

class Num {
  value: number;
  constructor(value: number) {
    this.value = value;
  }
  visit(visitor: Visitor): number {
    return visitor.visitNum(this);
  }
  toString(): string {
    return String(this.value);
  }
}

class BinOp {
  left: Expr;
  op: Op;
  right: Expr;
  constructor(left: Expr, op: Op, right: Expr) {
    this.left = left;
    this.op = op;
    this.right = right;
  }
  visit(visitor: Visitor): number {
    return visitor.visitBinOp(this);
  }
  toString(): string {
    let l: string;
    let r: string;
    if (this.left instanceof Num) {
      l = this.left.toString();
    } else {
      l = `(${this.right})`;
    }
    if (this.right instanceof Num) {
      r = this.right.toString();
    } else {
      r = `(${this.right})`;
    }
    return `${l}${this.op === Op.ADD ? "+" : "*"}${r}`;
  }
}

class Visitor {
  visitNum(num: Num): number {
    return num.value;
  }
  visitBinOp(binOp: BinOp): number {
    switch (binOp.op) {
      case Op.ADD: {
        return binOp.left.visit(this) + binOp.right.visit(this);
      }
      case Op.MULT: {
        return binOp.left.visit(this) * binOp.right.visit(this);
      }
    }
  }
}

function tokenize(str: string): [Token, string][] {
  str = str.trim();
  let tokens = new Array<[Token, string]>();
  let curr = "";
  for (let i = 0; i < str.length; i++) {
    curr += str[i].trim();
    let next = i < str.length - 1 ? str[i + 1].trim() : "-";
    if (curr === "") {
      continue;
    } else if (!isNaN(+curr) && isNaN(+next)) {
      tokens.push([Token.NUMBER, curr]);
      curr = "";
    } else if (curr === "+") {
      tokens.push([Token.ADD, curr]);
      curr = "";
    } else if (curr === "*") {
      tokens.push([Token.MULT, curr]);
      curr = "";
    } else if (curr === "(") {
      tokens.push([Token.OPEN, curr]);
      curr = "";
    } else if (curr === ")") {
      tokens.push([Token.CLOSE, curr]);
      curr = "";
    }
  }
  return tokens;
}

function parseLeftToRight(tokens: [Token, string][]): Expr {
  return parseBinOp(tokens, 0)[0];
}

function parseAddBeforeMult(tokens: [Token, string][]): Expr {
  return parseMult(tokens, 0)[0];
}

function parseMult(tokens: [Token, string][], pos: number): [Expr, number] {
  let left: Expr;
  let right: Expr;
  [left, pos] = parseAdd(tokens, pos);
  if (pos >= tokens.length) {
    return [left, pos];
  }
  let [token, value] = tokens[pos];
  while (token === Token.MULT) {
    pos++;
    [right, pos] = parseAdd(tokens, pos);
    left = new BinOp(left, Op.MULT, right);
    if (pos >= tokens.length) {
      return [left, pos];
    }
    [token, value] = tokens[pos];
  }
  return [left, pos];
}

function parseAdd(tokens: [Token, string][], pos: number): [Expr, number] {
  let left: Expr;
  let right: Expr;
  [left, pos] = parseNumberAddBeforeMult(tokens, pos);
  if (pos >= tokens.length) {
    return [left, pos];
  }
  let [token, value] = tokens[pos];
  while (token === Token.ADD) {
    pos++;
    [right, pos] = parseNumberAddBeforeMult(tokens, pos);
    left = new BinOp(left, Op.ADD, right);
    if (pos >= tokens.length) {
      return [left, pos];
    }
    [token, value] = tokens[pos];
  }
  return [left, pos];
}

function parseBinOp(tokens: [Token, string][], pos: number): [Expr, number] {
  let left: Expr;
  let right: Expr;
  [left, pos] = parseNumberLeftToRight(tokens, pos);
  if (pos >= tokens.length) {
    return [left, pos];
  }
  let [token, value] = tokens[pos];
  while (token === Token.ADD || token === Token.MULT) {
    pos++;
    [right, pos] = parseNumberLeftToRight(tokens, pos);
    left = new BinOp(left, token === Token.ADD ? Op.ADD : Op.MULT, right);
    if (pos >= tokens.length) {
      return [left, pos];
    }
    [token, value] = tokens[pos];
  }
  return [left, pos];
}

function parseNumberLeftToRight(
  tokens: [Token, string][],
  pos: number
): [Expr, number] {
  const [token, value] = tokens[pos++];
  if (token === Token.NUMBER) {
    return [new Num(+value), pos];
  } else if (token === Token.OPEN) {
    let e: Expr;
    [e, pos] = parseBinOp(tokens, pos);
    pos++;
    return [e, pos];
  }
}

function parseNumberAddBeforeMult(
  tokens: [Token, string][],
  pos: number
): [Expr, number] {
  const [token, value] = tokens[pos++];
  if (token === Token.NUMBER) {
    return [new Num(+value), pos];
  } else if (token === Token.OPEN) {
    let e: Expr;
    [e, pos] = parseMult(tokens, pos);
    pos++;
    return [e, pos];
  }
}

function evaluate(expr: Expr): number {
  const visitor = new Visitor();
  let result = 0;
  if (expr instanceof Num) {
    result = visitor.visitNum(expr);
  } else {
    result = visitor.visitBinOp(expr);
  }
  return result;
}

const input = readFileSync("./input.txt", "utf8").split("\n");

console.log(
  `Part 1: ${input.reduce(
    (sum, str) => (sum += evaluate(parseLeftToRight(tokenize(str)))),
    0
  )}`
);
console.log(
  `Part 2: ${input.reduce(
    (sum, str) => (sum += evaluate(parseAddBeforeMult(tokenize(str)))),
    0
  )}`
);
