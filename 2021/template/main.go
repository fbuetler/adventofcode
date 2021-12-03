package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("./input.txt")
	solveOne(input)
	solveTwo(input)
}

func solveOne(input []string) {

	fmt.Printf("Part 1: \n")
}

func solveTwo(input []string) {

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
