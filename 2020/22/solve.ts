import { readFileSync } from "fs";

function playGamePart1(player1: number[], player2: number[]): number {
  let round = 1;
  while (player1.length !== 0 && player2.length !== 0) {
    console.log(`-- Round ${round} --`);
    console.log(`Player 1's deck: ${player1}`);
    console.log(`Player 2's deck: ${player2}`);
    const card1 = player1.shift();
    const card2 = player2.shift();
    console.log(`Player 1 play: ${card1}`);
    console.log(`Player 2 play: ${card2}`);
    if (card1 > card2) {
      player1.push(card1);
      player1.push(card2);
      console.log("Player 1 wins the round!");
    } else {
      player2.push(card2);
      player2.push(card1);
      console.log("Player 2 wins the round!");
    }
    round++;
  }
  console.log(player1);
  console.log(player2);
  return calcScore(player1, player2);
}

function playGamePart2(
  player1: number[],
  player2: number[],
  game: number
): number {
  let round = 1;
  let winner = 0;
  let history1 = new Array<Array<number>>();
  let history2 = new Array<Array<number>>();
  console.log(`=== Game ${game} ===`);
  while (player1.length !== 0 && player2.length !== 0) {
    console.log(`\n-- Round ${round} (Game ${game}) --`);
    console.log(`Player 1's deck: ${player1}`);
    console.log(`Player 2's deck: ${player2}`);

    if (
      history1.filter((el) => JSON.stringify(el) === JSON.stringify(player1))
        .length !== 0 ||
      history2.filter((el) => JSON.stringify(el) === JSON.stringify(player2))
        .length !== 0
    ) {
      console.log("Player 1 wins due the reoccuring round configuration");
      return 1;
    }
    history1.push(JSON.parse(JSON.stringify(player1)));
    history2.push(JSON.parse(JSON.stringify(player2)));

    const card1 = player1.shift();
    const card2 = player2.shift();
    console.log(`Player 1 play: ${card1}`);
    console.log(`Player 2 play: ${card2}`);

    if (player1.length >= card1 && player2.length >= card2) {
      console.log("Playing a sub-game to determine the winner...\n");
      winner = playGamePart2(
        JSON.parse(JSON.stringify(player1.slice(0, card1))),
        JSON.parse(JSON.stringify(player2.slice(0, card2))),
        game + 1
      );
    } else {
      if (card1 > card2) {
        winner = 1;
      } else {
        winner = 2;
      }
    }
    console.log(`Player ${winner} wins round ${round} of game ${game}!`);

    if (winner === 1) {
      player1.push(card1);
      player1.push(card2);
    } else {
      player2.push(card2);
      player2.push(card1);
    }

    round++;
  }
  console.log(`The winner of game ${game} is palyer ${winner}\n`);
  if (game === 1) {
    console.log(
      `\n== Post-game results ==\nPlayer 1's deck: ${player1}\nPlayer 2's deck: ${player2}`
    );
    return calcScore(player1, player2);
  } else {
    console.log(`...anyway, back to game ${game - 1}`);
    return winner;
  }
}

function calcScore(player1: number[], player2: number[]): number {
  let score = 0;
  if (player1.length !== 0) {
    player1.forEach(
      (card, index, arr) => (score += (arr.length - index) * card)
    );
  } else {
    player2.forEach(
      (card, index, arr) => (score += (arr.length - index) * card)
    );
  }
  return score;
}

const input = readFileSync("./input.txt", "utf8").split("\n\n");

let player1 = input[0]
  .split("\n")
  .slice(1)
  .map((el) => +el);
let player2 = input[1]
  .split("\n")
  .slice(1)
  .map((el) => +el);

console.log(
  `Part 1: ${playGamePart1(
    JSON.parse(JSON.stringify(player1)),
    JSON.parse(JSON.stringify(player2))
  )}`
);

console.log(`Part 2: ${playGamePart2(player1, player2, 1)}`);
