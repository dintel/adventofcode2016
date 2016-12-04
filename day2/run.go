package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type keypad struct {
	current int
}

func newKeypad() *keypad {
	return &keypad{current: 5}
}

func (k *keypad) move(direction string) {
	switch direction {
	case "U":
		if k.current-3 > 0 {
			k.current -= 3
		}
	case "D":
		if k.current+3 < 10 {
			k.current += 3
		}
	case "L":
		if k.current != 1 && k.current != 4 && k.current != 7 {
			k.current--
		}
	case "R":
		if k.current != 3 && k.current != 6 && k.current != 9 {
			k.current++
		}
	}
}

func (k *keypad) move2(direction string) {
	switch direction {
	case "U":
		if k.current == 13 {
			k.current = 11
		} else if k.current == 10 || k.current == 11 || k.current == 12 || k.current == 8 || k.current == 7 || k.current == 6 {
			k.current -= 4
		} else if k.current == 3 {
			k.current = 1
		}
	case "D":
		if k.current == 1 {
			k.current = 3
		} else if k.current == 2 || k.current == 3 || k.current == 4 || k.current == 6 || k.current == 7 || k.current == 8 {
			k.current += 4
		} else if k.current == 11 {
			k.current = 13
		}
	case "L":
		if k.current == 3 || k.current == 4 || k.current == 6 || k.current == 7 || k.current == 8 || k.current == 9 || k.current == 11 || k.current == 12 {
			k.current--
		}
	case "R":
		if k.current == 2 || k.current == 3 || k.current == 5 || k.current == 6 || k.current == 7 || k.current == 8 || k.current == 10 || k.current == 11 {
			k.current++
		}
	}
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
		keypad := newKeypad()
		lines := strings.Split(string(data), "\n")
		code := make([]int, len(lines))
		for i, line := range lines {
			for _, ch := range line {
				keypad.move(string(ch))
			}
			code[i] = keypad.current
		}
		codeStr := ""
		for _, n := range code {
			codeStr = fmt.Sprintf("%s%d", codeStr, n)
		}
		log.Printf("Code to bathroom is %s", codeStr)
	case 2:
		keypad := newKeypad()
		lines := strings.Split(string(data), "\n")
		code := make([]int, len(lines))
		for i, line := range lines {
			for _, ch := range line {
				keypad.move2(string(ch))
			}
			code[i] = keypad.current
		}
		codeStr := ""
		for _, n := range code {
			if n < 10 {
				codeStr = fmt.Sprintf("%s%d", codeStr, n)
			} else if n == 10 {
				codeStr = fmt.Sprintf("%sA", codeStr)
			} else if n == 11 {
				codeStr = fmt.Sprintf("%sB", codeStr)
			} else if n == 12 {
				codeStr = fmt.Sprintf("%sC", codeStr)
			} else if n == 13 {
				codeStr = fmt.Sprintf("%sD", codeStr)
			}
		}
		log.Printf("Code to bathroom is %s", codeStr)
	}
}
