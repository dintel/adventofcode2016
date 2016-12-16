package day16

import (
	"io/ioutil"
	"log"
	"os"
)

func inv(n byte) byte {
	if n == 0 {
		return 1
	}
	return 0
}

func dragonCurve(in []byte) []byte {
	result := make([]byte, len(in)*2+1)
	for i := range in {
		result[i] = in[i]
	}
	result[len(in)] = 0
	r := len(in) + 1
	for i := len(in) - 1; i >= 0; i-- {
		result[r] = inv(in[i])
		r++
	}
	return result
}

func checksum(in []byte) []byte {
	result := make([]byte, len(in)/2)
	for i := 0; i < len(in); i += 2 {
		result[i/2] = inv(in[i] ^ in[i+1])
	}
	if len(result)%2 == 0 {
		return checksum(result)
	}
	return result
}

func checksumString(sum []byte) string {
	result := make([]byte, len(sum))
	for i := range sum {
		if sum[i] == 0 {
			result[i] = '0'
		} else {
			result[i] = '1'
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
	input := make([]byte, len(data))
	for i := range data {
		if data[i] == '0' {
			input[i] = 0
		} else {
			input[i] = 1
		}
	}

	switch part {
	case 1:
		diskLen := 272
		disk := dragonCurve(input)
		for len(disk) < diskLen {
			disk = dragonCurve(disk)
		}
		disk = disk[:diskLen]
		log.Printf("Checksum of disk is %s", checksumString(checksum(disk)))
	case 2:
		diskLen := 35651584
		disk := dragonCurve(input)
		for len(disk) < diskLen {
			disk = dragonCurve(disk)
		}
		disk = disk[:diskLen]
		log.Printf("Checksum of disk is %s", checksumString(checksum(disk)))
	}
}
