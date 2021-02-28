package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := conwayCubes(ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

var diffs = [3]int{-1, 0, 1}

func conwayCubes(input string, part int) int {
	activeNodes := parseInput(input)

	diffsW := []int{0}
	if part == 2 {
		diffsW = []int{-1, 0, 1}
	}

	for cycles := 0; cycles < 6; cycles++ {
		toCheck := map[[4]int]bool{}
		fmt.Printf("Simulating cycle: %v\n", cycles+1)

		for coord := range activeNodes {
			for _, dx := range diffs {
				for _, dy := range diffs {
					for _, dz := range diffs {
						for _, dw := range diffsW {
							toCheck[[4]int{
								coord[0] + dx,
								coord[1] + dy,
								coord[2] + dz,
								coord[3] + dw}] = true
						}
					}
				}
			}
		}

		fmt.Printf("Positions to check: %d\n", len(toCheck))
		nextState := map[[4]int]bool{}
		for coord := range toCheck {
			// check all neighbors around this coord
			var countNeighbors int
			for _, dx := range diffs {
				for _, dy := range diffs {
					for _, dz := range diffs {
						for _, dw := range diffsW {
							if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
								x, y, z, w := coord[0]+dx, coord[1]+dy, coord[2]+dz, coord[3]+dw
								neighCoord := [4]int{x, y, z, w}
								if isActive, ok := activeNodes[neighCoord]; ok && isActive {
									countNeighbors++
								}
							}
						}
					}
				}
			}

			if wasActive, ok := activeNodes[coord]; ok && wasActive {
				if countNeighbors == 2 || countNeighbors == 3 {
					nextState[coord] = true
				}
			} else {
				// inactive originally
				if countNeighbors == 3 {
					nextState[coord] = true
				}
			}
		}

		activeNodes = nextState
		fmt.Printf("Active cubes: %d\n", len(activeNodes))
	}

	// cubes after 6 cycles
	return len(activeNodes)
}

// this is not perfectly generalized because arrays in go have to be sized at compile
// time, and slices can't be used to map keys because they're not trivial to compare
// they could be compared by converting it into a string... but that's annoying
func parseInput(input string) map[[4]int]bool {
	setActiveNodes := map[[4]int]bool{}

	for i, line := range strings.Split(input, "\n") {
		for j, cell := range strings.Split(line, "") {

			if cell == "#" {
				// start z and w coords at zero
				n := [4]int{i, j, 0, 0}
				setActiveNodes[n] = true
			}
		}
	}
	return setActiveNodes
}

/*
ReadFile is a wrapper over io/ioutil.ReadFile but also determines the
dynamic absolute path to the file.
*/
func ReadFile(pathFromCaller string) string {
	// Docs: https://golang.org/pkg/runtime/#Caller
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	// parse directory with pathFromCaller (which could be relative to Directory)
	absolutePath := path.Join(path.Dir(filename), pathFromCaller)

	// read the entire file & return the byte slice as a string
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}
	// trim off new lines and tabs at end of input files
	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}
