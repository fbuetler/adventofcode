package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("./input.txt")

	var minX, maxX, minY, maxY int
	fmt.Sscanf(input[0], "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)

	solve(minX, maxX, minY, maxY)
}

func solve(minX, maxX, minY, maxY int) {
	x := 0
	y := 0
	highestY := 0
	hitCounter := 0
	for velX := 1; velX <= maxX; velX++ {
		for velY := minY; velY < 1000; velY++ {
			hy, hit := simulateShoot(x, y, velX, velY, minX, maxX, minY, maxY)
			if hit {
				hitCounter++
				if hy > highestY {
					highestY = hy
				}
			}
		}
	}

	fmt.Printf("Part 1: Highest y position reached and still hit the target is %d\n", highestY)
	fmt.Printf("Part 2: Target can be hit with %d different initial velocity values\n", hitCounter)
}

func simulateShoot(x, y, velX, velY, minX, maxX, minY, maxY int) (int, bool) {
	highestY := 0
	for !hitTarget(x, y, minX, maxX, minY, maxY) {
		if overshotTarget(x, y, minX, maxX, minY, maxY) {
			return 0, false
		}
		if targetUnreachable(x, minX, velX) {
			return 0, false
		}
		x, y, velX, velY = shootStep(x, y, velX, velY)
		if y > highestY {
			highestY = y
		}
	}
	return highestY, true
}

func shootStep(x, y, velX, velY int) (int, int, int, int) {
	x += velX
	y += velY
	if velX > 0 {
		velX--
	} else if velX < 0 {
		velX++
	}
	velY--
	return x, y, velX, velY
}

func hitTarget(x, y, minX, maxX, minY, maxY int) bool {
	return minX <= x && x <= maxX && minY <= y && y <= maxY
}

func overshotTarget(x, y, minX, maxX, minY, maxY int) bool {
	return maxX < x || y < minY
}

func targetUnreachable(x, minX, velX int) bool {
	return x < minX && velX == 0
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
