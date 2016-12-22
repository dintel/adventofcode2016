package day22

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type gridNode struct {
	x         int
	y         int
	size      int
	used      int
	available int
}

func newGridNode(x, y, size, used, available int) *gridNode {
	return &gridNode{
		x:         x,
		y:         y,
		size:      size,
		used:      used,
		available: available,
	}
}

func viablePairs(nodes []*gridNode) int {
	result := 0
	for i := range nodes {
		node := nodes[i]
		if node.used != 0 {
			for j := range nodes {
				if i != j && nodes[j].available >= node.used {
					result++
				}
			}
		}
	}
	return result
}

type coord struct {
	x, y int
}

func distance(from, to coord) float64 {
	deltaX := float64(to.x - from.x)
	deltaY := float64(to.y - from.y)
	return math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
}

type coordQueue struct {
	q      []coord
	target coord
}

func newCoordQueue(start, target coord) *coordQueue {
	result := &coordQueue{
		q:      make([]coord, 1),
		target: target,
	}
	result.q[0] = start
	return result
}

func (q *coordQueue) Pop() coord {
	var result coord
	result, q.q = q.q[0], q.q[1:]
	return result
}

func (q *coordQueue) Append(c coord) {
	q.q = append(q.q, c)
}

func (q *coordQueue) Len() int {
	return len(q.q)
}

func (q *coordQueue) Swap(i, j int) {
	q.q[i], q.q[j] = q.q[j], q.q[i]
}

func (q *coordQueue) Less(i, j int) bool {
	return distance(q.q[i], q.target) < distance(q.q[j], q.target)
}

type grid struct {
	nodes  map[coord]*gridNode
	width  int
	height int
}

func newGrid() *grid {
	return &grid{
		nodes:  make(map[coord]*gridNode),
		width:  0,
		height: 0,
	}
}

func (g *grid) String() string {
	result := "\n "
	for x := 0; x < g.width; x++ {
		result = fmt.Sprintf("%s%d", result, x%10)
	}
	result = fmt.Sprintf("%s\n", result)
	for y := 0; y < g.height; y++ {
		result = fmt.Sprintf("%s%d", result, y%10)
		for x := 0; x < g.width; x++ {
			c := coord{x: x, y: y}
			code := "."
			if g.nodes[c].size > 100 {
				code = "#"
			} else if g.nodes[c].used == 0 {
				code = "_"
			}
			result = fmt.Sprintf("%s%s", result, code)
		}
		result = fmt.Sprintf("%s\n", result)
	}
	return result
}

func (g *grid) addNode(node *gridNode) {
	c := coord{x: node.x, y: node.y}
	g.nodes[c] = node
	if node.x+1 > g.width {
		g.width = node.x + 1
	}
	if node.y+1 > g.height {
		g.height = node.y + 1
	}
}

func (g *grid) findEmpty() coord {
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			c := coord{x: x, y: y}
			if g.nodes[c].used == 0 {
				return c
			}
		}
	}
	log.Printf("Could not find coordinates with empty node")
	return coord{x: -1, y: -1}
}

func (g *grid) left(c coord) (coord, bool) {
	if c.x == 0 {
		return c, false
	}
	return coord{x: c.x - 1, y: c.y}, true
}

func (g *grid) right(c coord) (coord, bool) {
	if c.x == g.width-1 {
		return c, false
	}
	return coord{x: c.x + 1, y: c.y}, true
}

func (g *grid) up(c coord) (coord, bool) {
	if c.y == g.height-1 {
		return c, false
	}
	return coord{x: c.x, y: c.y + 1}, true
}

func (g *grid) down(c coord) (coord, bool) {
	if c.y == 0 {
		return c, false
	}
	return coord{x: c.x, y: c.y - 1}, true
}

func (g *grid) shortestPath(from, to coord) []coord {
	prev := make(map[coord]coord)
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			c := coord{x: x, y: y}
			prev[c] = coord{x: -1, y: -1}
		}
	}

	queue := newCoordQueue(from, to)
	var current coord
	for queue.Len() != 0 {
		current = queue.Pop()
		if current == to {
			break
		}

		up, exists := g.up(current)
		if exists && g.nodes[up].size < 100 && g.nodes[up].used <= g.nodes[current].size {
			if prev[up].x == -1 && prev[up].y == -1 {
				prev[up] = current
				queue.Append(up)
			}
		}

		right, exists := g.right(current)
		if exists && g.nodes[right].size < 100 && g.nodes[right].used <= g.nodes[current].size {
			if prev[right].x == -1 && prev[right].y == -1 {
				prev[right] = current
				queue.Append(right)
			}
		}

		down, exists := g.down(current)
		if exists && g.nodes[down].size < 100 && g.nodes[down].used <= g.nodes[current].size {
			if prev[down].x == -1 && prev[down].y == -1 {
				prev[down] = current
				queue.Append(down)
			}
		}

		left, exists := g.left(current)
		if exists && g.nodes[left].size < 100 && g.nodes[left].used <= g.nodes[current].size {
			if prev[left].x == -1 && prev[left].y == -1 {
				prev[left] = current
				queue.Append(left)
			}
		}
		sort.Sort(queue)
	}

	result := make([]coord, 0)
	current = to
	for current != from {
		result = append(result, current)
		current = prev[current]
	}
	result = append(result, current)
	return result
}

func (g *grid) move(from, to coord) {
	g.nodes[to].used = g.nodes[from].used
	g.nodes[to].available = g.nodes[to].size - g.nodes[to].used
	g.nodes[from].used = 0
	g.nodes[from].available = g.nodes[from].size - g.nodes[from].used
}

func (g *grid) execPath(path []coord) coord {
	for i := len(path) - 1; i > 0; i-- {
		g.move(path[i-1], path[i])
	}
	return path[0]
}

func (g *grid) block(c coord) {
	g.nodes[c].size += 100
}

func (g *grid) unblock(c coord) {
	g.nodes[c].size -= 100
}

func Run(part int) {
	if len(os.Args) != 4 {
		log.Fatalf("Expected input file parameter")
	}
	filename := os.Args[3]
	log.Printf("Loading file %s", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read input file - %s", err)
	}

	lines := strings.Split(string(data), "\n")
	lines = lines[2:]
	nodes := make([]*gridNode, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		size, _ := strconv.Atoi(strings.Trim(parts[1], "T"))
		used, _ := strconv.Atoi(strings.Trim(parts[2], "T"))
		avail, _ := strconv.Atoi(strings.Trim(parts[3], "T"))
		parts = strings.Split(parts[0], "-")
		x, _ := strconv.Atoi(strings.Trim(parts[1], "x"))
		y, _ := strconv.Atoi(strings.Trim(parts[2], "y"))
		nodes[i] = newGridNode(x, y, size, used, avail)
	}

	switch part {
	case 1:
		log.Printf("There are %d viable pair of nodes", viablePairs(nodes))
	case 2:
		g := newGrid()
		for i := range nodes {
			g.addNode(nodes[i])
		}
		empty := g.findEmpty()
		target := coord{x: g.width - 2, y: 0}
		dataAt := coord{x: g.width - 1, y: 0}
		final := coord{x: 0, y: 0}
		shortestPath := g.shortestPath(empty, target)
		empty = g.execPath(shortestPath)
		minMoves := len(shortestPath) - 1
		for dataAt != final {
			g.move(dataAt, empty)
			dataAt, empty = empty, dataAt
			if dataAt == final {
				break
			}
			minMoves++
			g.block(dataAt)
			target.x = dataAt.x - 1
			target.y = dataAt.y
			shortestPath = g.shortestPath(empty, target)
			g.unblock(dataAt)
			empty = g.execPath(shortestPath)
			minMoves += len(shortestPath) - 1
		}
		g.move(coord{x: target.x, y: target.y + 1}, target)
		minMoves++
		log.Printf("Minimal number of moves to get data to (0,0) is %d", minMoves)
	}
}
