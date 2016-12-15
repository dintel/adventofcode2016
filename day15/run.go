package day15

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type disc struct {
	n             int
	positions     int
	startPosition int
}

func newDisc(n, positions, startPosition int) *disc {
	return &disc{
		n:             n,
		positions:     positions,
		startPosition: startPosition,
	}
}

func (d *disc) position(t int) int {
	return (d.startPosition + t) % d.positions
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
	discs := make([]*disc, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		nStr := strings.TrimPrefix(parts[1], "#")
		n, _ := strconv.Atoi(nStr)
		positions, _ := strconv.Atoi(parts[3])
		startPosition, _ := strconv.Atoi(strings.Trim(parts[len(parts)-1], "."))
		discs = append(discs, newDisc(n, positions, startPosition))
	}

	switch part {
	case 1:
		t := 0
		for {
			wouldFall := true
			for i := 0; i < len(discs); i++ {
				if discs[i].position(t+i+1) != 0 {
					wouldFall = false
					break
				}
			}
			if wouldFall {
				break
			}
			t++
		}
		log.Printf("Push the button at time=%d to get capsule", t)
	case 2:
		discs = append(discs, newDisc(len(discs)+1, 11, 0))
		t := 0
		for {
			wouldFall := true
			for i := 0; i < len(discs); i++ {
				if discs[i].position(t+i+1) != 0 {
					wouldFall = false
					break
				}
			}
			if wouldFall {
				break
			}
			t++
		}
		log.Printf("Push the button at time=%d to get capsule", t)
	}
}
