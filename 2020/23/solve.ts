import { readFileSync } from "fs";

class Circle {
  nextCup = new Map<number, number>();
  current: number;
  constructor(cups: number[]) {
    this.current = cups[0];
    for (let i = 0; i < cups.length - 1; i++) {
      this.nextCup.set(cups[i], cups[i + 1]);
    }
    this.nextCup.set(cups[cups.length - 1], cups[0]);
  }

  pickThreeCupsAfterCurrent(): number[] {
    const pickedUp = new Array<number>();
    let next = this.nextCup.get(this.current);
    for (let i = 0; i < 3; i++) {
      pickedUp.push(next);
      next = this.nextCup.get(next);
    }
    this.nextCup.set(this.current, next);
    return pickedUp;
  }

  selectDestinationCup(pickedUp: number[]): number {
    let cup = this.current - 1;
    while (pickedUp.includes(cup) || cup === 0) {
      cup--;
      if (cup <= 0) {
        cup = this.nextCup.size;
      }
    }
    return cup;
  }

  placeCupsAfterDestination(cups: number[], destination: number) {
    let next = this.nextCup.get(destination);
    this.nextCup.set(destination, cups[0]);
    this.nextCup.set(cups[cups.length - 1], next);
    this.current = this.nextCup.get(this.current);
  }

  labelsAfterCupOne(len: number): Array<number> {
    let cups = new Array<number>();
    let next = this.nextCup.get(1);
    let i = 0;
    while (i < len) {
      cups.push(next);
      next = this.nextCup.get(next);
      i++;
    }
    return cups;
  }

  toString(): string {
    let str = ` (${this.current})`;
    let next = this.nextCup.get(this.current);
    let i = 1;
    while (next !== this.current) {
      str += ` ${next}`;
      next = this.nextCup.get(next);
      if (i % 100 === 0) {
        str += "\n";
      }
      i++;
    }
    return str;
  }
}

function part1(cups: number[]) {
  let move = 0;
  const circle = new Circle(cups);
  while (move < 100) {
    const pickedUp = circle.pickThreeCupsAfterCurrent();
    const destination = circle.selectDestinationCup(pickedUp);
    circle.placeCupsAfterDestination(pickedUp, destination);
    move++;
  }
  console.log(`Part 1: ${circle.labelsAfterCupOne(8).join("")}`);
}

function part2(cups: number[]) {
  const totalCups = 1000000;
  const totalMoves = 10000000;

  const millionCups = cups;
  let nextLabel = 10;
  while (millionCups.length < totalCups) {
    millionCups.push(nextLabel);
    nextLabel++;
  }

  const bigCircle = new Circle(millionCups);
  let move = 0;
  while (move < totalMoves) {
    const pickedUp = bigCircle.pickThreeCupsAfterCurrent();
    const destination = bigCircle.selectDestinationCup(pickedUp);
    bigCircle.placeCupsAfterDestination(pickedUp, destination);
    move++;
  }

  console.log(
    `Part 2: ${bigCircle.labelsAfterCupOne(2).reduce((a, b) => a * b)}`
  );
}

const input = readFileSync("./input.txt", "utf8");
console.log(`input: ${input}`);
const cups = input.split("").map((el) => +el);

part1(cups);
part2(cups);
