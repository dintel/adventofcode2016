package day1

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type location struct {
	x int
	y int
}

type locations []location

func (locs locations) Len() int {
	return len(locs)
}

func (locs locations) Swap(i, j int) {
	locs[i], locs[j] = locs[j], locs[i]
}

func (locs locations) Less(i, j int) bool {
	if locs[i].x == locs[j].x {
		return locs[i].y < locs[j].y
	}
	return locs[i].x < locs[j].x
}

type path struct {
	locations []location
	loc       location
	dir       string
	x         int
	y         int
}

func newPath() *path {
	result := &path{
		locations: make([]location, 0),
		loc:       location{x: 0, y: 0},
		dir:       "N",
		x:         0,
		y:         0,
	}
	result.add()
	return result
}

func move(curDirection, where string, loc location, steps int) (newDirection string, newLoc location) {
	newLoc = loc
	switch curDirection {
	case "N":
		if where == "L" {
			newDirection = "W"
		} else {
			newDirection = "E"
		}
	case "E":
		if where == "L" {
			newDirection = "N"
		} else {
			newDirection = "S"
		}
	case "S":
		if where == "L" {
			newDirection = "E"
		} else {
			newDirection = "W"
		}
	case "W":
		if where == "L" {
			newDirection = "S"
		} else {
			newDirection = "N"
		}
	}

	switch newDirection {
	case "N":
		newLoc.y += steps
	case "E":
		newLoc.x += steps
	case "S":
		newLoc.y -= steps
	case "W":
		newLoc.x -= steps
	}
	return
}

func (p *path) duplicate() bool {
	for _, l := range p.locations {
		if l == p.loc {
			return true
		}
	}
	return false
}

func (p *path) add() {
	p.locations = append(p.locations, p.loc)
	sort.Sort(locations(p.locations))
}

func (p *path) move(where string, steps int) *location {
	switch p.dir {
	case "N":
		if where == "L" {
			p.dir = "W"
		} else {
			p.dir = "E"
		}
	case "E":
		if where == "L" {
			p.dir = "N"
		} else {
			p.dir = "S"
		}
	case "S":
		if where == "L" {
			p.dir = "E"
		} else {
			p.dir = "W"
		}
	case "W":
		if where == "L" {
			p.dir = "S"
		} else {
			p.dir = "N"
		}
	}

	var result *location = nil
	switch p.dir {
	case "N":
		for i := 0; i < steps; i++ {
			p.loc.y++
			if p.duplicate() {
				result = new(location)
				*result = p.loc
			}
			p.add()
		}
	case "E":
		for i := 0; i < steps; i++ {
			p.loc.x++
			if p.duplicate() {
				result = new(location)
				*result = p.loc
			}
			p.add()
		}
	case "S":
		for i := 0; i < steps; i++ {
			p.loc.y--
			if p.duplicate() {
				result = new(location)
				*result = p.loc
			}
			p.add()
		}
	case "W":
		for i := 0; i < steps; i++ {
			p.loc.x--
			if p.duplicate() {
				result = new(location)
				*result = p.loc
			}
			p.add()
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

	switch part {
	case 1:
		arr := strings.Split(string(data), ", ")
		direction := "N"
		loc := location{x: 0, y: 0}
		for i, str := range arr {
			where := str[0:1]
			steps, err := strconv.Atoi(str[1:])
			if err != nil {
				log.Fatalf("Failed parsing step %d (%s) - %s", i, str, err)
			}
			direction, loc = move(direction, where, loc, steps)
		}
		xDir := "North"
		yDir := "East"
		if loc.x < 0 {
			xDir = "South"
			loc.x = -loc.x
		}
		if loc.y < 0 {
			yDir = "West"
			loc.y = -loc.y
		}
		moves := len(arr)
		log.Printf("Easter Bunny HQ position after %d moves - %d to %s, %d to %s (total %d blocks away)", moves, loc.x, xDir, loc.y, yDir, loc.x+loc.y)
	case 2:
		arr := strings.Split(string(data), ", ")
		p := newPath()
		var duplicateLocation *location
		for i, str := range arr {
			where := str[0:1]
			steps, err := strconv.Atoi(str[1:])
			if err != nil {
				log.Fatalf("Failed parsing step %d (%s) - %s", i, str, err)
			}
			duplicateLocation = p.move(where, steps)
			if duplicateLocation != nil {
				break
			}
		}
		xDir := "North"
		yDir := "East"
		if duplicateLocation.x < 0 {
			xDir = "South"
			duplicateLocation.x = -duplicateLocation.x
		}
		if duplicateLocation.y < 0 {
			yDir = "West"
			duplicateLocation.y = -duplicateLocation.y
		}
		log.Printf("Easter Bunny HQ position - %d to %s, %d to %s (total %d blocks away)", duplicateLocation.x, xDir, duplicateLocation.y, yDir, duplicateLocation.x+duplicateLocation.y)
	}
}
