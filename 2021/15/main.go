package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	dst  int
	cost int
}

type PriorityQueue struct {
	queue []*Node
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		queue: []*Node{},
	}
}

func (q *PriorityQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *PriorityQueue) add(el *Node) {
	if q.isEmpty() {
		q.queue = append(q.queue, el)
		return
	}

	pos := sort.Search(len(q.queue), func(i int) bool {
		return q.queue[i].cost >= el.cost
	})

	q.queue = append(q.queue[:pos], append([]*Node{el}, q.queue[pos:]...)...)
}

func (q *PriorityQueue) remove() *Node {
	if q.isEmpty() {
		return nil
	}
	el := q.queue[0]
	q.queue = q.queue[1:]
	return el
}

type Graph struct {
	vertices int

	adjMatrix [][]int

	adjList map[int][]*Node
	pq      *PriorityQueue
	settled map[int]bool
	dist    []int
}

func newGraph(vertices int) *Graph {
	g := Graph{
		vertices:  vertices,
		adjMatrix: make([][]int, vertices),
		pq:        NewPriorityQueue(),
		adjList:   make(map[int][]*Node),
		settled:   make(map[int]bool),
		dist:      make([]int, vertices),
	}
	for i := 0; i < vertices; i++ {
		g.adjMatrix[i] = make([]int, vertices)
		g.adjList[i] = make([]*Node, 0)
	}
	return &g
}

func (g *Graph) init(cave [][]int) {
	for i, l := range cave {
		for j := range l {
			x := i*len(l) + j
			// right
			if j < len(l)-1 {
				y := i*len(l) + j + 1
				g.adjMatrix[x][y] = cave[i][j+1]
				g.adjList[x] = append(g.adjList[x], &Node{dst: y, cost: cave[i][j+1]})
			}
			// left
			if j > 0 {
				y := i*len(l) + j - 1
				g.adjMatrix[x][y] = cave[i][j-1]
				g.adjList[x] = append(g.adjList[x], &Node{dst: y, cost: cave[i][j-1]})
			}
			// down
			if i < len(cave)-1 {
				y := (i+1)*len(l) + j
				g.adjMatrix[x][y] = cave[i+1][j]
				g.adjList[x] = append(g.adjList[x], &Node{dst: y, cost: cave[i+1][j]})
			}
			// up
			if i > 0 {
				y := (i-1)*len(l) + j
				g.adjMatrix[x][y] = cave[i-1][j]
				g.adjList[x] = append(g.adjList[x], &Node{dst: y, cost: cave[i-1][j]})
			}
		}
	}
}

func (g *Graph) dijkstra(start int) []int {
	dist := make([]int, g.vertices)
	for i := range dist {
		dist[i] = math.MaxInt
	}

	dist[start] = 0
	visited := map[int]bool{}
	for i := 0; i < g.vertices; i++ {
		x := closestNotVisited(dist, visited)
		visited[x] = true

		for y := 0; y < g.vertices; y++ {
			if g.adjMatrix[x][y] > 0 && !visited[y] && dist[y] > dist[x]+g.adjMatrix[x][y] {
				dist[y] = dist[x] + g.adjMatrix[x][y]
			}
		}
	}

	return dist
}

func closestNotVisited(dist []int, visited map[int]bool) int {
	idx := 0
	min := math.MaxInt
	for i, a := range dist {
		if a < min && !visited[i] {
			min = a
			idx = i
		}
	}
	return idx
}

func (g *Graph) fastDijkstra(start int) []int {
	for i := range g.dist {
		g.dist[i] = math.MaxInt
	}

	g.pq.add(&Node{dst: start, cost: 0})
	g.dist[start] = 0

	for len(g.settled) != g.vertices {
		u := g.pq.remove().dst
		if _, ok := g.settled[u]; ok {
			continue
		}

		g.settled[u] = true
		g.processNeighbours(u)
	}

	return g.dist
}

func (g *Graph) processNeighbours(u int) {
	for _, v := range g.adjList[u] {
		if _, ok := g.settled[v.dst]; ok {
			continue
		}
		distance := g.dist[u] + v.cost
		if distance < g.dist[v.dst] {
			g.dist[v.dst] = distance
		}
		g.pq.add(&Node{dst: v.dst, cost: g.dist[v.dst]})
	}
}

func main() {
	input := readInput("./input.txt")

	cave := [][]int{}
	for _, l := range input {
		c := []int{}
		for _, b := range l {
			d, _ := strconv.Atoi(string(b))
			c = append(c, d)
		}
		cave = append(cave, c)
	}

	solveOne(cave)
	solveTwo(cave)
}

func solveOne(cave [][]int) {
	g := newGraph(len(cave) * len(cave[0]))
	g.init(cave)
	dists := g.dijkstra(0)
	fmt.Printf("Part 1: %v\n", dists[g.vertices-1])
}

func solveTwo(cave [][]int) {
	entireCave := cave
	for i := 0; i < 4; i++ {
		for _, cl := range cave {
			entireCaveLine := []int{}
			for _, c := range cl {
				entireCaveLine = append(entireCaveLine, (c+i)%9+1)
			}
			entireCave = append(entireCave, entireCaveLine)
		}
	}

	for j, cl := range entireCave {
		addCl := []int{}
		for i := 0; i < 4; i++ {
			for k := range cl {
				addCl = append(addCl, (cl[k]+i)%9+1)
			}
		}
		entireCave[j] = append(entireCave[j], addCl...)
	}

	g := newGraph(len(entireCave) * len(entireCave[0]))
	g.init(entireCave)
	dists := g.fastDijkstra(0)
	fmt.Printf("Part 2: %v\n", dists[g.vertices-1])
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
