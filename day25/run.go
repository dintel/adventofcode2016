package day25

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation int
	operand1  string
	operand2  string
}

type processor struct {
	instructions []instruction
	registers    map[string]int
	pc           int
	output       []int
}

func newProcessor(instructions []instruction, registers int) *processor {
	result := &processor{
		instructions: instructions,
		registers:    make(map[string]int),
		pc:           0,
		output:       make([]int, 0),
	}
	last := rune('a' + registers)
	for i := 'a'; i < last; i++ {
		result.registers[string(i)] = 0
	}
	return result
}

func (p *processor) reset() {
	for r := range p.registers {
		p.registers[r] = 0
	}
	p.pc = 0
	p.output = make([]int, 0)
}

func (p *processor) exec() {
	i := p.instructions[p.pc]
	switch i.operation {
	case 1:
		n, err := strconv.Atoi(i.operand1)
		if err != nil {
			n = p.registers[i.operand1]
		}
		p.registers[i.operand2] = n
		p.pc++
	case 2:
		p.registers[i.operand1]++
		p.pc++
	case 3:
		p.registers[i.operand1]--
		p.pc++
	case 4:
		n, err := strconv.Atoi(i.operand1)
		if err != nil {
			n = p.registers[i.operand1]
		}
		if n != 0 {
			jump, _ := strconv.Atoi(i.operand2)
			p.pc += jump
		} else {
			p.pc++
		}
	case 5:
		if p.registers[i.operand1] != 0 && p.registers[i.operand1] != 1 {
			p.pc = -1
			return
		}
		if len(p.output) > 0 {
			if p.output[len(p.output)-1] == 0 && p.registers[i.operand1] != 1 {
				p.pc = -1
				return
			}
			if p.output[len(p.output)-1] == 1 && p.registers[i.operand1] != 0 {
				p.pc = -1
				return
			}
		}
		p.output = append(p.output, p.registers[i.operand1])
		p.pc++
	}
}

func (p *processor) run(n int) {
	for p.pc < len(p.instructions) && p.pc >= 0 && n > 0 {
		p.exec()
		n--
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

	lines := strings.Split(string(data), "\n")
	instructions := make([]instruction, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "cpy":
			instructions = append(instructions, instruction{
				operation: 1,
				operand1:  parts[1],
				operand2:  parts[2],
			})
		case "inc":
			instructions = append(instructions, instruction{
				operation: 2,
				operand1:  parts[1],
			})
		case "dec":
			instructions = append(instructions, instruction{
				operation: 3,
				operand1:  parts[1],
			})
		case "jnz":
			instructions = append(instructions, instruction{
				operation: 4,
				operand1:  parts[1],
				operand2:  parts[2],
			})
		case "out":
			instructions = append(instructions, instruction{
				operation: 5,
				operand1:  parts[1],
			})
		}
	}
	p := newProcessor(instructions, 4)

	switch part {
	case 1:
		a := 0
		valid := false
		for !valid {
			a++
			p.reset()
			p.registers["a"] = a
			p.run(100000)
			if p.pc >= 0 && p.pc < len(p.instructions) {
				log.Printf("Found candidate a=%d", a)
				valid = true
			}
		}
	case 2:
		log.Printf("Part 2 has no answer")
	}
}
