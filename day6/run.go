package day6

import (
	"io/ioutil"
	"log"
	"os"
)

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
	_ = data

	switch part {
	case 1:
	case 2:
		log.Print("Part 2 is not implemented yet")
	}
}
