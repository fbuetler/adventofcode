package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type snailfishNumber struct {
	parent *snailfishNumber

	left  *snailfishNumber
	right *snailfishNumber

	value int
}

func (sf *snailfishNumber) sumMagnitude() int {
	sum := 0
	if sf.left != nil {
		sum += sf.left.sumMagnitude() * 3
	}
	if sf.right != nil {
		sum += sf.right.sumMagnitude() * 2
	}
	if sf.left == nil && sf.right == nil {
		sum += sf.value
	}
	return sum
}

func main() {
	input := readInput("./input.txt")
	solve(input)
}

func solve(numbers []string) {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum := add(result, numbers[i])
		result = reduce(sum)
	}

	sf, _ := parseSnailfischNumber(result, 0)
	fmt.Printf("Part 1: %d\n", sf.sumMagnitude())

	maxMagnitude := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}

			s := add(numbers[i], numbers[j])
			r := reduce(s)
			sf, _ := parseSnailfischNumber(r, 0)
			m := sf.sumMagnitude()
			if m > maxMagnitude {
				maxMagnitude = m
			}

			s = add(numbers[j], numbers[i])
			r = reduce(s)
			sf, _ = parseSnailfischNumber(r, 0)
			m = sf.sumMagnitude()
			if m > maxMagnitude {
				maxMagnitude = m
			}
		}
	}
	fmt.Printf("Part 2: %d\n", maxMagnitude)
}

func leftMostPairWithDepthAndNeighbours(number string, depth int) (int, int, int, int) {
	// find pair
	currentDepth := 0
	pairStartIndex := 0
	pairEndIndex := 0
	deapthReached := false
	for i, c := range number {
		if c == '[' {
			currentDepth++
		} else if c == ']' {
			currentDepth--
		}
		if currentDepth == depth {
			pairStartIndex = i
			deapthReached = true
			break
		}
	}

	if !deapthReached {
		return -1, -1, -1, -1
	}

	for i := pairStartIndex; i < len(number); i++ {
		if number[i] == '[' {
			pairStartIndex = i
		}
		if number[i] == ']' {
			pairEndIndex = i + 1
			break
		}
	}
	// find left neighbour
	lnIndex := -1
	for i := pairStartIndex; i >= 0; i-- {
		if _, err := strconv.Atoi(string(number[i])); err != nil {
			continue
		}
		for number[i] != '[' && number[i] != ',' {
			i--
		}
		lnIndex = i + 1
		break
	}
	// find right neighbour
	rnIndex := -1
	for i := pairEndIndex; i < len(number); i++ {
		if _, err := strconv.Atoi(string(number[i])); err == nil {
			rnIndex = i
			break
		}
	}

	return lnIndex, pairStartIndex, pairEndIndex, rnIndex
}

func leftMostNumberGreaterThan(number string, limit int) int {
	numberIndex := -1

	i := 0
	for i < len(number) {
		// find start of number
		if _, err := strconv.Atoi(string(number[i])); err != nil {
			i++
			continue
		}

		n, _ := numberAt(number, i)
		if n > limit {
			numberIndex = i
			break
		}
		i += len(fmt.Sprintf("%d", n))
	}
	return numberIndex
}

func numberAt(number string, i int) (int, int) {
	j := i
	for number[j] != ',' && number[j] != ']' {
		j++
	}

	n, err := strconv.Atoi(string(number[i:j]))
	if err != nil {
		return -1, 0
	}
	return n, j - i
}

func add(n1, n2 string) string {
	return fmt.Sprintf("[%s,%s]", n1, n2)
}

func reduce(number string) string {
	// fmt.Printf("after addition: %s\n", number)
	for {
		_, startInx, endIdx, _ := leftMostPairWithDepthAndNeighbours(number, 5)
		if startInx != -1 && endIdx != -1 {
			prevNumber := number
			number = explode(number)
			// fmt.Printf("after explode: %s\n", number)
			if prevNumber != number {
				continue
			}
		}

		limit := 10
		nIndx := leftMostNumberGreaterThan(number, limit-1)
		ruleApplied := false
		if nIndx != -1 {
			number = split(number)
			// fmt.Printf("after split: %s\n", number)
			ruleApplied = true
		}
		if !ruleApplied {
			break
		}
	}
	return number
}

func explode(number string) string {
	lnIdx, startInx, endIdx, rnIdx := leftMostPairWithDepthAndNeighbours(number, 5)

	var lv, rv int
	fmt.Sscanf(number[startInx:endIdx], "[%d,%d]", &lv, &rv)

	if rnIdx != -1 {
		rn, len := numberAt(number, rnIdx)
		number = number[:rnIdx] + strconv.Itoa(rv+rn) + number[rnIdx+len:]
	}

	number = number[:startInx] + "0" + number[endIdx:]

	if lnIdx != -1 {
		ln, len := numberAt(number, lnIdx)
		number = number[:lnIdx] + strconv.Itoa(lv+ln) + number[lnIdx+len:]
	}

	return number
}

func split(number string) string {
	nIdx := leftMostNumberGreaterThan(number, 9)
	if nIdx == -1 {
		return number
	}

	n, len := numberAt(number, nIdx)
	split := fmt.Sprintf("[%d,%d]", int(math.Floor(float64(n)/2)), int(math.Ceil(float64(n)/2)))

	return number[:nIdx] + split + number[nIdx+len:]
}

func parseSnailfischNumber(str string, curr int) (*snailfishNumber, int) {
	sf := &snailfishNumber{}
	curr++ // consume [

	// left pair elment
	if str[curr] == '[' {
		subsf, subcurr := parseSnailfischNumber(str, curr)
		subsf.parent = sf
		sf.left = subsf
		curr = subcurr
	} else {
		n, _ := strconv.Atoi(string(str[curr]))
		sf.left = &snailfishNumber{value: n}
		curr++ // one digit number and comma
	}

	curr++

	// right pair elment
	if str[curr] == '[' {
		subsf, subcurr := parseSnailfischNumber(str, curr)
		subsf.parent = sf
		sf.right = subsf
		curr = subcurr
	} else {
		n, _ := strconv.Atoi(string(str[curr]))
		sf.right = &snailfishNumber{value: n}
		curr++ // one digit number and comma
	}

	curr++ // consume [

	return sf, curr
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
