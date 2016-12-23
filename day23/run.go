package day23

import (
	"fmt"
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
	operand3  string
	operand4  string
}

func (i instruction) String() string {
	switch i.operation {
	case 1:
		return fmt.Sprintf("cpy %s %s", i.operand1, i.operand2)
	case 2:
		return fmt.Sprintf("inc %s", i.operand1)
	case 3:
		return fmt.Sprintf("dec %s", i.operand1)
	case 4:
		return fmt.Sprintf("jnz %s %s", i.operand1, i.operand2)
	case 5:
		return fmt.Sprintf("tgl %s", i.operand1)
	case 6:
		return fmt.Sprintf("mul %s %s %s %s", i.operand1, i.operand2, i.operand3, i.operand4)
	case 7:
		return fmt.Sprintf("nil")
	}
	return "Unknown instruction"
}

func (p *processor) optimizeMultiply() {
	max := len(p.instructions) - 5
	muls := make([]int, 0)
	for i := 0; i < max; i++ {
		if p.instructions[i].operation == 1 && p.instructions[i+1].operation == 2 && p.instructions[i+2].operation == 3 && p.instructions[i+3].operation == 4 && p.instructions[i+4].operation == 3 && p.instructions[i+5].operation == 4 && p.instructions[i+2].operand1 == p.instructions[i+3].operand1 && p.instructions[i+4].operand1 == p.instructions[i+5].operand1 {
			muls = append(muls, i)
		}
	}
	for i := len(muls) - 1; i >= 0; i-- {
		idx := muls[i]
		p.instructions[idx] = instruction{
			operation: 6,
			operand1:  p.instructions[idx+1].operand1,
			operand2:  p.instructions[idx].operand1,
			operand3:  p.instructions[idx+5].operand1,
			operand4:  p.instructions[idx+3].operand1,
		}
		p.instructions[idx+1].operation = 7
		p.instructions[idx+2].operation = 7
		p.instructions[idx+3].operation = 7
		p.instructions[idx+4].operation = 7
		p.instructions[idx+5].operation = 7
	}
	log.Printf("Optimized %d loops with multiply", len(muls))
}

type processor struct {
	instructions []instruction
	registers    map[string]int
	pc           int
}

func newProcessor(instructions []instruction, registers int) *processor {
	result := &processor{
		instructions: instructions,
		registers:    make(map[string]int),
		pc:           0,
	}
	last := rune('a' + registers)
	for i := 'a'; i < last; i++ {
		result.registers[string(i)] = 0
	}
	return result
}

func (p *processor) simpleLoop(j int) bool {
	if j >= 0 {
		return false
	}
	j = -j
	for i := 1; i <= j; i++ {
		if p.instructions[p.pc-i].operation != 2 && p.instructions[p.pc-i].operation != 3 {
			return false
		}
	}
	return true
}

func (p *processor) optimize(j int) {
	// p.pc += j
	// return
	jnz := p.instructions[p.pc]
	m := make(map[string]int)
	j = -j
	for i := 1; i <= j; i++ {
		instr := p.instructions[p.pc-i]
		if _, err := strconv.Atoi(instr.operand1); err == nil {
			continue
		}
		if _, exists := m[instr.operand1]; !exists {
			m[instr.operand1] = 0
		}
		switch instr.operation {
		case 2:
			m[instr.operand1]++
		case 3:
			m[instr.operand1]--
		default:
			log.Fatalf("Illegal instruction during simple loop optimization - %s", instr)
		}
	}
	mul := p.registers[jnz.operand1] / -m[jnz.operand1]
	for r, v := range m {
		p.registers[r] += v * mul
	}
	p.pc++
}

func (p *processor) exec() {
	i := p.instructions[p.pc]
	switch i.operation {
	case 1:
		n, err := strconv.Atoi(i.operand1)
		if err != nil {
			n = p.registers[i.operand1]
		}
		_, err = strconv.Atoi(i.operand2)
		if err != nil {
			p.registers[i.operand2] = n
		}
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
			jump, err := strconv.Atoi(i.operand2)
			if err != nil {
				jump = p.registers[i.operand2]
			}
			if p.simpleLoop(jump) {
				p.optimize(jump)
			} else {
				p.pc += jump
			}
		} else {
			p.pc++
		}
	case 5:
		n, err := strconv.Atoi(i.operand1)
		if err != nil {
			n = p.registers[i.operand1]
		}
		idx := p.pc + n
		if idx < len(p.instructions) && idx >= 0 {
			if p.instructions[idx].operand2 == "" {
				if p.instructions[idx].operation == 2 {
					p.instructions[idx].operation = 3
				} else {
					p.instructions[idx].operation = 2
				}
			} else {
				if p.instructions[idx].operation == 4 {
					p.instructions[idx].operation = 1
				} else {
					p.instructions[idx].operation = 4
				}
			}
		}
		p.pc++
	case 6:
		p.registers[i.operand1] += p.registers[i.operand2] * p.registers[i.operand3]
		p.registers[i.operand3] = 0
		p.registers[i.operand4] = 0
		p.pc++
	case 7:
		p.pc++
	}
}

func (p *processor) run() {
	for p.pc < len(p.instructions) && p.pc >= 0 {
		p.exec()
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
		case "tgl":
			instructions = append(instructions, instruction{
				operation: 5,
				operand1:  parts[1],
			})
		default:
			log.Fatalf("Unknown instruction %s", parts[0])
		}
	}
	p := newProcessor(instructions, 4)

	switch part {
	case 1:
		p.registers["a"] = 7
		p.run()
		log.Printf("After running program register a value is %d", p.registers["a"])
	case 2:
		p.registers["a"] = 12
		p.optimizeMultiply()
		p.run()
		log.Printf("After running program register a value is %d", p.registers["a"])
	}
}
