package day10

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type bot struct {
	chips       []int
	chipsPassed []int
	lowOutput   bool
	low         int
	highOutput  bool
	high        int
}

func newBot(low, high int, lowOutput, highOutput bool) *bot {
	return &bot{
		chips:       make([]int, 0),
		chipsPassed: make([]int, 0),
		low:         low,
		lowOutput:   lowOutput,
		high:        high,
		highOutput:  highOutput,
	}
}

func (b *bot) process() (lowN, lowV, highN, highV int, outputLow, outputHigh bool) {
	sort.Ints(b.chips)
	lowV = b.chips[0]
	highV = b.chips[1]
	lowN = b.low
	highN = b.high
	outputLow = b.lowOutput
	outputHigh = b.highOutput

	b.chips = make([]int, 0)
	return
}

func (b *bot) add(v int) {
	b.chips = append(b.chips, v)
	b.chipsPassed = append(b.chipsPassed, v)
	sort.Ints(b.chipsPassed)
}

func (b *bot) ready() bool {
	return len(b.chips) == 2
}

type factory struct {
	bots    map[int]*bot
	outputs map[int]int
	inputs  map[int][]int
}

func newFactory() *factory {
	return &factory{
		bots:    make(map[int]*bot),
		outputs: make(map[int]int),
		inputs:  make(map[int][]int),
	}
}

func (f *factory) addBot(n, low, high int, lowOutput, highOutput bool) {
	f.bots[n] = newBot(low, high, lowOutput, highOutput)
}

func (f *factory) addInput(bot, value int) {
	if _, exists := f.inputs[bot]; !exists {
		f.inputs[bot] = make([]int, 0)
	}
	f.inputs[bot] = append(f.inputs[bot], value)
}

func (f *factory) compute() {
	for n, in := range f.inputs {
		for _, v := range in {
			f.bots[n].add(v)
		}
	}
	queue := make([]*bot, 0)
	for _, b := range f.bots {
		if b.ready() {
			queue = append(queue, b)
		}
	}
	var b *bot
	for len(queue) != 0 {
		b, queue = queue[0], queue[1:]
		lowN, lowV, highN, highV, outputLow, outputHigh := b.process()
		if outputLow {
			f.outputs[lowN] = lowV
		} else {
			f.bots[lowN].add(lowV)
			if f.bots[lowN].ready() {
				queue = append(queue, f.bots[lowN])
			}
		}
		if outputHigh {
			f.outputs[highN] = highV
		} else {
			f.bots[highN].add(highV)
			if f.bots[highN].ready() {
				queue = append(queue, f.bots[highN])
			}
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
		lines := strings.Split(string(data), "\n")
		f := newFactory()
		for _, line := range lines {
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "value":
				value, _ := strconv.Atoi(parts[1])
				bot, _ := strconv.Atoi(parts[5])
				f.addInput(bot, value)
			case "bot":
				bot, _ := strconv.Atoi(parts[1])
				lowOutput := parts[5] == "output"
				low, _ := strconv.Atoi(parts[6])
				highOutput := parts[10] == "output"
				high, _ := strconv.Atoi(parts[11])
				f.addBot(bot, low, high, lowOutput, highOutput)
			}
		}
		log.Printf("Parsed %d lines (%d bots, %d inputs)", len(lines), len(f.bots), len(f.inputs))
		f.compute()
		log.Print("Finished computing")
		for n, bot := range f.bots {
			if len(bot.chipsPassed) >= 2 && bot.chipsPassed[0] == 17 && bot.chipsPassed[1] == 61 {
				log.Printf("Found bot %d that processed chips 17 and 61", n)
				return
			}
		}
	case 2:
		lines := strings.Split(string(data), "\n")
		f := newFactory()
		for _, line := range lines {
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "value":
				value, _ := strconv.Atoi(parts[1])
				bot, _ := strconv.Atoi(parts[5])
				f.addInput(bot, value)
			case "bot":
				bot, _ := strconv.Atoi(parts[1])
				lowOutput := parts[5] == "output"
				low, _ := strconv.Atoi(parts[6])
				highOutput := parts[10] == "output"
				high, _ := strconv.Atoi(parts[11])
				f.addBot(bot, low, high, lowOutput, highOutput)
			}
		}
		log.Printf("Parsed %d lines (%d bots, %d inputs)", len(lines), len(f.bots), len(f.inputs))
		f.compute()
		log.Print("Finished computing")
		log.Printf("output0=%d, output1=%d, output2=%d, result=%d", f.outputs[0], f.outputs[1], f.outputs[2], f.outputs[0]*f.outputs[1]*f.outputs[2])
	}
}
