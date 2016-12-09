package day8

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type screen struct {
	width  int
	height int
	values [][]bool
}

func newScreen(width, height int) *screen {
	result := &screen{
		width:  width,
		height: height,
		values: make([][]bool, width),
	}
	for i := 0; i < width; i++ {
		result.values[i] = make([]bool, height)
	}
	return result
}

func (s *screen) rect(width, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s.values[x][y] = true
		}
	}
}

func (s *screen) rotateRow(row, n int) {
	for i := 0; i < n; i++ {
		carry := s.values[s.width-1][row]
		for j := 0; j < s.width; j++ {
			s.values[j][row], carry = carry, s.values[j][row]
		}
	}
}

func (s *screen) rotateColumn(column, n int) {
	for i := 0; i < n; i++ {
		var last bool
		last, s.values[column] = s.values[column][len(s.values[column])-1], s.values[column][:len(s.values[column])-1]
		s.values[column] = append([]bool{last}, s.values[column]...)
	}
}

func (s *screen) countOn() int {
	result := 0
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			if s.values[x][y] {
				result++
			}
		}
	}
	return result
}

func (s *screen) String() string {
	result := make([]string, s.height)
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			if s.values[x][y] {
				result[y] = fmt.Sprintf("%s%s", result[y], "*")
			} else {
				result[y] = fmt.Sprintf("%s%s", result[y], " ")
			}
		}
	}
	return strings.Join(result, "\n")
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
		lines := strings.Split(string(data), "\n")
		s := newScreen(50, 6)
		for _, line := range lines {
			parts := strings.Split(line, " ")
			if parts[0] == "rect" {
				size := strings.Split(parts[1], "x")
				width, _ := strconv.Atoi(size[0])
				height, _ := strconv.Atoi(size[1])
				s.rect(width, height)
			} else if parts[0] == "rotate" {
				if parts[1] == "row" {
					coord := strings.Split(parts[2], "=")
					row, _ := strconv.Atoi(coord[1])
					n, _ := strconv.Atoi(parts[4])
					s.rotateRow(row, n)
				} else if parts[1] == "column" {
					coord := strings.Split(parts[2], "=")
					column, _ := strconv.Atoi(coord[1])
					n, _ := strconv.Atoi(parts[4])
					s.rotateColumn(column, n)
				} else {
					log.Fatalf("Failed parsing line '%s'", line)
				}
			} else {
				log.Fatalf("Failed parsing line '%s'", line)
			}
		}
		log.Printf("After executing %d instructions, there are %d lights on", len(lines), s.countOn())
	case 2:
		lines := strings.Split(string(data), "\n")
		s := newScreen(50, 6)
		for _, line := range lines {
			parts := strings.Split(line, " ")
			if parts[0] == "rect" {
				size := strings.Split(parts[1], "x")
				width, _ := strconv.Atoi(size[0])
				height, _ := strconv.Atoi(size[1])
				s.rect(width, height)
			} else if parts[0] == "rotate" {
				if parts[1] == "row" {
					coord := strings.Split(parts[2], "=")
					row, _ := strconv.Atoi(coord[1])
					n, _ := strconv.Atoi(parts[4])
					s.rotateRow(row, n)
				} else if parts[1] == "column" {
					coord := strings.Split(parts[2], "=")
					column, _ := strconv.Atoi(coord[1])
					n, _ := strconv.Atoi(parts[4])
					s.rotateColumn(column, n)
				} else {
					log.Fatalf("Failed parsing line '%s'", line)
				}
			} else {
				log.Fatalf("Failed parsing line '%s'", line)
			}
		}
		log.Print("Following is what display shows")
		log.Printf("\n%s", s)
	}
}
