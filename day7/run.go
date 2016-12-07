package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func isAbba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func findAbas(s string) []string {
	result := make([]string, 0)
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			result = append(result, s[i:i+3])
		}
	}
	return result
}

func reverseAba(s string) string {
	return fmt.Sprintf("%s%s%s", string(s[1]), string(s[0]), string(s[1]))
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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
		tlsIps := 0
		for _, line := range lines {
			parts := strings.FieldsFunc(line, func(ch rune) bool {
				return ch == '[' || ch == ']'
			})
			tls := false
			for i := 0; i < len(parts); i += 2 {
				if isAbba(parts[i]) {
					tls = true
				}
			}
			for i := 1; i < len(parts); i += 2 {
				if isAbba(parts[i]) {
					tls = false
				}
			}
			if tls {
				tlsIps++
			}
		}
		log.Printf("Found %d IPs that support TLS", tlsIps)
	case 2:
		lines := strings.Split(string(data), "\n")
		sslIps := 0
		for _, line := range lines {
			parts := strings.FieldsFunc(line, func(ch rune) bool {
				return ch == '[' || ch == ']'
			})
			superAbas := make([]string, 0)
			for i := 0; i < len(parts); i += 2 {
				superAbas = append(superAbas, findAbas(parts[i])...)
			}
			if len(superAbas) == 0 {
				continue
			}
			hyperAbas := make([]string, 0)
			for i := 1; i < len(parts); i += 2 {
				hyperAbas = append(hyperAbas, findAbas(parts[i])...)
			}
			if len(hyperAbas) == 0 {
				continue
			}
			for _, superAba := range superAbas {
				if stringInSlice(reverseAba(superAba), hyperAbas) {
					sslIps++
					break
				}
			}
		}
		log.Printf("Found %d IPs that support SSL", sslIps)
	}
}
