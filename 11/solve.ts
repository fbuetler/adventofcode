import { readFileSync } from "fs";

enum Position {
  FLOOR,
  EMPTY,
  OCCUPIED,
}

class WaitingArea {
  layout: Position[][];
  rows: number;
  cols: number;
  tolerance: number;

  constructor(input: string[], tolerance: number) {
    this.rows = input.length;
    this.cols = input[0].length;
    this.layout = new Array<Position[]>();
    for (let row = 0; row < input.length; row++) {
      this.layout.push(new Array<Position>());
      for (let col = 0; col < input[row].length; col++) {
        if (input[row][col] === ".") {
          this.layout[row].push(Position.FLOOR);
        } else if (input[row][col] === "L") {
          this.layout[row].push(Position.EMPTY);
        } else {
          this.layout[row].push(Position.OCCUPIED);
        }
      }
    }
    this.tolerance = tolerance;
  }

  getAdjacentSeats(row: number, col: number): [number, number][] {
    let seats = new Array<[number, number]>();
    if (row < this.rows - 1) {
      seats.push([row + 1, col]); // down
      if (col > 0) {
        seats.push([row + 1, col - 1]); // down left
      }
      if (col < this.cols - 1) {
        seats.push([row + 1, col + 1]); // down right
      }
    }
    if (row > 0) {
      seats.push([row - 1, col]); // up
      if (col > 0) {
        seats.push([row - 1, col - 1]); // up left
      }
      if (col < this.cols - 1) {
        seats.push([row - 1, col + 1]); // up right
      }
    }
    if (col > 0) {
      seats.push([row, col - 1]); // left
    }
    if (col < this.cols - 1) {
      seats.push([row, col + 1]); // right
    }
    return seats;
  }

  getFirstSeatInLineOfSights(row: number, col: number): [number, number][] {
    let leftBoundsCheck = (row: number, col: number, i: number, j: number) =>
      col - j >= 0;
    let rightBoundsCheck = (row: number, col: number, i: number, j: number) =>
      col + j < this.cols;
    let upBoundsCheck = (row: number, col: number, i: number, j: number) =>
      row - i >= 0;
    let downBoundsCheck = (row: number, col: number, i: number, j: number) =>
      row + i < this.rows;
    let left = (col: number, j: number) => col - j;
    let right = (col: number, j: number) => col + j;
    let up = (row: number, i: number) => row - i;
    let down = (row: number, i: number) => row + i;
    let identity = (n: number, m: number) => n;

    let directions: [
      (a: number, b: number, c: number, d: number) => boolean,
      (a: number, b: number) => number,
      (a: number, b: number) => number
    ][] = [
      [rightBoundsCheck, identity, right], // right
      [downBoundsCheck, down, identity], // down
      [leftBoundsCheck, identity, left], // left
      [upBoundsCheck, up, identity], // up
      [
        (row: number, col: number, i: number, j: number) =>
          rightBoundsCheck(row, col, i, j) && upBoundsCheck(row, col, i, j),
        up,
        right,
      ], // right up
      [
        (row: number, col: number, i: number, j: number) =>
          rightBoundsCheck(row, col, i, j) && downBoundsCheck(row, col, i, j),
        down,
        right,
      ], // right down
      [
        (row: number, col: number, i: number, j: number) =>
          leftBoundsCheck(row, col, i, j) && downBoundsCheck(row, col, i, j),
        down,
        left,
      ], // left down
      [
        (row: number, col: number, i: number, j: number) =>
          leftBoundsCheck(row, col, i, j) && upBoundsCheck(row, col, i, j),
        up,
        left,
      ], //left up
    ];

    let seats = new Array<[number, number]>();
    directions.forEach(([boundsCheck, r, c]) => {
      let i = 0;
      let j = 0;
      while (
        (i === 0 && j === 0) ||
        (boundsCheck(row, col, i, j) &&
          this.layout[r(row, i)][c(col, j)] === Position.FLOOR)
      ) {
        i++;
        j++;
      }
      if (boundsCheck(row, col, i, j)) {
        seats.push([r(row, i), c(col, j)]);
      }
      i = 0;
      j = 0;
    });
    return seats;
  }

  simulateRound(): number {
    let changeCounter = 0;
    let layoutCopy = JSON.parse(JSON.stringify(this.layout)); // clone without reference
    for (let row = 0; row < this.rows; row++) {
      for (let col = 0; col < this.cols; col++) {
        let seats: [number, number][];
        if (this.tolerance === 4) {
          seats = this.getAdjacentSeats(row, col);
        } else {
          seats = this.getFirstSeatInLineOfSights(row, col);
        }
        let emptyAdjacentSeats = seats.filter(
          ([row, col]) =>
            this.layout[row][col] === Position.EMPTY ||
            this.layout[row][col] === Position.FLOOR
        ).length;
        let occupiedAdjacentSeats = seats.filter(
          ([row, col]) => this.layout[row][col] === Position.OCCUPIED
        ).length;

        if (
          this.layout[row][col] === Position.EMPTY &&
          emptyAdjacentSeats === seats.length
        ) {
          layoutCopy[row][col] = Position.OCCUPIED;
          changeCounter++;
        } else if (
          this.layout[row][col] === Position.OCCUPIED &&
          occupiedAdjacentSeats >= this.tolerance
        ) {
          layoutCopy[row][col] = Position.EMPTY;
          changeCounter++;
        }
      }
    }
    this.layout = layoutCopy;
    return changeCounter;
  }

  getNumberOfOccupiedSeats(): number {
    let counter = 0;
    for (let row = 0; row < this.rows; row++) {
      for (let col = 0; col < this.cols; col++) {
        if (this.layout[row][col] === Position.OCCUPIED) {
          counter++;
        }
      }
    }
    return counter;
  }

  toString(): string {
    let str = "";
    for (let row = 0; row < this.rows; row++) {
      for (let col = 0; col < this.cols; col++) {
        if (this.layout[row][col] === Position.EMPTY) {
          str += "L";
        } else if (this.layout[row][col] === Position.FLOOR) {
          str += ".";
        } else {
          str += "#";
        }
      }
      str += "\n";
    }
    return str;
  }
}

const input = readFileSync("./input.txt", "utf8").split("\n");

let parts: [string, number][] = [
  ["Part 1", 4],
  ["Part 2", 5],
];

parts.forEach(([name, tolerance]) => {
  const wa = new WaitingArea(input, tolerance);
  let change = -1;
  while (change !== 0) {
    change = wa.simulateRound();
  }
  console.log(`${name}: ${wa.getNumberOfOccupiedSeats()}`);
});
