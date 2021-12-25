package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("./input.txt")

	grid := [][]string{}
	for _, l := range input {
		g := []string{}
		for _, c := range l {
			g = append(g, string(c))
		}
		grid = append(grid, g)
	}

	solveOne(grid)
}

func solveOne(grid [][]string) {
	// fmt.Println("Initial state:")
	// print(grid)

	step := 0
	for {
		moves := 0
		// move east
		newGrid := copy(grid)
		for i, l := range grid {
			for j, c := range l {
				if c != ">" {
					continue
				}
				right := (j + 1) % len(l)
				if grid[i][right] == "." {
					newGrid[i][j] = "."
					newGrid[i][right] = ">"
					moves++
				}
			}
		}
		grid = newGrid

		// move south
		newGrid = copy(grid)
		for i, l := range grid {
			for j, c := range l {
				if c != "v" {
					continue
				}
				down := (i + 1) % len(grid)
				if grid[down][j] == "." {
					newGrid[i][j] = "."
					newGrid[down][j] = "v"
					moves++
				}
			}
		}
		grid = newGrid

		step++
		// fmt.Printf("After step %d:\n", step)
		// print(grid)
		if moves == 0 {
			break
		}
	}
	fmt.Printf("Part 1: %d\n", step)
}

func print(grid [][]string) {
	for _, l := range grid {
		for _, c := range l {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func copy(grid [][]string) [][]string {
	c := [][]string{}
	for _, l := range grid {
		cl := []string{}
		for _, el := range l {
			cl = append(cl, el)
		}
		c = append(c, cl)
	}
	return c
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
