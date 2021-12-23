package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var roomNr = map[string]int{
	"A": 2,
	"B": 4,
	"C": 6,
	"D": 8,
}

var energyPerStep = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

type burrow struct {
	pos []pos
}

type pos struct {
	specie string
	x      int
	y      int
}

func (b *burrow) possibleMoves(x, y int) []pos {
	stype := b.atPos(x, y)
	if stype == "." {
		return []pos{}
	}

	// skip if it is in its destination room
	if x == roomNr[stype] && b.roomIsSettled(roomNr[stype], stype) {
		return []pos{}
	}

	moves := []pos{}
	// if an amphipod is in the hallway, it will move in its room iff
	// its their destination room and that room contains no amphipods from another type
	nr := roomNr[stype]
	if b.roomIsMovable(nr, stype) {
		spot := 1
		if b.roomIsEmpty(nr) {
			spot = 2
		}
		if b.pathExists(x, y, nr, spot) {
			moves = append(moves, pos{"", nr, spot})
		}
	}
	// if its in the hallway and cant go to its room, it wont move
	if y == 0 {
		return moves
	}

	// move directly to destination room

	// move to the hallway
	for i := 0; i < 11; i++ {
		// an amphipod never wait in front of a room
		if i == 2 || i == 4 || i == 6 || i == 8 {
			continue
		}
		if !b.pathExists(x, y, i, 0) {
			continue
		}
		moves = append(moves, pos{"", i, 0})
	}

	return moves
}

func (b *burrow) move(fromX, fromY, toX, toY int) int {
	stype := b.atPos(fromX, fromY)
	if stype == "." {
		log.Fatal("cannot move empty space")
	}

	for i, p := range b.pos {
		if p.equals(pos{"", fromX, fromY}) {
			b.pos[i].x = toX
			b.pos[i].y = toY
		}
	}

	steps := abs(toX-fromX) + abs(toY-fromY)
	return energyPerStep[stype] * steps
}

func (b *burrow) roomIsMovable(nr int, stype string) bool {
	if b.roomIsEmpty(nr) {
		return true
	}
	fstSpot := b.atPos(nr, 1)
	sndSpot := b.atPos(nr, 2)
	return fstSpot == "." && (sndSpot == "." || sndSpot == stype)
}

func (b *burrow) roomIsSettled(nr int, stype string) bool {
	fstSpot := b.atPos(nr, 1)
	sndSpot := b.atPos(nr, 2)
	return (fstSpot == "." || fstSpot == stype) && sndSpot == stype
}

func (b *burrow) roomIsEmpty(nr int) bool {
	return b.atPos(nr, 1) == "." && b.atPos(nr, 2) == "."
}

func (b *burrow) pathExists(fromX, fromY, toX, toY int) bool {
	minX := min(fromX, toX)
	maxX := max(fromX, toX)
	minY := min(fromY, toY)
	maxY := max(fromY, toY)
	for i := minX; i <= maxX; i++ {
		if i == fromX {
			continue
		}
		if b.atPos(i, minY) != "." {
			return false
		}
	}
	nr := -1
	for _, v := range roomNr {
		if fromX == v || toX == v {
			nr = v
		}
	}
	for i := minY; i <= maxY; i++ {
		if i == fromY {
			continue
		}
		if b.atPos(nr, i) != "." {
			return false
		}
	}
	return true
}

func (b *burrow) complete() bool {
	for _, p := range b.pos {
		if p.specie == "" {
			continue
		}
		if p.x != roomNr[p.specie] {
			return false
		}
		if p.y == 0 {
			return false
		}
	}
	return true
}

func (b *burrow) toString() string {
	str := ""
	maxY := 0
	for _, p := range b.pos {
		maxY = max(maxY, p.y)
	}
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= 11; j++ {
			str += b.atPos(j, i)
		}
		str += "\n"
	}
	return str
}

func (b *burrow) atPos(x, y int) string {
	for _, p := range b.pos {
		if p.equals(pos{"", x, y}) && p.specie != "" {
			return p.specie
		}
	}

	if y > 0 && (x != 2 && x != 4 && x != 6 && x != 8) {
		return " "
	}
	if y == 0 && (x < 0 || x > 11) {
		return " "
	}
	maxY := 0
	for _, p := range b.pos {
		maxY = max(maxY, p.y)
	}
	if y < 0 || y > maxY {
		return " "
	}

	return "."
}

func (b *burrow) clone() burrow {
	c := burrow{
		pos: make([]pos, len(b.pos)),
	}
	copy(c.pos, b.pos)
	return c
}

func (p *pos) equals(other pos) bool {
	return p.x == other.x && p.y == other.y
}

func main() {
	input := readInput("./input.txt")

	b := burrow{
		pos: make([]pos, 8),
	}
	for _, room := range []int{3, 5, 7, 9} {
		sp1 := string(input[2][room])
		sp2 := string(input[3][room])
		b.pos = append(b.pos, pos{sp1, room - 1, 1})
		b.pos = append(b.pos, pos{sp2, room - 1, 2})
	}
	solve(b, 1)

	b.pos = make([]pos, 16)
	for _, room := range []int{3, 5, 7, 9} {
		sp1 := string(input[2][room])
		sp2 := string(input[3][room])
		b.pos = append(b.pos, pos{sp1, room - 1, 1})
		b.pos = append(b.pos, pos{sp2, room - 1, 4})
	}
	/*
		#D#C#B#A#
		#D#B#A#C#
	*/
	b.pos = append(b.pos, pos{"D", 2, 2})
	b.pos = append(b.pos, pos{"C", 4, 2})
	b.pos = append(b.pos, pos{"B", 6, 2})
	b.pos = append(b.pos, pos{"A", 8, 2})
	b.pos = append(b.pos, pos{"D", 2, 3})
	b.pos = append(b.pos, pos{"B", 4, 3})
	b.pos = append(b.pos, pos{"A", 6, 3})
	b.pos = append(b.pos, pos{"C", 8, 3})
	solve(b, 2)
}

func solve(b burrow, part int) {
	minEnergy, trace := simulate(b, 0, 1, "")
	fmt.Printf("Part %d: %d\n", part, minEnergy)
	fmt.Println(trace)
}

func simulate(b burrow, rootEnergy, level int, rootTrace string) (int, string) {
	minEnergy := int(math.Pow(10, 10))
	var levelTrace string
	for _, a := range b.pos {
		moves := b.possibleMoves(a.x, a.y)
		for _, m := range moves {
			bc := b.clone()
			stepEnergy := bc.move(a.x, a.y, m.x, m.y)
			stepStr := fmt.Sprintf("level %d, energy %d:\n%s\n\n", level, stepEnergy, bc.toString())
			if bc.complete() {
				return stepEnergy, stepStr
			}
			energy, stepTrace := simulate(bc, 0, level+1, rootTrace)
			if stepEnergy+energy < minEnergy {
				levelTrace = stepStr + stepTrace
			}
			minEnergy = min(minEnergy, stepEnergy+energy)
		}
	}
	return rootEnergy + minEnergy, rootTrace + levelTrace
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
