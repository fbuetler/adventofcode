package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./input.txt")

	energies := [][]int{}
	for _, line := range input {
		ls := strings.Split(line, "")
		es := []int{}
		for _, l := range ls {
			e, _ := strconv.Atoi(l)
			es = append(es, e)
		}
		energies = append(energies, es)
	}

	solve(energies)
}

func solve(energies [][]int) {
	solved := 0
	step := 0
	flashes := 0
	for true {
		increaseEnergyLevel(energies)

		flashed := make([][]bool, len(energies))
		for i := range flashed {
			flashed[i] = make([]bool, len(energies[0]))
		}
		for true {
			f := countFlashes(energies, flashed)
			if f == 0 {
				break
			}
			if allFlashed(energies) {
				fmt.Printf("Part 2: All octopuses flash after %d steps\n", step)
				solved++
			}
			flashes += f
		}
		if step == 100 {
			fmt.Printf("Part 1: There are %d flashes after %d steps\n", flashes, step)
			solved++
		}
		if solved == 2 {
			return
		}
		step++
	}
}

func increaseEnergyLevel(energies [][]int) {
	for i, es := range energies {
		for j := range es {
			energies[i][j]++
		}
	}
}

func countFlashes(energies [][]int, flashed [][]bool) int {
	flashes := 0
	for i, es := range energies {
		for j := range es {
			if energies[i][j] <= 9 {
				continue
			}

			flashes++
			energies[i][j] = 0
			flashed[i][j] = true

			if i > 0 {
				energies[i-1][j] += flashRessonance(i-1, j, flashed)
				if j > 0 {
					energies[i-1][j-1] += flashRessonance(i-1, j-1, flashed)
				}
				if j < len(es)-1 {
					energies[i-1][j+1] += flashRessonance(i-1, j+1, flashed)
				}
			}
			if i < len(energies)-1 {
				energies[i+1][j] += flashRessonance(i+1, j, flashed)
				if j > 0 {
					energies[i+1][j-1] += flashRessonance(i+1, j-1, flashed)
				}
				if j < len(es)-1 {
					energies[i+1][j+1] += flashRessonance(i+1, j+1, flashed)
				}
			}
			if j > 0 {
				energies[i][j-1] += flashRessonance(i, j-1, flashed)
			}
			if j < len(es)-1 {
				energies[i][j+1] += flashRessonance(i, j+1, flashed)
			}
		}
	}
	return flashes
}

func flashRessonance(i, j int, flashed [][]bool) int {
	if flashed[i][j] {
		return 0
	}
	return 1
}

func allFlashed(energies [][]int) bool {
	for i, es := range energies {
		for j := range es {
			if energies[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func print(energies [][]int) {
	for i, es := range energies {
		for j := range es {
			fmt.Printf("%d", energies[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
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
