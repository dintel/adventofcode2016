package day18

import (
	"io/ioutil"
	"log"
	"os"
)

type row []bool

func (r row) safeTiles() int {
	result := 0
	for _, tile := range r {
		if !tile {
			result++
		}
	}
	return result
}

func generateRow(prev row) row {
	result := make([]bool, len(prev))
	for i := range prev {
		l := false
		c := prev[i]
		r := false
		if i != 0 {
			l = prev[i-1]
		}
		if i != len(prev)-1 {
			r = prev[i+1]
		}
		result[i] = (l && c && !r) || (!l && c && r) || (l && !c && !r) || (!l && !c && r)
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
	firstRow := make([]bool, len(data))
	for i, b := range data {
		firstRow[i] = b == '^'
	}

	switch part {
	case 1:
		rows := make([]row, 40)
		rows[0] = row(firstRow)
		result := rows[0].safeTiles()
		for i := 1; i < len(rows); i++ {
			rows[i] = generateRow(rows[i-1])
			result += rows[i].safeTiles()
		}
		log.Printf("There are %d safe tiles", result)
	case 2:
		r := row(firstRow)
		result := r.safeTiles()
		for i := 1; i < 400000; i++ {
			r = generateRow(r)
			result += r.safeTiles()
		}
		log.Printf("There are %d safe tiles", result)
	}
}
