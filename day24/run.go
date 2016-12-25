package day24

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type path []point

func concatPaths(a, b path) path {
	result := path(make([]point, len(a)+len(b)-1))
	for i := range a {
		result[i] = a[i]
	}
	l := len(a) - 1
	for j := 1; j < len(b); j++ {
		result[l+j] = b[j]
	}
	return result
}

type hvacMap struct {
	height int
	width  int
	walls  map[point]bool
	paths  map[[2]point][]point
}

func newHvacMap(width, height int) *hvacMap {
	result := &hvacMap{
		height: height,
		width:  width,
		walls:  make(map[point]bool),
		paths:  make(map[[2]point][]point),
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := point{x: x, y: y}
			result.walls[p] = false
		}
	}
	return result
}

func (m *hvacMap) setPoint(x, y int, wall bool) {
	p := point{x: x, y: y}
	m.walls[p] = wall
}

func (m *hvacMap) up(p point) (point, bool) {
	result := point{x: p.x, y: p.y + 1}
	_, exists := m.walls[result]
	return result, exists
}

func (m *hvacMap) right(p point) (point, bool) {
	result := point{x: p.x + 1, y: p.y}
	_, exists := m.walls[result]
	return result, exists
}

func (m *hvacMap) down(p point) (point, bool) {
	result := point{x: p.x, y: p.y - 1}
	_, exists := m.walls[result]
	return result, exists
}

func (m *hvacMap) left(p point) (point, bool) {
	result := point{x: p.x - 1, y: p.y}
	_, exists := m.walls[result]
	return result, exists
}

func (m *hvacMap) shortestPath(from, to point) []point {
	if computedPath, exists := m.paths[[2]point{from, to}]; exists {
		return computedPath
	}
	prev := make(map[point]point)
	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			p := point{x: x, y: y}
			prev[p] = point{x: -1, y: -1}
		}
	}

	queue := newPointsQueue(from, to)
	var current point
	for queue.Len() != 0 {
		current = queue.Pop()
		if current == to {
			break
		}

		up, exists := m.up(current)
		if exists && !m.walls[up] && prev[up].x == -1 && prev[up].y == -1 {
			prev[up] = current
			queue.Append(up)
		}

		right, exists := m.right(current)
		if exists && !m.walls[right] && prev[right].x == -1 && prev[right].y == -1 {
			prev[right] = current
			queue.Append(right)
		}

		down, exists := m.down(current)
		if exists && !m.walls[down] && prev[down].x == -1 && prev[down].y == -1 {
			prev[down] = current
			queue.Append(down)
		}

		left, exists := m.left(current)
		if exists && !m.walls[left] && prev[left].x == -1 && prev[left].y == -1 {
			prev[left] = current
			queue.Append(left)
		}
	}

	result := make([]point, 0)
	current = to
	for current != from {
		result = append(result, current)
		current = prev[current]
	}
	result = append(result, current)
	result = reverse(result)
	m.paths[[2]point{from, to}] = result
	return result
}

func (m *hvacMap) bestPath(from point, targets []point) []point {
	if len(targets) == 1 {
		return m.shortestPath(from, targets[0])
	}
	bestPath := path(nil)
	for i := range targets {
		newTargets := filter(targets, i)
		p := concatPaths(m.shortestPath(from, targets[i]), m.bestPath(targets[i], newTargets))
		if bestPath == nil || len(bestPath) > len(p) {
			bestPath = p
		}
	}
	return bestPath
}

func (m *hvacMap) bestCycle(from point, targets []point, first point) []point {
	if len(targets) == 1 {
		return concatPaths(m.shortestPath(from, targets[0]), m.shortestPath(targets[0], first))
	}
	bestPath := path(nil)
	for i := range targets {
		newTargets := filter(targets, i)
		p := concatPaths(m.shortestPath(from, targets[i]), m.bestCycle(targets[i], newTargets, first))
		if bestPath == nil || len(bestPath) > len(p) {
			bestPath = p
		}
	}
	return bestPath
}

type pointsQueue struct {
	q      []point
	target point
}

func newPointsQueue(start, target point) *pointsQueue {
	result := &pointsQueue{
		q:      make([]point, 1),
		target: target,
	}
	result.q[0] = start
	return result
}

func (q *pointsQueue) Pop() point {
	var result point
	result, q.q = q.q[0], q.q[1:]
	return result
}

func (q *pointsQueue) Append(c point) {
	q.q = append(q.q, c)
}

func (q *pointsQueue) Len() int {
	return len(q.q)
}

func filter(points []point, i int) []point {
	result := make([]point, 0)
	for j := range points {
		if j != i {
			result = append(result, points[j])
		}
	}
	return result
}

func reverse(points []point) []point {
	result := make([]point, len(points))
	for i := range points {
		result[len(result)-1-i] = points[i]
	}
	return result
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
	hvac := newHvacMap(len(lines[0]), len(lines))
	targets := make([]point, 0)
	var start point
	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case '#':
				hvac.setPoint(x, y, true)
			case '.':
				hvac.setPoint(x, y, false)
			default:
				hvac.setPoint(x, y, false)
				p := point{x: x, y: y}
				if ch == '0' {
					start = p
				} else {
					targets = append(targets, p)
				}
			}
		}
	}

	switch part {
	case 1:
		log.Printf("Found best path with %d moves", len(hvac.bestPath(start, targets))-1)
	case 2:
		bestCycle := hvac.bestCycle(start, targets, start)
		log.Printf("Found best cycle (path including returning robot) with %d moves", len(bestCycle)-1)
	}
}
