package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Stack struct {
	els []string
}

func (s *Stack) Push(el string) {
	s.els = append(s.els, el)
}

func (s *Stack) Len() int {
	return len(s.els)
}

func (s *Stack) Pop() string {
	n := s.Len() - 1
	el := s.els[n]
	s.els = s.els[:n]
	return el
}

var syntaxErrorScore = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var autoCompletionScore = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var matchingClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func main() {
	input := readInput("./input.txt")
	solveOne(input)
	solveTwo(input)
}

func solveOne(lines []string) {
	score := 0
	for _, line := range lines {
		syntaxError, _ := findSyntaxError(line)
		score += syntaxErrorScore[syntaxError]
	}
	fmt.Printf("Part 1: The total syntax error score is %d\n", score)
}

func solveTwo(lines []string) {
	complScores := []int{}
	for _, line := range lines {
		syntaxError, s := findSyntaxError(line)
		if len(syntaxError) != 0 {
			continue
		}
		complScore := 0
		for s.Len() > 0 {
			open := s.Pop()
			complScore *= 5
			complScore += autoCompletionScore[matchingClose[open]]
		}
		complScores = append(complScores, complScore)
	}
	sort.Ints(complScores)
	middle := len(complScores) / 2
	fmt.Printf("Part 2: The middle score is %d\n", complScores[middle])
}

func findSyntaxError(line string) (string, *Stack) {
	chars := strings.Split(line, "")
	s := Stack{
		els: []string{},
	}
	for _, char := range chars {
		if _, ok := matchingClose[char]; ok {
			s.Push(char)
			continue
		}
		open := s.Pop()
		if char != matchingClose[open] {
			return char, nil
		}
	}
	return "", &s
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
