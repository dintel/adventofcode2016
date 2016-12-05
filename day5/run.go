package day5

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func nextChar(secret string, start int) (string, int) {
	h := md5.New()
	i := start
	for {
		testString := fmt.Sprintf("%s%d", secret, i)
		_, err := io.WriteString(h, testString)
		if err != nil {
			log.Fatalf("Error - %s", err)
		}
		sum := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(sum, "00000") {
			return string(sum[5:6]), i
		}
		h.Reset()
		i++
	}
}

func hasEmptyString(strs []string) bool {
	for _, str := range strs {
		if str == "" {
			return true
		}
	}
	return false
}

func nextValidChar(secret string, start int) (string, int, int) {
	h := md5.New()
	i := start
	for {
		testString := fmt.Sprintf("%s%d", secret, i)
		_, err := io.WriteString(h, testString)
		if err != nil {
			log.Fatalf("Error - %s", err)
		}
		sum := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(sum, "00000") {
			idx, err := strconv.Atoi(sum[5:6])
			if err == nil && idx >= 0 && idx < 8 {
				return string(sum[6:7]), i, idx
			}
		}
		h.Reset()
		i++
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
		i := 0
		result := ""
		var part string
		for j := 0; j < 8; j++ {
			part, i = nextChar(string(data), i)
			result = fmt.Sprintf("%s%s", result, part)
			log.Printf("Character %d is %s at %d iterations, partial code is %s", j, part, i, result)
			i++
		}
		log.Printf("Code to unlock door is %s", result)
	case 2:
		i := 0
		parts := make([]string, 8)
		var part string
		var idx int
		for {
			part, i, idx = nextValidChar(string(data), i)
			if parts[idx] == "" {
				parts[idx] = part
				log.Printf("Character at position %d is %s", idx, part)
			}
			if !hasEmptyString(parts) {
				break
			}
			i++
		}
		log.Printf("Code to unlock second door is %s", strings.Join(parts, ""))
	}
}
