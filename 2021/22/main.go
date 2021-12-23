package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

type reactor struct {
	switchedOn []region
}

type region struct {
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

type action struct {
	switchOn bool
	region   region
}

func (rc *reactor) isEmpty() bool {
	return len(rc.switchedOn) == 0
}

func (rc *reactor) add(rgs []region) {
	for _, rg := range rgs {
		rc.switchedOn = append(rc.switchedOn, rg)
	}
}

func (rc *reactor) remove(rgs []region) {
	for _, rg := range rgs {
		for i := len(rc.switchedOn) - 1; i >= 0; i-- {
			if reflect.DeepEqual(rc.switchedOn[i], rg) {
				rc.switchedOn = append(rc.switchedOn[:i], rc.switchedOn[i+1:]...)
			}
		}
	}

}

func (rc *reactor) spans() int {
	sum := 0
	for _, r := range rc.switchedOn {
		sum += (r.maxX - r.minX + 1) * (r.maxY - r.minY + 1) * (r.maxZ - r.minZ + 1)
	}
	return sum
}

func (rg *region) overlap(other region) *region {
	// check if there is an overlap
	if (rg.maxX < other.minX || rg.minX > other.maxX) ||
		(rg.maxY < other.minY || rg.minY > other.maxY) ||
		(rg.maxZ < other.minZ || rg.minZ > other.maxZ) {
		return nil
	}

	// compure overlap
	return &region{
		minX: max(rg.minX, other.minX),
		maxX: min(rg.maxX, other.maxX),
		minY: max(rg.minY, other.minY),
		maxY: min(rg.maxY, other.maxY),
		minZ: max(rg.minZ, other.minZ),
		maxZ: min(rg.maxZ, other.maxZ),
	}
}

func (rg *region) split(overlap region) []region {
	newRegions := []region{}
	/*
		the overlap is removed from the region by
		splitting the region up in multiple subregions
	*/
	if rg.minX != overlap.minX {
		newRegions = append(newRegions, region{rg.minX, overlap.minX - 1, rg.minY, rg.maxY, rg.minZ, rg.maxZ})
	}
	if rg.maxX != overlap.maxX {
		newRegions = append(newRegions, region{overlap.maxX + 1, rg.maxX, rg.minY, rg.maxY, rg.minZ, rg.maxZ})
	}
	if rg.minY != overlap.minY {
		newRegions = append(newRegions, region{overlap.minX, overlap.maxX, rg.minY, overlap.minY - 1, rg.minZ, rg.maxZ})
	}
	if rg.maxY != overlap.maxY {
		newRegions = append(newRegions, region{overlap.minX, overlap.maxX, overlap.maxY + 1, rg.maxY, rg.minZ, rg.maxZ})
	}
	if rg.minZ != overlap.minZ {
		newRegions = append(newRegions, region{overlap.minX, overlap.maxX, overlap.minY, overlap.maxY, rg.minZ, overlap.minZ - 1})
	}
	if rg.maxZ != overlap.maxZ {
		newRegions = append(newRegions, region{overlap.minX, overlap.maxX, overlap.minY, overlap.maxY, overlap.maxZ + 1, rg.maxZ})
	}

	return newRegions
}

func main() {
	input := readInput("./input.txt")

	var switchOn string
	var minX, maxX, minY, maxY, minZ, maxZ int
	actions := []action{}
	for _, l := range input {
		fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &switchOn, &minX, &maxX, &minY, &maxY, &minZ, &maxZ)
		actions = append(actions, action{
			switchOn == "on",
			region{minX,
				maxX,
				minY,
				maxY,
				minZ,
				maxZ,
			},
		})
	}

	solveOne(actions)
	solveTwo(actions)
}

func solveOne(actions []action) {
	size := 101
	cube := [101][101][101]bool{}
	for _, a := range actions {
		for i := max(a.region.minX, -50); i <= min(a.region.maxX, 50); i++ {
			for j := max(a.region.minY, -50); j <= min(a.region.maxY, 50); j++ {
				for k := max(a.region.minZ, -50); k <= min(a.region.maxZ, 50); k++ {
					cube[i+50][j+50][k+50] = a.switchOn
				}
			}
		}
	}

	turnedOn := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				if cube[i][j][k] {
					turnedOn++
				}
			}
		}
	}
	fmt.Printf("Part 1: %d cubes are turned on\n", turnedOn)
}

func solveTwo(actions []action) {
	rc := reactor{}
	for _, a := range actions {
		/*
			* a region that is switched on is always added
			* a region that is switch off is never added
			* if there is an overlap with existing regions,
			  then we always remove the overlap
			  (by splitting the existing region)
		*/
		toAdd := []region{}
		toRemove := []region{}
		if a.switchOn {
			toAdd = append(toAdd, a.region)
		}

		if rc.isEmpty() {
			rc.add(toAdd)
			continue
		}

		for _, rgOn := range rc.switchedOn {
			overlap := rgOn.overlap(a.region)
			if overlap == nil {
				continue
			}
			toRemove = append(toRemove, rgOn)
			splits := rgOn.split(*overlap)
			toAdd = append(toAdd, splits...)
		}

		rc.remove(toRemove)
		rc.add(toAdd)
	}

	turnedOn := rc.spans()
	fmt.Printf("Part 2: %d cubes are turned on\n", turnedOn)
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
