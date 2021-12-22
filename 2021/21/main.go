package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type dice struct {
	value      int
	totalRolls int
}

type player struct {
	id    int
	pos   int
	score int
}

func (d *dice) roll() int {
	v := d.value
	d.value = (d.value)%100 + 1
	d.totalRolls++
	return v
}

func (p *player) turn(d *dice) {
	s1 := d.roll()
	s2 := d.roll()
	s3 := d.roll()
	p.move(s1 + s2 + s3)
	// fmt.Printf("Player %d rolls %d+%d+%d and moves to space %d for a total score of %d.\n", p.id, s1, s2, s3, p.pos, p.score)
}

func (p *player) move(steps int) {
	p.pos = (p.pos+steps-1)%10 + 1
	p.score += p.pos
}

func (p *player) wins() bool {
	return p.score >= 1000
}

func main() {
	input := readInput("./input.txt")

	var pos1, pos2 int
	fmt.Sscanf(input[0], "Player 1 starting position: %d", &pos1)
	fmt.Sscanf(input[1], "Player 2 starting position: %d", &pos2)

	solveOne(pos1, pos2)
	solveTwo(pos1, pos2)
}

func solveOne(initPos1, initPos2 int) {
	d := &dice{
		value: 1,
	}
	p1 := player{
		id:    1,
		pos:   initPos1,
		score: 0,
	}
	p2 := player{
		id:    2,
		pos:   initPos2,
		score: 0,
	}
	for {
		p1.turn(d)
		if p1.wins() {
			break
		}
		p2.turn(d)
		if p2.wins() {
			break
		}
	}

	var s int
	if p1.wins() {
		s = p2.score
	} else {
		s = p1.score
	}
	fmt.Printf("Part 1: %d * %d = %d\n", s, d.totalRolls, s*d.totalRolls)
}

func solveTwo(initPos1, initPos2 int) {
	// pos1, pos2, score1, score2 -> number of universes with this state
	universes := [11][11][22][22]int{}
	universes[initPos1][initPos2][0][0] = 1

	for score1 := 0; score1 < 21; score1++ {
		for score2 := 0; score2 < 21; score2++ {
			for pos1 := 1; pos1 <= 10; pos1++ {
				for pos2 := 1; pos2 <= 10; pos2++ {
					for r1 := 1; r1 <= 3; r1++ {
						for r2 := 1; r2 <= 3; r2++ {
							for r3 := 1; r3 <= 3; r3++ {
								newPos1 := (pos1+r1+r2+r3-1)%10 + 1
								newScore1 := min(score1+newPos1, 21)
								if newScore1 == 21 {
									// if player 1 wins here, the game is over
									universes[newPos1][pos2][newScore1][score2] += universes[pos1][pos2][score1][score2]
								} else {
									// otherwise its player 2's turn
									for r4 := 1; r4 <= 3; r4++ {
										for r5 := 1; r5 <= 3; r5++ {
											for r6 := 1; r6 <= 3; r6++ {
												newPos2 := (pos2+r4+r5+r6-1)%10 + 1
												newScore2 := min(score2+newPos2, 21)
												universes[newPos1][newPos2][newScore1][newScore2] += universes[pos1][pos2][score1][score2]
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	wins1 := 0
	wins2 := 0
	for pos1 := 1; pos1 <= 10; pos1++ {
		for pos2 := 1; pos2 <= 10; pos2++ {
			for score := 0; score < 21; score++ {
				wins1 += universes[pos1][pos2][21][score]
				wins2 += universes[pos1][pos2][score][21]
			}
		}
	}
	fmt.Printf("Part 2: %d\n", max(wins1, wins2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
