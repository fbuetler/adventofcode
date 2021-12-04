package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

type Board struct {
	position map[int]Pair
	marked   [][]bool
	rows     [][]int
}

func (b *Board) init(rows [][]int) {
	b.rows = rows

	size := len(rows)
	b.marked = make([][]bool, size)
	for i := range b.marked {
		b.marked[i] = make([]bool, size)
	}

	b.position = make(map[int]Pair)
	for i, row := range rows {
		for j, rowEntry := range row {
			b.position[rowEntry] = Pair{x: i, y: j}
		}
	}
}

func (b *Board) reset() {
	for i, rows := range b.marked {
		for j := range rows {
			b.marked[i][j] = false
		}
	}
}

func (b *Board) mark(number int) {
	if pos, ok := b.position[number]; ok {
		b.marked[pos.x][pos.y] = true
	}
}

func (b *Board) wins() bool {
	for i, row := range b.rows {
		rowMarked := true
		colMarked := true
		for j := range row {
			rowMarked = rowMarked && b.marked[i][j]
			colMarked = colMarked && b.marked[j][i]
		}
		if rowMarked || colMarked {
			return true
		}
	}
	return false
}

func (b *Board) sumUnmarked() int {
	sum := 0
	for i, row := range b.rows {
		for j, rowEntry := range row {
			if !b.marked[i][j] {
				sum += rowEntry
			}
		}
	}
	return sum
}

func main() {
	input := readInput("./input.txt")

	// parse drawn numbers
	numbers := []int{}
	for _, n := range strings.Split(input[0], ",") {
		number, err := strconv.Atoi(string(n))
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// parse boards
	i := 2
	boards := []Board{}
	for i < len(input) {
		rows := [][]int{}
		for j := 0; j < 5; j++ {
			row := []int{}
			for _, r := range strings.Fields(input[i+j]) {
				rowEntry, err := strconv.Atoi(string(r))
				if err != nil {
					log.Fatal(err)
				}
				row = append(row, rowEntry)
			}
			rows = append(rows, row)
		}
		board := Board{}
		board.init(rows)
		boards = append(boards, board)
		i += 6
	}

	solveOne(numbers, boards)
	for _, board := range boards {
		board.reset()
	}
	solveTwo(numbers, boards)
}

func solveOne(numbers []int, boards []Board) {
	for _, number := range numbers {
		for _, board := range boards {
			board.mark(number)
			if board.wins() {
				sumUnmarked := board.sumUnmarked()
				fmt.Printf("Part 1: Final score %d\n", number*sumUnmarked)
				return
			}
		}
	}
}

func solveTwo(numbers []int, boards []Board) {
	remainingBoards := len(boards)
	winningBoards := make([]bool, len(boards))
	for _, number := range numbers {
		for j, board := range boards {
			if winningBoards[j] {
				continue
			}

			board.mark(number)
			if board.wins() {
				winningBoards[j] = true
				if remainingBoards == 1 {
					sumUnmarked := board.sumUnmarked()
					fmt.Printf("Part 2: Final score %d\n", number*sumUnmarked)
					return
				}
				remainingBoards--
			}
		}
	}
	fmt.Printf("Part 2: \n")
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
