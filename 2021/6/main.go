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
	ages := []int{}
	for _, a := range strings.Split(input[0], ",") {
		age, _ := strconv.Atoi(a)
		ages = append(ages, age)
	}

	solveOne(ages)
	solveTwo(ages)
}

func solveOne(ages []int) {
	days := 80
	fishes := calculate(ages, days)
	fmt.Printf("Part 1: There are %d lanternfishes after %d days\n", fishes, days)
}

func solveTwo(ages []int) {
	days := 256
	fishes := calculate(ages, days)
	fmt.Printf("Part 2: There are %d lanternfishes after %d days\n", fishes, days)
}

func simulate(ages []int, days int) int {
	// fmt.Printf("Initial state: %v\n", ages)
	for day := 1; day <= days; day++ {
		newAges := []int{}
		for i, age := range ages {
			if age == 0 {
				ages[i] = 6
				newAges = append(newAges, 8)
			} else {
				ages[i]--
			}
		}
		ages = append(ages, newAges...)
		// fmt.Printf("After %d day: %v\n", day, ages)
	}
	return len(ages)
}

func calculate(initialAges []int, days int) int {
	ages := make([]int, 9)
	for _, age := range initialAges {
		ages[age]++
	}
	// fmt.Printf("Initial state: %v\n", ages)
	for day := 1; day <= days; day++ {
		newAges := make([]int, 9)
		for i := 1; i < len(ages); i++ {
			newAges[i-1] = ages[i]
		}
		newAges[6] += ages[0]
		newAges[8] += ages[0]
		ages = newAges
		// fmt.Printf("After %d day: %v\n", day, ages)
	}
	totalFishes := 0
	for _, fishes := range ages {
		totalFishes += fishes
	}
	return totalFishes
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
