package day19

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

type elf struct {
	id         int
	hasPresent bool
}

type round []elf

func (r round) nextElfWithPresent(i int) int {
	i++
	if i >= len(r) {
		i = 0
	}
	for !r[i].hasPresent {
		i++
		if i >= len(r) {
			i = 0
		}
	}
	return i
}

func (r round) withPresent() int {
	for _, e := range r {
		if e.hasPresent {
			return e.id
		}
	}
	return -1
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

	elfCount, _ := strconv.Atoi(string(data))

	switch part {
	case 1:
		elves := round(make([]elf, elfCount))
		for i := 0; i < elfCount; i++ {
			elves[i].id = i + 1
			elves[i].hasPresent = true
		}
		i := 0
		hasPresent := elfCount
		for hasPresent != 1 {
			i = elves.nextElfWithPresent(i)
			elves[i].hasPresent = false
			hasPresent--
			i = elves.nextElfWithPresent(i)
		}
		log.Printf("Last remaining elf index is %d", elves.withPresent())
	case 2:
		n := float64(elfCount)
		x := math.Pow(3, math.Floor(math.Log2(n)/math.Log2(3)))
		result := n
		if x != n {
			result = math.Max(n-x, 2*n-3*x)
		}
		log.Printf("Last remaining elf index is %d", int(result))
	}
}
