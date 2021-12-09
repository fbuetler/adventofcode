package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readInput("./input.txt")
	heatmap := [][]int{}
	for _, line := range input {
		heatmapRow := []int{}
		for _, r := range line {
			height, _ := strconv.Atoi(string(r))
			heatmapRow = append(heatmapRow, height)
		}
		heatmap = append(heatmap, heatmapRow)
	}
	solveOne(heatmap)
	solveTwo(heatmap)
}

func solveOne(heatmap [][]int) {
	risk := 0
	lowPoints := findLowPoints(heatmap)
	for _, coords := range lowPoints {
		risk += heatmap[coords[0]][coords[1]] + 1
	}
	fmt.Printf("Part 1: The sum of the risk levels of all low points is %d\n", risk)
}

func solveTwo(heatmap [][]int) {
	basinSizes := []int{}
	lowPoints := findLowPoints(heatmap)
	for _, lowPoint := range lowPoints {
		basin := findBasin(heatmap, lowPoint[0], lowPoint[1], [][]int{})
		basinSizes = append(basinSizes, len(basin))
	}

	size := 1
	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})
	for i := 0; i < len(basinSizes) && i < 3; i++ {
		size *= basinSizes[i]
	}
	fmt.Printf("Part 2: The product of the size of the three largest basins is %d\n", size)
}

func findLowPoints(heatmap [][]int) [][]int {
	lowPoints := [][]int{}
	for i, heatmapRow := range heatmap {
		for j, height := range heatmapRow {
			if i > 0 && heatmap[i-1][j] <= height {
				continue
			}
			if i < len(heatmap)-1 && heatmap[i+1][j] <= height {
				continue
			}
			if j > 0 && heatmap[i][j-1] <= height {
				continue
			}
			if j < len(heatmapRow)-1 && heatmap[i][j+1] <= height {
				continue
			}
			lowPoints = append(lowPoints, []int{i, j})
		}
	}
	return lowPoints
}

func findBasin(heatmap [][]int, i int, j int, basin [][]int) [][]int {
	if contains(basin, i, j) {
		return basin
	}

	basin = append(basin, []int{i, j})

	height := heatmap[i][j]
	if i > 0 && heatmap[i-1][j] >= height && heatmap[i-1][j] != 9 {
		basin = findBasin(heatmap, i-1, j, basin)
	}
	if i < len(heatmap)-1 && heatmap[i+1][j] >= height && heatmap[i+1][j] != 9 {
		basin = findBasin(heatmap, i+1, j, basin)
	}
	if j > 0 && heatmap[i][j-1] >= height && heatmap[i][j-1] != 9 {
		basin = findBasin(heatmap, i, j-1, basin)
	}
	if j < len(heatmap[i])-1 && heatmap[i][j+1] >= height && heatmap[i][j+1] != 9 {
		basin = findBasin(heatmap, i, j+1, basin)
	}
	return basin
}

func contains(arr [][]int, el1, el2 int) bool {
	for _, els := range arr {
		if el1 == els[0] && el2 == els[1] {
			return true
		}
	}
	return false
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
