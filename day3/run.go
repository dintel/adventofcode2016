package day3

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func validTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
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
		valid := 0
		for _, line := range lines {
			sides := strings.Split(line, " ")
			if len(sides) != 3 {
				log.Printf("Invalid line %s", line)
				continue
			}
			a, err := strconv.Atoi(sides[0])
			if err != nil {
				log.Printf("Failed converting %s to number, skipping line '%s'", sides[0], line)
			}
			b, err := strconv.Atoi(sides[1])
			if err != nil {
				log.Printf("Failed converting %s to number, skipping line '%s'", sides[1], line)
			}
			c, err := strconv.Atoi(sides[2])
			if err != nil {
				log.Printf("Failed converting %s to number, skipping line '%s'", sides[3], line)
			}
			if validTriangle(a, b, c) {
				valid++
			}
		}
		log.Printf("Found %d valid triangles (out of %d analyzed)", valid, len(lines))
	case 2:
		lines := strings.Split(string(data), "\n")
		numbers := make([][3]int, len(lines))
		for i, line := range lines {
			sides := strings.Split(line, " ")
			if len(sides) != 3 {
				log.Printf("Invalid line %s", line)
				continue
			}
			a, err := strconv.Atoi(sides[0])
			if err != nil {
				log.Printf("Failed converting %s to number, skipping line '%s'", sides[0], line)
			}
			b, err := strconv.Atoi(sides[1])
			if err != nil {
				log.Printf("Failed converting %s to number, skipping line '%s'", sides[1], line)
			}
			c, err := strconv.Atoi(sides[2])
			if err != nil {
				log.Printf("Failed converting %s to number, skipping line '%s'", sides[3], line)
			}
			numbers[i][0] = a
			numbers[i][1] = b
			numbers[i][2] = c
		}

		for i := 0; i < len(numbers); i += 3 {
			numbers[i][1], numbers[i+1][0] = numbers[i+1][0], numbers[i][1]
			numbers[i][2], numbers[i+2][0] = numbers[i+2][0], numbers[i][2]
			numbers[i+1][2], numbers[i+2][1] = numbers[i+2][1], numbers[i+1][2]
		}

		valid := 0
		for _, line := range numbers {
			if validTriangle(line[0], line[1], line[2]) {
				valid++
			}
		}

		log.Printf("Found %d valid triangles (out of %d analyzed)", valid, len(lines))
	}
}
