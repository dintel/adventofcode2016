package day9

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func decompress(data []byte) []byte {
	buffer := bytes.NewBuffer(data)
	outBuffer := bytes.NewBufferString("")
	ch, err := buffer.ReadByte()
	lengthBuffer := bytes.NewBufferString("")
	nBuffer := bytes.NewBufferString("")
	tokenBuffer := bytes.NewBufferString("")
	var length, n int
	var token string
	for err == nil {
		if ch == '(' {
			ch, _ = buffer.ReadByte()
			for ch != 'x' {
				lengthBuffer.WriteByte(ch)
				ch, _ = buffer.ReadByte()
			}
			ch, _ = buffer.ReadByte()
			for ch != ')' {
				nBuffer.WriteByte(ch)
				ch, _ = buffer.ReadByte()
			}
			length, _ = strconv.Atoi(lengthBuffer.String())
			n, _ = strconv.Atoi(nBuffer.String())
			for i := 0; i < length; i++ {
				ch, _ = buffer.ReadByte()
				tokenBuffer.WriteByte(ch)
			}
			token = tokenBuffer.String()
			for i := 0; i < n; i++ {
				outBuffer.WriteString(token)
			}
			lengthBuffer.Reset()
			nBuffer.Reset()
			tokenBuffer.Reset()
		} else {
			outBuffer.WriteByte(ch)
		}
		ch, err = buffer.ReadByte()
	}
	return outBuffer.Bytes()
}

func decompress2(data []byte) int {
	buffer := bytes.NewBuffer(data)
	result := 0
	ch, err := buffer.ReadByte()
	lengthBuffer := bytes.NewBufferString("")
	nBuffer := bytes.NewBufferString("")
	tokenBuffer := bytes.NewBufferString("")
	var length, n int
	var tokenLength int
	for err == nil {
		if ch == '(' {
			ch, _ = buffer.ReadByte()
			for ch != 'x' {
				lengthBuffer.WriteByte(ch)
				ch, _ = buffer.ReadByte()
			}
			ch, _ = buffer.ReadByte()
			for ch != ')' {
				nBuffer.WriteByte(ch)
				ch, _ = buffer.ReadByte()
			}
			length, _ = strconv.Atoi(lengthBuffer.String())
			n, _ = strconv.Atoi(nBuffer.String())
			for i := 0; i < length; i++ {
				ch, _ = buffer.ReadByte()
				tokenBuffer.WriteByte(ch)
			}
			tokenLength = decompress2(tokenBuffer.Bytes())
			result += tokenLength * n
			lengthBuffer.Reset()
			nBuffer.Reset()
			tokenBuffer.Reset()
		} else {
			result++
		}
		ch, err = buffer.ReadByte()
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

	switch part {
	case 1:
		out := decompress(data)
		log.Printf("Encoded string length is %d", len(data))
		log.Printf("Decoded string length is %d", len(out))
	case 2:
		log.Printf("Encoded string length is %d", len(data))
		log.Printf("Decoded string length is %d", decompress2(data))
	}
}
