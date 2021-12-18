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

func (sf *snailfishNumber) toString() string {
	str := ""
	if sf.left != nil {
		str += "["
		str += sf.left.toString()
		str += ","
	}
	if sf.right != nil {
		str += sf.right.toString()
		str += "]"
	}
	if sf.left == nil && sf.right == nil {
		str += strconv.Itoa(sf.value)
	}
	return str
}

func (sf *snailfishNumber) add(other *snailfishNumber) *snailfishNumber {
	sum := &snailfishNumber{
		left:  sf,
		right: other,
	}
	sum.left.parent = sum
	sum.right.parent = sum
	sum.reduce()
	return sum
}

func (sf *snailfishNumber) reduce() {
	// fmt.Printf("after addition: %s\n", sf.toString())
	for {
		changes := sf.explode(0)
		if changes != 0 {
			// fmt.Printf("after explode: %s\n", sf.toString())
			continue
		}

		changes = sf.split()
		if changes == 0 {
			break
		}
		// fmt.Printf("after split: %s\n", sf.toString())
	}
}

func (sf *snailfishNumber) explode(depth int) int {
	if depth == 5 {
		parent := sf.parent
		ln := parent.findLeftNeighbour()
		if ln != nil {
			ln.value += parent.left.value
		}
		rn := parent.findRightNeighbour()
		if rn != nil {
			rn.value += parent.right.value
		}

		parent.left = nil
		parent.right = nil
		parent.value = 0
		return 1
	}

	if sf.left == nil {
		return 0
	}
	changes := sf.left.explode(depth + 1)
	if changes != 0 || sf.right == nil {
		return changes
	}
	return sf.right.explode(depth + 1)
}

func (sf *snailfishNumber) findLeftNeighbour() *snailfishNumber {
	prev := sf
	for number := sf.parent; number != nil; {
		if number.right == prev {
			prev = number
			number = number.left
			continue
		}
		if number.left == prev {
			if sf.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.right
	}

	return nil
}

func (sf *snailfishNumber) findRightNeighbour() *snailfishNumber {
	prev := sf
	for number := sf.parent; number != nil; {
		if number.left == prev {
			prev = number
			number = number.right
			continue
		}
		if number.right == prev {
			if sf.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.left
	}
	return nil
}

func (sf *snailfishNumber) split() int {
	if sf.left == nil && sf.right == nil {
		if sf.value < 10 {
			return 0
		}
		sf.left = &snailfishNumber{
			parent: sf,
			value:  int(math.Floor(float64(sf.value) / 2)),
		}
		sf.right = &snailfishNumber{
			parent: sf,
			value:  int(math.Ceil(float64(sf.value) / 2)),
		}
		sf.value = 0
		return 1
	}

	if sf.left == nil {
		return 0
	}
	changes := sf.left.split()
	if changes != 0 || sf.right == nil {
		return changes
	}
	return sf.right.split()
}

func (sf *snailfishNumber) magnitude() int {
	sum := 0
	if sf.left != nil {
		sum += sf.left.magnitude() * 3
	}
	if sf.right != nil {
		sum += sf.right.magnitude() * 2
	}
	if sf.left == nil && sf.right == nil {
		sum += sf.value
	}
	return sum
}

func (sf *snailfishNumber) clone() *snailfishNumber {
	clone := &snailfishNumber{}
	if sf.left != nil {
		clone.left = sf.left.clone()
		clone.left.parent = clone
	}
	if sf.right != nil {
		clone.right = sf.right.clone()
		clone.right.parent = clone
	}
	clone.value = sf.value
	return clone
}

func main() {
	input := readInput("./input.txt")

	sfs := []*snailfishNumber{}
	for _, l := range input {
		sf := parseSnailfischNumber(l)
		sfs = append(sfs, sf)
	}

	solve(sfs)
}

func solve(sfs []*snailfishNumber) {
	sum := sfs[0].clone()
	for i := 1; i < len(sfs); i++ {
		sum = sum.add(sfs[i].clone())
	}
	fmt.Printf("Part 1: The magnitude of the final sum is %d\n", sum.magnitude())

	mm := maxMagnitude(sfs)
	fmt.Printf("Part 2: Largest magnitude of any sum of two different snailfish numbers is %d\n", mm)
}

func parseSnailfischNumber(rawSfn string) *snailfishNumber {
	sf := &snailfishNumber{}
	root := sf

	for _, c := range rawSfn {
		switch c {
		case '[':
			sf.left = &snailfishNumber{parent: sf}
			sf.right = &snailfishNumber{parent: sf}
			sf = sf.left
		case ',':
			sf = sf.parent.right
		case ']':
			sf = sf.parent
		default:
			n, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			sf.value = n
		}
	}

	return root
}

func maxMagnitude(sfs []*snailfishNumber) int {
	maxMagnitude := 0
	for i := 0; i < len(sfs); i++ {
		for j := 0; j < len(sfs); j++ {
			if i == j {
				continue
			}

			m := sfs[i].clone().add(sfs[j].clone()).magnitude()
			if m > maxMagnitude {
				maxMagnitude = m
			}

			m = sfs[j].clone().add(sfs[i].clone()).magnitude()
			if m > maxMagnitude {
				maxMagnitude = m
			}
		}
	}
	return maxMagnitude
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
