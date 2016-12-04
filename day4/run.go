package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseLine(line string) (name string, sector int, checksum string) {
	first := strings.Split(line, "[")
	checksum = first[1][:len(first[1])-1]
	i := len(first[0]) - 1
	for string(first[0][i]) != "-" {
		i--
	}
	name = first[0][:i]
	sector, _ = strconv.Atoi(first[0][i+1:])
	return
}

func findBiggest(scores map[string]int) string {
	var keys []string
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	result := keys[0]
	for i := 1; i < len(keys); i++ {
		k := keys[i]
		if scores[k] > scores[result] {
			result = k
		}
	}
	return result
}

func checksum(name string) string {
	scores := make(map[string]int)
	for _, ch := range name {
		if ch == '-' {
			continue
		}
		s := string(ch)
		if _, exists := scores[s]; !exists {
			scores[s] = 0
		} else {
			scores[s]++
		}
	}
	result := ""
	for i := 0; i < 5; i++ {
		biggest := findBiggest(scores)
		result = fmt.Sprintf("%s%s", result, biggest)
		delete(scores, biggest)
	}
	return result
}

func decrypt(name string, sector int) string {
	return strings.Map(func(r rune) rune {
		if r == '-' {
			return ' '
		}
		s := int(r) + sector
		if s > 'z' {
			for s > 'z' {
				s -= 26
			}
		} else if s < 'a' {
			for s < 'a' {
				s += 26
			}
		}
		return rune(s)
	}, name)
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
		sum := 0
		for _, line := range lines {
			name, sector, chk := parseLine(line)
			if checksum(name) == chk {
				sum += sector
			}
		}
		log.Printf("Sum of sector IDs of real rooms is %d", sum)
	case 2:
		lines := strings.Split(string(data), "\n")
		neededSector := 0
		neededName := ""
		for _, line := range lines {
			name, sector, chk := parseLine(line)
			if checksum(name) == chk {
				realName := decrypt(name, sector)
				if strings.HasPrefix(realName, "north") {
					neededSector = sector
					neededName = realName
					break
				}
			}
		}
		log.Printf("Found sector ID %d - %s", neededSector, neededName)
	}
}
