import { readFileSync } from "fs";

class PasswordPolicy {
  lowerBound: number;
  upperBound: number;
  char: string;
  password: string;

  regex = new RegExp(`^(\\d+?)-(\\d+?)\\s(\\w?):\\s(\\w+?)$`);

  constructor(policy: string) {
    const match = policy.match(this.regex);
    this.lowerBound = +match[1];
    this.upperBound = +match[2];
    this.char = match[3];
    this.password = match[4];
  }

  countOccurences(): number {
    return this.password.split(this.char).length - 1;
  }

  isValidBySledRental(): boolean {
    const occurences = this.countOccurences();
    return this.lowerBound <= occurences && occurences <= this.upperBound;
  }

  xor(a: boolean, b: boolean): boolean {
    return (a || b) && !(a && b);
  }

  isValidByToboggan(): boolean {
    return this.xor(
      this.password[this.lowerBound - 1] === this.char,
      this.password[this.upperBound - 1] === this.char
    );
  }
}

function solve(input: PasswordPolicy[]) {
  let matchesPartOne = 0;
  let matchesPartTwo = 0;
  for (let i = 0; i < input.length; i++) {
    if (input[i].isValidBySledRental()) {
      matchesPartOne++;
    }
    if (input[i].isValidByToboggan()) {
      matchesPartTwo++;
    }
  }
  console.log("Part 1");
  console.log(matchesPartOne);
  console.log("Part 2");
  console.log(matchesPartTwo);
}

const data = readFileSync("./input.txt", "utf8");
const input = data.split("\n").map((el) => new PasswordPolicy(el));

solve(input);
