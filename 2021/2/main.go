package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	moves := readInput("./input.txt")
	solveOne(moves)
	solveTwo(moves)
}

func solveOne(moves []string) {
	pos := 0
	depth := 0
	for _, move := range moves {
		re := regexp.MustCompile("(?P<direction>\\w*) (?P<amount>\\d)")
		match := re.FindStringSubmatch(move)
		direction := match[1]
		amount, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "forward":
			pos += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	fmt.Printf("Part 1: Pos %d * depth %d = %d\n", pos, depth, pos*depth)
}

func solveTwo(moves []string) {
	pos := 0
	depth := 0
	aim := 0
	for _, move := range moves {
		re := regexp.MustCompile("(?P<direction>\\w*) (?P<amount>\\d)")
		match := re.FindStringSubmatch(move)
		direction := match[1]
		amount, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "forward":
			pos += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}
	fmt.Printf("Part 2: Pos %d * depth %d = %d\n", pos, depth, pos*depth)
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
