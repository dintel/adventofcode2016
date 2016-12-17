package day17

import (
	"crypto/md5"
	_ "fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

type state struct {
	secret []byte
	x      int
	y      int
	width  int
	height int
}

func startState(secret []byte, width, height int) *state {
	return &state{
		secret: secret,
		x:      0,
		y:      0,
		width:  width,
		height: height,
	}
}

func (s *state) afterMove(move byte) state {
	result := state{
		secret: make([]byte, len(s.secret)+1),
		x:      s.x,
		y:      s.y,
		width:  s.width,
		height: s.height,
	}
	copy(result.secret, s.secret)
	result.secret[len(s.secret)] = move
	switch move {
	case 'U':
		result.y--
	case 'D':
		result.y++
	case 'L':
		result.x--
	case 'R':
		result.x++
	}
	return result
}

func (s *state) availableMoves() (up, down, left, right bool) {
	sum := md5.Sum(s.secret)
	var hexes [4]byte
	hexes[0] = (sum[0] >> 4) & 0xf
	hexes[1] = sum[0] & 0xf
	hexes[2] = (sum[1] >> 4) & 0xf
	hexes[3] = sum[1] & 0xf
	up = hexes[0] > 10 && s.y > 0
	down = hexes[1] > 10 && s.y < s.height-1
	left = hexes[2] > 10 && s.x > 0
	right = hexes[3] > 10 && s.x < s.width-1
	return
}

func (s *state) nextStates() []state {
	up, down, left, right := s.availableMoves()
	resultLen := bool2int(up) + bool2int(down) + bool2int(left) + bool2int(right)
	result := make([]state, resultLen)
	i := 0
	if up {
		result[i] = s.afterMove('U')
		i++
	}
	if down {
		result[i] = s.afterMove('D')
		i++
	}
	if left {
		result[i] = s.afterMove('L')
		i++
	}
	if right {
		result[i] = s.afterMove('R')
		i++
	}
	return result
}

func (s *state) finished() bool {
	return s.x == s.width-1 && s.y == s.height-1
}

func (s *state) way(secret string) string {
	return strings.TrimPrefix(string(s.secret), secret)
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

	initialState := *startState(data, 4, 4)

	switch part {
	case 1:
		queue := make([]state, 1)
		queue[0] = initialState
		var current state
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]
			if current.finished() {
				log.Printf("Found a way to vault - %s", current.way(string(initialState.secret)))
				break
			}
			queue = append(queue, current.nextStates()...)
		}
	case 2:
		queue := make([]state, 1)
		queue[0] = initialState
		var current state
		maxWay := ""
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]
			if current.finished() {
				newWay := current.way(string(initialState.secret))
				if len(newWay) > len(maxWay) {
					maxWay = newWay
				}
			} else {
				queue = append(queue, current.nextStates()...)
			}
		}
		log.Printf("Longest found way length is %d", len(maxWay))
	}
}
