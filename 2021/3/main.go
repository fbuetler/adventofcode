package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput("./input.txt")
	solveOne(input)
	solveTwo(input)
}

func solveOne(input []string) {
	gamma := 0
	epsilon := 0
	factor := 1
	for i := len(input[0]) - 1; i >= 0; i-- {
		mostCommon := mostCommonAtPos(input, i)
		leastCommon := (mostCommon + 1) % 2

		gamma += mostCommon * factor
		epsilon += leastCommon * factor
		factor *= 2
	}

	fmt.Printf("Part 1: power consumption is %d\n", gamma*epsilon)
}

func solveTwo(input []string) {
	mcMatches := input
	lcMatches := input
	i := 0
	for len(mcMatches) > 1 || len(lcMatches) > 1 {
		mostCommon := mostCommonAtPos(mcMatches, i)
		leastCommon := (mostCommonAtPos(lcMatches, i) + 1) % 2

		mcMatches = filterMatches(mcMatches, i, mostCommon)
		lcMatches = filterMatches(lcMatches, i, leastCommon)

		i++
	}

	oxygen, err := strconv.ParseInt(mcMatches[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2, err := strconv.ParseInt(lcMatches[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 2: life support rating is %d\n", oxygen*co2)
}

func filterMatches(input []string, pos int, common int) []string {
	if len(input) <= 1 {
		return input
	}
	matches := []string{}
	for _, bits := range input {
		if string(bits[pos]) == strconv.Itoa(common) {
			matches = append(matches, bits)
		}
	}
	return matches
}

func mostCommonAtPos(input []string, pos int) int {
	zeroes := 0
	for _, bits := range input {
		if string(bits[pos]) == "0" {
			zeroes++
		}
	}
	mostCommon := 0
	if zeroes <= len(input)/2 {
		mostCommon = 1
	}
	return mostCommon
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
