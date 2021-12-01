package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	depths := readInput("./input.txt")
	solveOne(depths)
	solveTwo(depths)
}

func solveOne(depths []int) {
	deeper := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			deeper++
		}
	}
	fmt.Printf("Part 1: Goes %d times deeper\n", deeper)
}

func solveTwo(depths []int) {
	deeper := 0
	for i := 1; i < len(depths)-2; i++ {
		if depths[i-1]+depths[i]+depths[i+1] < depths[i]+depths[i+1]+depths[i+2] {
			deeper++
		}
	}
	fmt.Printf("Part 2: Goes %d times deeper\n", deeper)
}

func readInput(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []int{}
	for scanner.Scan() {
		rawNumber := scanner.Text()
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
