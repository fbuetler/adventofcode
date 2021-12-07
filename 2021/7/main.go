package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./input.txt")
	positions := []int{}
	for _, s := range strings.Split(input[0], ",") {
		position, _ := strconv.Atoi(s)
		positions = append(positions, position)
	}
	solveOne(positions)
	solveTwo(positions)
}

func solveOne(positions []int) {
	minimalFuel, optimalPos := fuelNeeded(positions, func(start int, end int) int {
		return abs(start - end)
	})
	fmt.Printf("Part 1: Optimal position is at %d with %d fuel used\n", optimalPos, minimalFuel)
}

func solveTwo(positions []int) {
	minimalFuel, optimalPos := fuelNeeded(positions, func(start int, end int) int {
		n := abs(start - end)
		return (n*n + n) / 2
	})
	fmt.Printf("Part 2: Optimal position is at %d with %d fuel used\n", optimalPos, minimalFuel)
}

func fuelNeeded(positions []int, step func(int, int) int) (int, int) {
	max := max(positions)

	fuel := make([]int, max+1)
	for i := 0; i < len(fuel); i++ {
		f := 0
		for _, pos := range positions {
			f += step(i, pos)
		}
		fuel[i] = f
	}
	return findOptimum(fuel)
}

func findOptimum(fuel []int) (int, int) {
	minimalFuel := math.MaxInt64
	optimalPos := 0
	for i, f := range fuel {
		if f < minimalFuel {
			minimalFuel = f
			optimalPos = i
		}
	}
	return minimalFuel, optimalPos
}

func max(as []int) int {
	max := 0
	for _, a := range as {
		if a > max {
			max = a
		}
	}
	return max
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
