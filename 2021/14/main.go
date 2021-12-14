package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	input := readInput("./input.txt")

	template := input[0]
	rules := map[string]string{}
	for i := 2; i < len(input); i++ {
		var in string
		var out rune
		fmt.Sscanf(input[i], "%s -> %c", &in, &out)
		rules[string(in)] = string(out)
	}

	solveOne(template, rules)
	solveTwo(template, rules)
}

func solveOne(template string, rules map[string]string) {
	steps := 10
	difference := simulate(template, rules, steps)
	fmt.Printf("Part 1: Most common element minus least common element is %d\n", difference)
}

func solveTwo(template string, rules map[string]string) {
	steps := 40
	difference := calculate(template, rules, steps)
	fmt.Printf("Part 2: Most common element minus least common element is %d\n", difference)
}

func simulate(template string, rules map[string]string, steps int) int {
	for i := 1; i <= steps; i++ {
		nextTemplate := ""
		for j := 0; j < len(template)-1; j++ {
			l := string(template[j])
			r := string(template[j+1])
			o := rules[l+r]
			nextTemplate += l + o
		}
		nextTemplate += string(template[len(template)-1]) // fence post problem
		template = nextTemplate
	}

	occurs := map[string]int{}
	for _, c := range template {
		occurs[string(c)]++
	}

	return differenceOfExtrema(occurs)
}

func calculate(template string, rules map[string]string, steps int) int {
	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		l := string(template[i])
		r := string(template[i+1])
		pairs[l+r]++
	}

	for i := 1; i <= steps; i++ {
		nextPairs := map[string]int{}
		for pair, occurences := range pairs {
			l := string(pair[0])
			r := string(pair[1])
			o := rules[pair]
			nextPairs[l+o] += occurences
			nextPairs[o+r] += occurences
		}
		pairs = nextPairs
	}

	occurs := map[string]int{}
	for pair, occurences := range pairs {
		l := string(pair[0])
		occurs[l] += occurences
	}
	// fence post problem
	lastChar := string(template[len(template)-1])
	occurs[lastChar]++

	return differenceOfExtrema(occurs)
}

func differenceOfExtrema(occurs map[string]int) int {
	max := 0
	min := math.MaxInt
	for _, o := range occurs {
		if o > max {
			max = o
		}
		if o < min {
			min = o
		}
	}

	return max - min
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
