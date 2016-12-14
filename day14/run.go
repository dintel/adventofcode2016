package day14

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type pad struct {
	h       hash.Hash
	hashes  map[string]string
	secret  string
	counter int
	stretch bool
}

func newPad(secret string, stretch bool) *pad {
	return &pad{
		h:       md5.New(),
		hashes:  make(map[string]string),
		secret:  secret,
		counter: 1,
		stretch: stretch,
	}
}

func (p *pad) getHash(i int) string {
	src := fmt.Sprintf("%s%d", p.secret, i)
	sum, exists := p.hashes[src]
	if !exists {
		_, err := io.WriteString(p.h, src)
		if err != nil {
			log.Fatalf("Error - %s", err)
		}
		sum = fmt.Sprintf("%x", p.h.Sum(nil))
		p.hashes[src] = sum
		p.h.Reset()
		if p.stretch {
			current := sum
			for i := 0; i < 2016; i++ {
				_, err := io.WriteString(p.h, current)
				if err != nil {
					log.Fatalf("Error - %s", err)
				}
				current = fmt.Sprintf("%x", p.h.Sum(nil))
				p.h.Reset()
			}
			p.hashes[src] = current
			sum = current
		}
	}
	return sum
}

func (p *pad) checkCurrentTriple() byte {
	sum := p.getHash(p.counter)
	for i := 0; i < len(sum)-2; i++ {
		if sum[i] == sum[i+1] && sum[i+1] == sum[i+2] {
			return sum[i]
		}
	}
	return 0
}

func (p *pad) checkNextThousand(b byte) bool {
	for i := 1; i < 1001; i++ {
		sum := p.getHash(p.counter + i)
		for i := 0; i < len(sum)-4; i++ {
			if sum[i] == b && sum[i] == sum[i+1] && sum[i+1] == sum[i+2] && sum[i+2] == sum[i+3] && sum[i+3] == sum[i+4] {
				return true
			}
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
		p := newPad(string(data), false)
		found := 0
		for found != 64 {
			b := p.checkCurrentTriple()
			if b != 0 && p.checkNextThousand(b) {
				found++
			}
			p.counter++
		}
		log.Printf("Found last key at index %d", p.counter-1)
	case 2:
		p := newPad(string(data), true)
		found := 0
		for found != 64 {
			b := p.checkCurrentTriple()
			if b != 0 && p.checkNextThousand(b) {
				found++
			}
			p.counter++
		}
		log.Printf("Found last key at index %d", p.counter-1)
	}
}
