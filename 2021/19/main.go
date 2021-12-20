package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type coord struct {
	x int
	y int
	z int
}

type scanner struct {
	beacons []coord
}

type pattern struct {
	origin coord
	coords []coord
}

/*
1 0 0
0 0 -1
0 1 0
*/
func (s *scanner) rotateX(w int) {
	w = w % 4
	for i := 0; i < w; i++ {
		bs := []coord{}
		for _, b := range s.beacons {
			bs = append(bs, coord{
				x: b.x,
				y: -b.z,
				z: b.y,
			})
		}
		s.beacons = bs
	}
}

/*
0 0 1
0 1 0
-1 0 0
*/
func (s *scanner) rotateY(w int) {
	w = w % 4
	for i := 0; i < w; i++ {
		bs := []coord{}
		for _, b := range s.beacons {
			bs = append(bs, coord{
				x: b.z,
				y: b.y,
				z: -b.x,
			})
		}
		s.beacons = bs
	}
}

/*
0 -1 0
1 0 0
0 0 1
*/
func (s *scanner) rotateZ(w int) {
	w = w % 4
	for i := 0; i < w; i++ {
		bs := []coord{}
		for _, b := range s.beacons {
			bs = append(bs, coord{
				x: -b.y,
				y: b.x,
				z: b.z,
			})
		}
		s.beacons = bs
	}
}

/*
fix point in set A
calc relative positions to it

for every point in set B
calc relative positions to it

check if at least 12 points have the same coords
*/
func (s *scanner) isNeighbour(other scanner) *coord {
	for i := range s.beacons {
		p1 := pattern{
			origin: s.beacons[i],
		}
		p1.calibrate(append(s.clone().beacons[:i], s.clone().beacons[i+1:]...))
		for j := range other.beacons {
			p2 := pattern{
				origin: other.beacons[j],
			}
			p2.calibrate(append(other.clone().beacons[:j], other.clone().beacons[j+1:]...))
			if p1.fits(p2) {
				n := p1.origin.subtract(p2.origin)
				return &n
			}
		}
	}
	return nil
}

func (s *scanner) clone() scanner {
	var c scanner
	for _, b := range s.beacons {
		c.beacons = append(c.beacons, b.clone())
	}
	return c
}

func (p *pattern) calibrate(coords []coord) {
	for _, c := range coords {
		p.coords = append(p.coords, c.subtract(p.origin))
	}
}

func (p *pattern) fits(other pattern) bool {
	matches := 1
	for _, c1 := range p.coords {
		for _, c2 := range other.coords {
			if reflect.DeepEqual(c1, c2) {
				matches++
			}
		}
	}
	return matches >= 3
}

func (c *coord) add(other coord) coord {
	return coord{
		x: c.x + other.x,
		y: c.y + other.y,
		z: c.z + other.z,
	}
}

func (c *coord) subtract(other coord) coord {
	return coord{
		x: c.x - other.x,
		y: c.y - other.y,
		z: c.z - other.z,
	}
}

func (c *coord) clone() coord {
	return coord{
		x: c.x,
		y: c.y,
		z: c.z,
	}
}

func main() {
	input := readInput("./input.txt")

	i := 0
	scanners := []*scanner{}
	for i < len(input) {
		scanner := scanner{
			beacons: []coord{},
		}
		for i < len(input) {
			l := input[i]
			i++
			if strings.HasPrefix(l, "---") {
				continue
			}
			if len(l) == 0 {
				break
			}
			var c coord
			fmt.Sscanf(l, "%d,%d,%d", &c.x, &c.y, &c.z)
			scanner.beacons = append(scanner.beacons, c)
		}
		scanners = append(scanners, &scanner)
	}

	solve(scanners)
}

func solve(scanners []*scanner) {
	sr := calcScannerRelPos(scanners)
	fmt.Println(sr)
	sa := calcScannerAbsPos(sr, len(scanners))
	fmt.Println(sa)

	totalBeacons := countBeacons(scanners, sa)
	fmt.Printf("Part 1: There are %d beacons in total\n", totalBeacons)

	maxManhattenDistance := calcMaxManhattenDistance(sa)
	fmt.Printf("Part 2: The largest Manhatten distance between any two scanners is %d\n", maxManhattenDistance)
}

func calcScannerRelPos(scanners []*scanner) map[string]coord {
	scannerRelatives := map[string]coord{}
	for i, s1 := range scanners {
		for j, s2 := range scanners {
			if i >= j {
				continue
			}
			fmt.Printf("%d %d\n", i, j)
			c := isNeighbourInAnyRotation(s1, s2)
			if c != nil {
				scannerRelatives[fmt.Sprintf("%d,%d", i, j)] = *c
			}
		}
	}
	return scannerRelatives
}

func isNeighbourInAnyRotation(s1, s2 *scanner) *coord {
	c1 := s1.clone()
	for k := 0; k < 4; k++ {
		for l := 0; l < 4; l++ {
			for m := 0; m < 4; m++ {
				c2 := s2.clone()
				c2.rotateX(k)
				c2.rotateY(l)
				c2.rotateZ(m)
				if c := c1.isNeighbour(c2); c != nil {
					s2.rotateX(k)
					s2.rotateY(l)
					s2.rotateZ(m)
					return c
				}
			}
		}
	}
	return nil
}

func calcScannerAbsPos(sr map[string]coord, limit int) map[int]coord {
	sa := map[int]coord{
		0: {0, 0, 0},
	}

	for len(sa) < limit {
		for s, c := range sr {
			var s1, s2 int
			fmt.Sscanf(s, "%d,%d", &s1, &s2)
			if ref, ok := sa[s1]; ok {
				if _, ok := sa[s2]; ok {
					continue
				}
				sa[s2] = c.add(ref)
			}
		}
	}

	return sa
}

func countBeacons(scanners []*scanner, scannerPos map[int]coord) int {
	seen := map[coord]bool{}
	for i, s := range scanners {
		for _, b := range s.beacons {
			pos := scannerPos[i]
			c := pos.add(b)
			seen[c] = true
		}
	}
	return len(seen)
}

func calcMaxManhattenDistance(scannerPos map[int]coord) int {
	maxDistance := 0
	for i, pos1 := range scannerPos {
		for j, pos2 := range scannerPos {
			if i >= j {
				continue
			}
			manhattenDistance := calcManhattenDistance(pos1, pos2)
			maxDistance = max(maxDistance, manhattenDistance)
		}
	}
	return maxDistance
}

func calcManhattenDistance(c1, c2 coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y) + abs(c1.z-c2.z)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
