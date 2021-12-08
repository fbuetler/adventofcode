package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func main() {
	input := readInput("./input.txt")
	signals := [][]string{}
	outputs := [][]string{}
	for _, in := range input {
		segments := strings.Split(in, "|")
		s := strings.Split(strings.Trim(segments[0], " "), " ")
		o := strings.Split(strings.Trim(segments[1], " "), " ")
		signals = append(signals, s)
		outputs = append(outputs, o)
	}
	solveOne(signals, outputs)
	solveTwo(signals, outputs)
}

func solveOne(signals [][]string, outputs [][]string) {
	appearances := 0
	for _, outs := range outputs {
		for _, o := range outs {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				appearances++
			}
		}
	}

	fmt.Printf("Part 1: Digits 1, 4, 7 and 8 appear %d times\n", appearances)
}

func solveTwo(signals [][]string, outputs [][]string) {
	sum := 0
	for i, sigs := range signals {
		digits := calculateDigits(sigs)

		output := 0
		for pos, o := range outputs[i] {
			digit := index(digits, sortString(o))
			output += pow(10, 4-pos-1) * digit
		}
		sum += output
	}

	fmt.Printf("Part 2: The sum of all output values is %d\n", sum)
}

func calculateDigits(signals []string) []string {
	digits := make([]string, 10)

	// determine 1, 4, 7, 8 (unique lengths)
	for _, s := range signals {
		if len(s) == 2 {
			digits[1] = sortString(s)
			continue
		}
		if len(s) == 4 {
			digits[4] = sortString(s)
			continue
		}
		if len(s) == 3 {
			digits[7] = sortString(s)
			continue
		}
		if len(s) == 7 {
			digits[8] = sortString(s)
			continue
		}
	}

	corner := strings.Join(difference(digits[4], digits[1]), "")

	// determine 2, 3, 5 (length 5)
	for _, s := range signals {
		if len(s) != 5 {
			continue
		}
		if len(difference(s, digits[1])) == 3 {
			digits[3] = sortString(s)
			continue
		}
		if len(difference(s, corner)) == 3 {
			digits[5] = sortString(s)
			continue
		}
		digits[2] = sortString(s)
	}

	// determine 0, 6, 9 (length 6)
	for _, s := range signals {
		if len(s) != 6 {
			continue
		}
		if len(difference(s, digits[4])) == 2 {
			digits[9] = sortString(s)
			continue
		}
		if len(difference(s, corner)) == 4 {
			digits[6] = sortString(s)
			continue
		}
		digits[0] = sortString(s)
	}

	return digits
}

// return what chars are in s1 and not in s2 (s1 - s2)
func difference(s1 string, s2 string) []string {
	diff := []string{}
	for _, r := range s1 {
		s := string(r)
		if !strings.Contains(s2, s) {
			diff = append(diff, s)
		}
	}
	return diff
}

func index(ss []string, str string) int {
	for i, s := range ss {
		if s == str {
			return i
		}
	}
	// fmt.Printf("not found %s in %v\n", str, ss)
	return -1
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func pow(a int, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
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
