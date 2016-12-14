package day13

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func bitsSet(n int) int {
	var result int
	for result = 0; n != 0; result++ {
		n &= n - 1
	}
	return result
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

type point struct {
	x, y     int
	wall     bool
	distance int
}

type maze struct {
	points [][]*point
	height int
	width  int
}

func newMaze(width, height int) *maze {
	result := &maze{points: make([][]*point, width), height: height, width: width}
	for i := range result.points {
		result.points[i] = make([]*point, height)
		for j := range result.points[i] {
			result.points[i][j] = new(point)
			result.points[i][j].x = i
			result.points[i][j].y = j
			result.points[i][j].distance = -1
			result.points[i][j].wall = false
		}
	}
	return result
}

func (m *maze) adj(x, y int) []*point {
	result := make([]*point, 0, 10)
	startX := x - 1
	if startX < 0 {
		startX = 0
	}
	endX := x + 1
	if endX > m.width-1 {
		endX = m.width - 1
	}

	startY := y - 1
	if startY < 0 {
		startY = 0
	}
	endY := y + 1
	if endY > m.height-1 {
		endY = m.height - 1
	}

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			if i == x && j == y {
				continue
			}
			if i != x && j != y {
				continue
			}
			if !m.points[i][j].wall {
				result = append(result, m.points[i][j])
			}
		}
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

	fav, _ := strconv.Atoi(string(data))
	m := newMaze(50, 50)
	for x := 0; x < len(m.points); x++ {
		for y := 0; y < len(m.points[x]); y++ {
			m.points[x][y].wall = !even(bitsSet(x*x + 3*x + 2*x*y + y + y*y + fav))
		}
	}

	switch part {
	case 1:
		target := m.points[31][39]
		queue := make([]*point, 1)
		queue[0] = m.points[1][1]
		m.points[1][1].distance = 0
		var current *point
		for {
			current, queue = queue[0], queue[1:]
			if current.x == target.x && current.y == target.y {
				log.Printf("Found path with %d steps", current.distance)
				break
			} else {
				adj := m.adj(current.x, current.y)
				for _, p := range adj {
					if p.distance == -1 {
						p.distance = current.distance + 1
						queue = append(queue, p)
					}
				}
			}
		}
	case 2:
		queue := make([]*point, 1)
		queue[0] = m.points[1][1]
		m.points[1][1].distance = 0
		var current *point
		for len(queue) != 0 {
			current, queue = queue[0], queue[1:]
			adj := m.adj(current.x, current.y)
			for _, p := range adj {
				if p.distance == -1 {
					p.distance = current.distance + 1
					queue = append(queue, p)
				}
			}
		}
		result := 0
		for x := 0; x < len(m.points); x++ {
			for y := 0; y < len(m.points[x]); y++ {
				if m.points[x][y].distance <= 50 && m.points[x][y].distance != -1 {
					result++
				}
			}
		}
		log.Printf("There are %d locations that can be reached in 50 steps", result)
	}
}
