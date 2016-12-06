package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func mostCommonChar(s string) string {
	m := make(map[rune]int)
	for _, ch := range s {
		if _, exists := m[ch]; !exists {
			m[ch] = 0
		}
		m[ch]++
	}
	max := 0
	var result rune
	for ch, n := range m {
		if max < n {
			max = n
			result = ch
		}
	}
	return string(result)
}

func leastCommonChar(s string) string {
	m := make(map[rune]int)
	for _, ch := range s {
		if _, exists := m[ch]; !exists {
			m[ch] = 0
		}
		m[ch]++
	}
	min := len(s)
	var result rune
	for ch, n := range m {
		if min > n {
			min = n
			result = ch
		}
	}
	return string(result)
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
		columns := make([]string, len(lines[0]))
		for _, line := range lines {
			for i, ch := range line {
				columns[i] = fmt.Sprintf("%s%s", columns[i], string(ch))
			}
		}
		result := ""
		for _, col := range columns {
			result = fmt.Sprintf("%s%s", result, mostCommonChar(col))
		}
		log.Printf("Error-corrected version of the message is '%s'", result)
	case 2:
		lines := strings.Split(string(data), "\n")
		columns := make([]string, len(lines[0]))
		for _, line := range lines {
			for i, ch := range line {
				columns[i] = fmt.Sprintf("%s%s", columns[i], string(ch))
			}
		}
		result := ""
		for _, col := range columns {
			result = fmt.Sprintf("%s%s", result, leastCommonChar(col))
		}
		log.Printf("Error-corrected version of the message is '%s'", result)
	}
}
