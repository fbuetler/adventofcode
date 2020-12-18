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
    let next = i < str.length - 1 ? str[i + 1].trim() : "";
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
  console.log(tokens);
  return tokens;
}

function parse(tokens: [Token, string][]): Expr {
  return parseAdd(tokens, 0)[0];
}

function parseAdd(tokens: [Token, string][], pos: number): [Expr, number] {
  let left: Expr;
  let right: Expr;
  [left, pos] = parseMult(tokens, pos);
  if (pos >= tokens.length) {
    return [left, pos];
  }
  let [token, value] = tokens[pos];
  while (token === Token.ADD) {
    pos++;
    [right, pos] = parseMult(tokens, pos);
    left = new BinOp(left, Op.ADD, right);
    if (pos >= tokens.length) {
      return [left, pos];
    }
    [token, value] = tokens[pos];
  }
  return [left, pos];
}

function parseMult(tokens: [Token, string][], pos: number): [Expr, number] {
  let left: Expr;
  let right: Expr;
  [left, pos] = parseNumber(tokens, pos);
  if (pos >= tokens.length) {
    return [left, pos];
  }
  let [token, value] = tokens[pos];
  while (token === Token.MULT) {
    pos++;
    [right, pos] = parseNumber(tokens, pos);
    left = new BinOp(left, Op.MULT, right);
    if (pos >= tokens.length) {
      return [left, pos];
    }
    [token, value] = tokens[pos];
  }
  return [left, pos];
}

function parseNumber(tokens: [Token, string][], pos: number): [Expr, number] {
  const [token, value] = tokens[pos++];
  if (token === Token.NUMBER) {
    return [new Num(+value), pos];
  } else if (token === Token.OPEN) {
    let e: Expr;
    [e, pos] = parseAdd(tokens, pos);
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
  console.log(result);
  return result;
}

const input = readFileSync("./example.txt", "utf8").split("\n");

console.log(
  `Part 1: ${input.reduce(
    (sum, str) => (sum += evaluate(parse(tokenize(str)))),
    0
  )}`
);
