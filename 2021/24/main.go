package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	w int
	z int
}
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].w < p[j].w
}

type bound struct {
	min int
	max int
}

type stateInterval map[string]bound

type state struct {
	vars map[string]int
	min  int
	max  int
}

var consts = []map[string]int{}

func (s state) apply(leftOp, rightOp string, op func(int, int) int) {
	var right int
	if d, err := strconv.Atoi(rightOp); err == nil {
		right = d
	} else {
		right = s.vars[rightOp]
	}
	s.vars[leftOp] = op(s.vars[leftOp], right)
}

func (s state) copy() state {
	c := state{
		vars: map[string]int{},
	}
	for k, v := range s.vars {
		c.vars[k] = v
	}
	c.min = s.min
	c.max = s.max
	return c
}

func main() {
	input := readInput("./input.txt")
	// read constants that differ in each block
	for i := 0; i < 14; i++ {
		a, _ := strconv.Atoi(input[i*18+4][6:])
		b, _ := strconv.Atoi(input[i*18+5][6:])
		c, _ := strconv.Atoi(input[i*18+15][6:])
		m := map[string]int{
			"a": a,
			"b": b,
			"c": c,
		}
		consts = append(consts, m)
	}
	solve(input)
}

func solve(instrs []string) {
	minMonad, maxMonad := computeMonad()
	fmt.Printf("Part 1: %d\n", maxMonad)
	fmt.Printf("Part 2: %d\n", minMonad)
}

/*	approach 5:
	nextZ = z / a, 				iff z%26+b == w
	nextZ = (z/a)*26 + w + c, 	otherwise

	solving the formular from above for z we get:
		z = nextZ * a + r
		z = (nextZ - w - c) / 26 * a + r
	where r is the remainder of the integer division that was cut off

	we cannot decide which case to take, as it depends on z, therefore
	we just try both and check which one actually works
*/
func computeMonad() (int, int) {
	possibleInputs := make([]map[int]pairs, 14)
	zs := []int{0}
	for i := 13; i >= 0; i-- {
		a := consts[i]["a"]
		b := consts[i]["b"]
		c := consts[i]["c"]
		possibleInputs[i] = map[int]pairs{}
		prevZs := []int{}
		for w := 1; w <= 9; w++ {
			for _, z := range zs {
				for r := 0; r < a; r++ {
					// try case one
					prevZ := z*a + r
					if prevZ%26+b == w && prevZ/a == z {
						prevZs = append(prevZs, prevZ)
						possibleInputs[i][prevZ] = append(possibleInputs[i][prevZ], pair{w, z})
					}
					// try case two
					prevZ = (z-w-c)/26*a + r
					if prevZ%26+b != w && prevZ/a*26+w+c == z {
						prevZs = append(prevZs, prevZ)
						possibleInputs[i][prevZ] = append(possibleInputs[i][prevZ], pair{w, z})
					}
				}
			}
		}
		fmt.Printf("Backtracking block %d\n", i)
		zs = prevZs
	}

	// z has to be 0 at the start and the end
	_, minMonad := backtraceMonad(0, 0, 0, possibleInputs, false)
	_, maxMonad := backtraceMonad(0, 0, 0, possibleInputs, true)

	return minMonad, maxMonad
}

func backtraceMonad(block, z, monadPart int, possibleInputs []map[int]pairs, largest bool) (bool, int) {
	if block == 14 {
		return true, monadPart
	}
	if _, ok := possibleInputs[block][z]; !ok {
		return false, 0
	}

	pairs := possibleInputs[block][z]
	if largest {
		sort.Sort(sort.Reverse(pairs))
	} else {
		sort.Sort(pairs)
	}
	for _, pair := range pairs {
		m := monadPart + pair.w*int(math.Pow(10, 13-float64(block)))
		complete, monad := backtraceMonad(block+1, pair.z, m, possibleInputs, largest)
		if complete {
			return true, monad
		}
	}
	return false, 0
}

/*  approach 4:
if we analyse our code we observe the following:
* there are 14 blocks of 18 instructions each starting with an 'inp'
* all blocks have the same instructions in the same order
* the only differences between two blocks are some constants on line 5, 6 and 16
* x and y are always nulified before they are used, hence only z and w are interesting
* analysing a block yields the following formula:
	nextZ = z / a, 				iff z%26+b == w
	nextZ = (z/a)*26 + w + c, 	otherwise

inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w		inp w
mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0		mul x 0
add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z		add x z
mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26	mod x 26
div z 1		div z 1		div z 1		div z 1		div z 1		div z 26	div z 1		div z 26	div z 26	div z 1		div z 26	div z 26	div z 26	div z 26
add x 10	add x 12	add x 13	add x 13	add x 14	add x -2	add x 11	add x -15	add x -10	add x 10	add x -10	add x -4	add x -1	add x -1
eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w		eql x w
eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0		eql x 0
mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0
add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25	add y 25
mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x
add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1		add y 1
mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y		mul z y
mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0		mul y 0
add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w		add y w
add y 0		add y 6		add y 4		add y 2		add y 9		add y 1		add y 10	add y 6		add y 4		add y 6		add y 3		add y 9		add y 1	5	add y 5
mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x		mul y x
add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y		add z y

*/
func isMonad(monad int) bool {
	z := 0
	for i := 0; i < 14; i++ {
		w, _ := strconv.Atoi(string(strconv.Itoa(monad)[i]))
		a := consts[i]["a"]
		b := consts[i]["b"]
		c := consts[i]["c"]

		if z%26+b == w {
			z = z / a
		} else {
			z = (z/a)*26 + w + c
		}
	}
	return z == 0
}

// approach 3: analyse keeps track of all possible states the programm can have
func analyse(instrs []string, initialState state) []state {
	states := []state{initialState}

	for _, instr := range instrs {
		// fmt.Printf("Instr %d\n", i)

		instrParts := strings.Split(instr, " ")
		op := instrParts[0]
		left := instrParts[1]
		var right string
		if op != "inp" {
			right = instrParts[2]
		}

		newStates := []state{}
		for _, state := range states {
			// apply instruction to state
			switch op {
			case "inp":
				for i := 1; i <= 9; i++ {
					newState := state.copy()
					newState.vars[left] = i
					newState.min = newState.min*10 + i
					newState.max = newState.max*10 + i
					if idx := containsState(states, newState); idx >= 0 {
						states[idx].min = min(states[idx].min, newState.min)
						states[idx].max = max(states[idx].max, newState.max)
					} else {
						newStates = append(newStates, newState)
					}
				}
			case "add":
				state.apply(left, right, func(a, b int) int {
					return a + b
				})
			case "mul":
				state.apply(left, right, func(a, b int) int {
					return a * b
				})
			case "div":
				state.apply(left, right, func(a, b int) int {
					return a / b
				})
			case "mod":
				state.apply(left, right, func(a, b int) int {
					return a % b
				})
			case "eql":
				state.apply(left, right, func(a, b int) int {
					if a == b {
						return 1
					}
					return 0
				})
			}
		}
		if len(newStates) != 0 {
			states = newStates
			// fmt.Printf("States: %d\n", len(states))
		}
	}

	return states
}

// approach 2: analyseIntervals analyses the program and computes the interval each variable can have at any given point
func analyseIntervals(instrs []string) stateInterval {
	// initial state
	state := stateInterval{
		"w": bound{0, 0},
		"x": bound{0, 0},
		"y": bound{0, 0},
		"z": bound{0, 0},
	}

	for _, instr := range instrs {
		instrParts := strings.Split(instr, " ")
		op := instrParts[0]
		leftVar := instrParts[1]

		left := state[leftVar]
		var right bound
		if op != "inp" {
			if d, err := strconv.Atoi(instrParts[2]); err == nil {
				right = bound{d, d}
			} else {
				right = state[instrParts[2]]
			}
		}

		// apply instruction to state
		switch op {
		case "inp":
			state[leftVar] = bound{1, 9}
		case "add":
			state[leftVar] = bound{left.min + right.min, left.max + right.max}
		case "mul":
			a1b1 := left.min * right.min
			a1b2 := left.min * right.max
			a2b1 := left.max * right.min
			a2b2 := left.max * right.max
			min := min(a1b1, min(a1b2, min(a2b1, a2b2)))
			max := max(a1b1, max(a1b2, max(a2b1, a2b2)))
			state[leftVar] = bound{min, max}
		case "div":
			// luckily 0 is not part of an interval
			a1b1 := left.min / right.max
			a1b2 := left.min / right.min
			a2b1 := left.max / right.max
			a2b2 := left.max / right.min
			min := min(a1b1, min(a1b2, min(a2b1, a2b2)))
			max := max(a1b1, max(a1b2, max(a2b1, a2b2)))
			state[leftVar] = bound{min, max}
		case "mod":
			state[leftVar] = bound{0, right.max - 1}
		case "eql":
			var b bound
			if left.max < right.min || left.min > right.max {
				b = bound{0, 0}

			} else if left.min == right.min && left.max == right.max {
				b = bound{1, 1}
			} else {
				b = bound{0, 1}
			}
			state[leftVar] = b
		}
	}

	return state
}

// approach 1: proccess executes thes program with a given monad
func process(instrs []string, monad int) bool {
	mi := 0
	m := strings.Split(strconv.Itoa(monad), "")

	w, x, y, z := 0, 0, 0, 0
	for _, instr := range instrs {
		instrParts := strings.Split(instr, " ")

		var left, right *int
		if instrParts[1] == "w" {
			left = &w
		} else if instrParts[1] == "x" {
			left = &x
		} else if instrParts[1] == "y" {
			left = &y
		} else if instrParts[1] == "z" {
			left = &z
		} else {
			l, _ := strconv.Atoi(instrParts[1])
			left = &l
		}

		if instrParts[0] != "inp" {
			if instrParts[2] == "w" {
				right = &w
			} else if instrParts[2] == "x" {
				right = &x
			} else if instrParts[2] == "y" {
				right = &y
			} else if instrParts[2] == "z" {
				right = &z
			} else {
				r, _ := strconv.Atoi(instrParts[2])
				right = &r
			}
		}

		switch instrParts[0] {
		case "inp":
			a, _ := strconv.Atoi(m[mi])
			mi++
			*left = a
		case "add":
			*left = *left + *right
		case "mul":
			*left = *left * *right
		case "div":
			*left = *left / *right
		case "mod":
			*left = *left % *right
		case "eql":
			if *left == *right {
				*left = 1
			} else {
				*left = 0
			}
		}
	}

	fmt.Printf("w=%d, x=%d, y=%d, z=%d\n", w, x, y, z)
	return z == 0
}

// some helpers down here
func containsState(states []state, el state) int {
	for i, state := range states {
		allMatch := true
		for _, variable := range []string{"w", "x", "y", "z"} {
			allMatch = allMatch && (state.vars[variable] == el.vars[variable])
		}
		if allMatch {
			return i
		}
	}
	return -1
}

func contains(arr []string, el string) bool {
	for _, a := range arr {
		if a == el {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
