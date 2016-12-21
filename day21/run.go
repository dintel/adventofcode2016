package day21

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func swapPosition(s string, a, b int) string {
	result := []byte(s)
	result[a], result[b] = result[b], result[a]
	return string(result)
}

func swapLetter(s string, a, b byte) string {
	result := []byte(s)
	for i := range result {
		if result[i] == a {
			result[i] = b
		} else if result[i] == b {
			result[i] = a
		}
	}
	return string(result)
}

func rotateRight(s string, n int) string {
	result := []byte(s)
	l := len(s)
	for i := 0; i < n; i++ {
		result = append(result[l-1:], result[:l-1]...)
	}
	return string(result)
}

func rotateLeft(s string, n int) string {
	result := []byte(s)
	for i := 0; i < n; i++ {
		result = append(result[1:], result[0])
	}
	return string(result)
}

func rotateLetter(s string, ch byte) string {
	n := strings.IndexByte(s, ch)
	if n >= 4 {
		n += 2
	} else {
		n += 1
	}
	return rotateRight(s, n)
}

func unrotateLetter(s string, ch byte) string {
	result := s
	for i := 0; i < len(s); i++ {
		result = rotateLeft(s, i)
		if rotateLetter(result, ch) == s {
			return result
		}
	}
	log.Printf("Failed unrotating %s by %c", s, ch)
	return s
}

func reverse(s string, a, b int) string {
	result := []byte(s)
	max := (b - a) / 2
	for i := 0; i <= max; i++ {
		result[a+i], result[b-i] = result[b-i], result[a+i]
	}
	return string(result)
}

func move(s string, a, b int) string {
	result := []byte(s)
	ch := result[a]
	result = append(result[:a], result[a+1:]...)
	if b == len(s)-1 {
		result = append(result, ch)
	} else {
		part := make([]byte, len(result[b:])+1)
		part[0] = ch
		for i := 1; i < len(part); i++ {
			part[i] = result[b+i-1]
		}
		result = append(result[:b], part...)
	}
	return string(result)
}

func scramble(password string, lines []string) string {
	result := password
	for _, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "swap":
			switch parts[1] {
			case "position":
				a, _ := strconv.Atoi(parts[2])
				b, _ := strconv.Atoi(parts[5])
				result = swapPosition(result, a, b)
			case "letter":
				result = swapLetter(result, parts[2][0], parts[5][0])
			}
		case "rotate":
			switch parts[1] {
			case "left":
				a, _ := strconv.Atoi(parts[2])
				result = rotateLeft(result, a)
			case "right":
				a, _ := strconv.Atoi(parts[2])
				result = rotateRight(result, a)
			case "based":
				result = rotateLetter(result, parts[6][0])
			}
		case "reverse":
			a, _ := strconv.Atoi(parts[2])
			b, _ := strconv.Atoi(parts[4])
			result = reverse(result, a, b)
		case "move":
			a, _ := strconv.Atoi(parts[2])
			b, _ := strconv.Atoi(parts[5])
			result = move(result, a, b)
		}
	}
	return result
}

func unscramble(password string, lines []string) string {
	result := password
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "swap":
			switch parts[1] {
			case "position":
				a, _ := strconv.Atoi(parts[2])
				b, _ := strconv.Atoi(parts[5])
				result = swapPosition(result, a, b)
			case "letter":
				result = swapLetter(result, parts[2][0], parts[5][0])
			}
		case "rotate":
			switch parts[1] {
			case "left":
				a, _ := strconv.Atoi(parts[2])
				result = rotateRight(result, a)
			case "right":
				a, _ := strconv.Atoi(parts[2])
				result = rotateLeft(result, a)
			case "based":
				result = unrotateLetter(result, parts[6][0])
			}
		case "reverse":
			a, _ := strconv.Atoi(parts[2])
			b, _ := strconv.Atoi(parts[4])
			result = reverse(result, a, b)
		case "move":
			a, _ := strconv.Atoi(parts[2])
			b, _ := strconv.Atoi(parts[5])
			result = move(result, b, a)
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
	lines := strings.Split(string(data), "\n")

	switch part {
	case 1:
		password := "abcdefgh"
		result := scramble(password, lines)
		log.Printf("Password is %s", password)
		log.Printf("Scrambled password is %s", result)
	case 2:
		password := "fbgdceah"
		result := unscramble(password, lines)
		log.Printf("Scrambled password is %s", password)
		log.Printf("Unscrambled password is %s", result)
		log.Printf("Rescrambled password is %s", scramble(result, lines))
	}
}
