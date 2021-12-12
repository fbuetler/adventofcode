package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const unlimitedVisites = -1

type graph struct {
	vertices       []string
	edges          map[string][]string
	allowedVisites map[string]int
}

func NewGraph() *graph {
	var g graph
	g.vertices = []string{}
	g.edges = map[string][]string{}
	g.allowedVisites = map[string]int{}
	return &g
}

func (g *graph) AddEdge(a, b string) {
	g.vertices = append(g.vertices, a)
	g.vertices = append(g.vertices, b)

	g.edges[a] = append(g.edges[a], b)
	g.edges[b] = append(g.edges[b], a)

	g.allowedVisites[a] = 1
	g.allowedVisites[b] = 1

	if strings.ToUpper(a) == a {
		g.allowedVisites[a] = unlimitedVisites
	}
	if strings.ToUpper(b) == b {
		g.allowedVisites[b] = unlimitedVisites
	}
}

func (g *graph) FindAllPaths(src, dst string, allPaths *[]string) {
	visited := map[string]int{}
	path := []string{}
	g.findAllPathsHelper(src, dst, visited, path, allPaths)
}

func (g *graph) findAllPathsHelper(curr, dst string, visited map[string]int, path []string, allPaths *[]string) {
	visited[curr]++
	path = append(path, curr)

	if curr == dst {
		p := strings.Join(path, ",")
		if !contains(*allPaths, p) {
			*allPaths = append(*allPaths, p)
		}
	} else {
		for _, next := range g.edges[curr] {
			if visited[next] < g.allowedVisites[next] || g.allowedVisites[next] == unlimitedVisites {
				g.findAllPathsHelper(next, dst, visited, path, allPaths)
			}
		}
	}

	n := len(path) - 1
	path = path[:n]
	visited[curr]--
}

func main() {
	input := readInput("./input.txt")

	g := NewGraph()
	for _, line := range input {
		vertices := strings.Split(line, "-")
		g.AddEdge(vertices[0], vertices[1])
	}

	solveOne(g)
	solveTwo(g)
}

func solveOne(g *graph) {
	allPaths := []string{}
	g.FindAllPaths("start", "end", &allPaths)
	fmt.Printf("Part 1: There are %d paths through this cave system\n", len(allPaths))
}

func solveTwo(g *graph) {
	allPaths := []string{}
	for isAllowedTwice, val := range g.allowedVisites {
		if val == unlimitedVisites || isAllowedTwice == "start" || isAllowedTwice == "end" {
			continue
		}
		g.allowedVisites[isAllowedTwice] = 2
		g.FindAllPaths("start", "end", &allPaths)
		g.allowedVisites[isAllowedTwice] = 1
	}
	fmt.Printf("Part 2: There are %d paths through this cave system\n", len(allPaths))
}

func contains(aa []string, b string) bool {
	for _, a := range aa {
		if a == b {
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
