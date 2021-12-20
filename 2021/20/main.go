package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type image struct {
	pixels map[coord]bool
	bounds *bounds
}

type coord struct {
	x int
	y int
}

type bounds struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (b *bounds) outOfBounds(x, y int) bool {
	return x < b.minX || x > b.maxX || y < b.minY || y > b.maxY
}

func (b *bounds) extend() {
	b.minX--
	b.maxX++
	b.minY--
	b.maxY++
}

func (img *image) enhance(filter map[int]bool, enhancements int) {
	for i := 0; i < enhancements; i++ {
		var defPixel string
		if !filter[0] || i%2 == 0 {
			defPixel = "0"
		} else {
			defPixel = "1"
		}
		enhancedImg := img.encanceRound(filter, defPixel)
		img.pixels = enhancedImg.pixels
		img.bounds = enhancedImg.bounds
	}
}

func (img *image) encanceRound(filter map[int]bool, defPixel string) *image {
	enhancedImage := &image{
		pixels: map[coord]bool{},
		bounds: img.bounds,
	}
	for i := img.bounds.minY - 1; i <= img.bounds.maxY+1; i++ {
		for j := img.bounds.minX - 1; j <= img.bounds.maxX+1; j++ {
			enhancedImage.pixels[coord{j, i}] = img.enhancePixel(j, i, filter, defPixel)
		}
	}
	enhancedImage.bounds.extend()
	return enhancedImage
}

func (img *image) enhancePixel(x, y int, filter map[int]bool, defPixel string) bool {
	binRaw := ""
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if img.bounds.outOfBounds(j, i) {
				binRaw += defPixel
			} else if img.pixels[coord{j, i}] {
				binRaw += "1"
			} else {
				binRaw += "0"
			}
		}
	}
	bin, _ := strconv.ParseInt(binRaw, 2, 64)
	return filter[int(bin)]
}

func (img *image) countLitPixels() int {
	litPixels := 0
	for _, p := range img.pixels {
		if p {
			litPixels++
		}
	}
	return litPixels
}

func (img *image) toString() string {
	str := ""
	for i := img.bounds.minY; i <= img.bounds.maxY; i++ {
		for j := img.bounds.minX; j <= img.bounds.maxX; j++ {
			if img.pixels[coord{j, i}] {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	str += "\n"
	return str
}

func main() {
	input := readInput("./input.txt")

	filter := map[int]bool{}
	for i, c := range input[0] {
		filter[i] = c == '#'
	}

	image := &image{
		pixels: map[coord]bool{},
		bounds: &bounds{
			minX: 0,
			maxX: len(input[2]),
			minY: 0,
			maxY: len(input[2:]),
		},
	}
	for i, l := range input[2:] {
		for j, c := range l {
			image.pixels[coord{j, i}] = c == '#'
		}
	}

	solve(image, filter)
}

func solve(image *image, filter map[int]bool) {
	image.enhance(filter, 2)
	litPixels := image.countLitPixels()
	fmt.Printf("Part 1: %d pixels are lit after enhancing twice\n", litPixels)

	image.enhance(filter, 50-2)
	litPixels = image.countLitPixels()
	fmt.Printf("Part 2: %d pixels are lit after enhancing 50 times\n", litPixels)

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
