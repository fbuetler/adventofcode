import { readFileSync } from "fs";

enum Direction {
  NORTH,
  EAST,
  SOUTH,
  WEST,
}

function rotate(action: string, amount: number, byWaypoint: boolean) {
  if (byWaypoint) {
    let relativeX = x - shipX;
    let relativeY = y - shipY;
    if (
      (action === "R" && amount === 90) ||
      (action === "L" && amount === 270)
    ) {
      const tmp = relativeX;
      relativeX = relativeY;
      relativeY = -tmp;
      x = shipX + relativeX;
      y = shipY + relativeY;
    } else if ((action === "R" || action === "L") && amount === 180) {
      x = shipX - relativeX;
      y = shipY - relativeY;
    } else if (
      (action === "R" && amount === 270) ||
      (action === "L" && amount === 90)
    ) {
      const tmp = relativeX;
      relativeX = -relativeY;
      relativeY = tmp;
      x = shipX + relativeX;
      y = shipY + relativeY;
    }
  } else {
    let r = 0;
    if (
      (action === "R" && amount === 90) ||
      (action === "L" && amount === 270)
    ) {
      r = 1;
    } else if ((action === "R" || action === "L") && amount === 180) {
      r = 2;
    } else if (
      (action === "R" && amount === 270) ||
      (action === "L" && amount === 90)
    ) {
      r = 3;
    }
    direction = (direction + r) % 4;
  }
}

function moveForward(action: string, amount: number, byWaypoint: boolean) {
  if (byWaypoint) {
    for (let i = 0; i < amount; i++) {
      const relativeX = x - shipX;
      const relativeY = y - shipY;
      shipX = x;
      shipY = y;
      x += relativeX;
      y += relativeY;
    }
  } else {
    switch (direction) {
      case Direction.NORTH: {
        shipY += amount;
        break;
      }
      case Direction.EAST: {
        shipX += amount;
        break;
      }
      case Direction.SOUTH: {
        shipY -= amount;
        break;
      }
      case Direction.WEST: {
        shipX -= amount;
        break;
      }
    }
  }
}

function moveInDirection(action: string, amount: number, byWaypoint: boolean) {
  if (byWaypoint) {
    switch (action) {
      case "N": {
        y += amount;
        break;
      }
      case "E": {
        x += amount;
        break;
      }
      case "S": {
        y -= amount;
        break;
      }
      case "W": {
        x -= amount;
        break;
      }
    }
  } else {
    switch (action) {
      case "N": {
        shipY += amount;
        break;
      }
      case "E": {
        shipX += amount;
        break;
      }
      case "S": {
        shipY -= amount;
        break;
      }
      case "W": {
        shipX -= amount;
        break;
      }
    }
  }
}

function navigate(action: string, amount: number, byWaypoint: boolean) {
  switch (action) {
    case "N": {
    }
    case "E": {
    }
    case "S": {
    }
    case "W": {
      moveInDirection(action, amount, byWaypoint);
      break;
    }
    case "R": {
    }
    case "L": {
      rotate(action, amount % 360, byWaypoint);
      break;
    }
    case "F": {
      moveForward(action, amount, byWaypoint);
      break;
    }
  }
}

const input = readFileSync("./input.txt", "utf8").split("\n");

let x: number;
let y: number;
let shipX: number;
let shipY: number;
let direction: Direction;
const parts: [string, boolean][] = [
  ["Part 1", false],
  ["Part 2", true],
];
parts.forEach(([part, byWaypoint]) => {
  x = 10;
  y = 1;
  shipX = 0;
  shipY = 0;
  direction = Direction.EAST;
  for (let i = 0; i < input.length; i++) {
    let action = input[i].substring(0, 1);
    let amount = +input[i].substring(1);
    navigate(action, amount, byWaypoint);
  }
  console.log(`${part}: ${Math.abs(shipX) + Math.abs(shipY)}`);
});
