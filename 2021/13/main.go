package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type dot struct {
	x int
	y int
}

type instr struct {
	axis string
	line int
}

func main() {
	input := readInput("./input.txt")

	isInstr := false
	dots := map[dot]bool{}
	instrs := []instr{}
	for _, l := range input {
		if len(l) == 0 {
			isInstr = true
			continue
		}
		if isInstr {
			var axis rune
			var line int
			fmt.Sscanf(l, "fold along %c=%d", &axis, &line)
			instrs = append(instrs, instr{
				string(axis),
				line,
			})
		} else {
			var x, y int
			fmt.Sscanf(l, "%d,%d", &x, &y)
			dots[dot{x, y}] = true
		}
	}

	solveOne(dots, instrs)
	solveTwo(dots, instrs)
}

func solveOne(dots map[dot]bool, instrs []instr) {
	instr := instrs[0]
	dots = fold(dots, instr.axis, instr.line)

	visibleDots := len(dots)
	fmt.Printf("Part 1: After the first fold there are %d dots visible\n", visibleDots)
}

func solveTwo(dots map[dot]bool, instrs []instr) {
	for _, instr := range instrs {
		dots = fold(dots, instr.axis, instr.line)
	}

	fmt.Println("Part 2:")
	printPaper(dots)
}

func fold(dots map[dot]bool, axis string, line int) map[dot]bool {
	foldedDots := map[dot]bool{}

	if axis == "x" {
		for d := range dots {
			if d.x > line {
				foldedDots[dot{line - (d.x - line), d.y}] = true
			}
			if d.x < line {
				foldedDots[d] = true
			}
		}
	} else {
		for d := range dots {
			if d.y > line {
				foldedDots[dot{d.x, line - (d.y - line)}] = true
			}
			if d.y < line {
				foldedDots[d] = true
			}
		}
	}

	return foldedDots
}

func printPaper(dots map[dot]bool) {
	maxX := 0
	maxY := 0
	for d := range dots {
		maxX = max(maxX, d.x)
		maxY = max(maxY, d.y)
	}

	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if _, exists := dots[dot{j, i}]; exists {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
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
