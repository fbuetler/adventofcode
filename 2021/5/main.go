package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	input := readInput("./input.txt")

	lines := []line{}
	re := regexp.MustCompile("(?P<x1>\\d*),(?P<y1>\\d*) -> (?P<x2>\\d*),(?P<y2>\\d*)")
	for _, i := range input {
		match := re.FindStringSubmatch(i)
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])
		l := line{
			x1,
			y1,
			x2,
			y2,
		}
		lines = append(lines, l)
	}

	solveOne(lines)
	solveTwo(lines)
}

func solveOne(lines []line) {
	vents := createVents(lines)

	for _, l := range lines {
		drawHorizontalLines(vents, l)
		drawVerticalLines(vents, l)
	}

	overlaps := countOverlaps(vents)
	fmt.Printf("Part 1: Number of overlapping vents %d\n", overlaps)
}

func solveTwo(lines []line) {
	vents := createVents(lines)

	for _, l := range lines {
		drawHorizontalLines(vents, l)
		drawVerticalLines(vents, l)
		drawDiagonalLines(vents, l)
	}

	overlaps := countOverlaps(vents)
	fmt.Printf("Part 2: Number of overlapping vents %d\n", overlaps)
}

func createVents(lines []line) [][]int {
	maxX := 0
	maxY := 0
	for _, line := range lines {
		maxX = max(maxX, line.x1)
		maxX = max(maxX, line.x2)
		maxY = max(maxY, line.y1)
		maxY = max(maxY, line.y2)
	}

	vents := make([][]int, maxY+1)
	for i := range vents {
		vents[i] = make([]int, maxX+1)
	}
	return vents
}

func printVents(vents [][]int) {
	for _, vent := range vents {
		for _, v := range vent {
			if v == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", v)
			}
		}
		fmt.Printf("\n")
	}
}

func drawVerticalLines(vents [][]int, l line) {
	if l.x1 == l.x2 {
		minY := min(l.y1, l.y2)
		maxY := max(l.y1, l.y2)
		for y := minY; y <= maxY; y++ {
			vents[y][l.x1]++
		}
	}
}

func drawHorizontalLines(vents [][]int, l line) {
	if l.y1 == l.y2 {
		minX := min(l.x1, l.x2)
		maxX := max(l.x1, l.x2)
		for x := minX; x <= maxX; x++ {
			vents[l.y1][x]++
		}
	}
}

func drawDiagonalLines(vents [][]int, l line) {
	if abs(l.x1-l.x2) == abs(l.y1-l.y2) {
		var x1 int
		// var x2 int
		var y1 int
		var y2 int
		if l.x1 < l.x2 {
			// left to right
			x1 = l.x1
			y1 = l.y1
			// x2 = l.x2
			y2 = l.y2
		} else {
			x1 = l.x2
			y1 = l.y2
			// x2 = l.x1
			y2 = l.y1
		}

		difference := y1 - y2
		if difference >= 0 {
			//increasing
			for j := 0; j <= difference; j++ {
				vents[y1-j][x1+j]++
			}
		} else {
			// decreasing
			for j := 0; j <= abs(difference); j++ {
				vents[y1+j][x1+j]++
			}
		}
	}
}

func countOverlaps(vents [][]int) int {
	overlaps := 0
	for _, vent := range vents {
		for _, v := range vent {
			if v > 1 {
				overlaps++
			}
		}
	}
	return overlaps
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
