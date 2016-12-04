package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dintel/adventofcode2016/day1"
	"github.com/dintel/adventofcode2016/day10"
	"github.com/dintel/adventofcode2016/day11"
	"github.com/dintel/adventofcode2016/day12"
	"github.com/dintel/adventofcode2016/day13"
	"github.com/dintel/adventofcode2016/day14"
	"github.com/dintel/adventofcode2016/day15"
	"github.com/dintel/adventofcode2016/day16"
	"github.com/dintel/adventofcode2016/day17"
	"github.com/dintel/adventofcode2016/day18"
	"github.com/dintel/adventofcode2016/day19"
	"github.com/dintel/adventofcode2016/day2"
	"github.com/dintel/adventofcode2016/day20"
	"github.com/dintel/adventofcode2016/day21"
	"github.com/dintel/adventofcode2016/day22"
	"github.com/dintel/adventofcode2016/day23"
	"github.com/dintel/adventofcode2016/day24"
	"github.com/dintel/adventofcode2016/day25"
	"github.com/dintel/adventofcode2016/day3"
	"github.com/dintel/adventofcode2016/day4"
	"github.com/dintel/adventofcode2016/day5"
	"github.com/dintel/adventofcode2016/day6"
	"github.com/dintel/adventofcode2016/day7"
	"github.com/dintel/adventofcode2016/day8"
	"github.com/dintel/adventofcode2016/day9"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <day-number> <part-number> parameters\n", os.Args[0])
	}

	// Parse day number and validate it
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Could not parse day number: %s", err)
	}
	if day < 1 || day > 25 {
		log.Fatalln("Day number must be between 1 and 25")
	}

	// Parse part number and validate it
	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Could not parse part number: %s", err)
	}
	if part < 1 || part > 2 {
		log.Fatalln("Part number must be erither 1 or 2")
	}

	switch day {
	case 1:
		day1.Run(part)
	case 2:
		day2.Run(part)
	case 3:
		day3.Run(part)
	case 4:
		day4.Run(part)
	case 5:
		day5.Run(part)
	case 6:
		day6.Run(part)
	case 7:
		day7.Run(part)
	case 8:
		day8.Run(part)
	case 9:
		day9.Run(part)
	case 10:
		day10.Run(part)
	case 11:
		day11.Run(part)
	case 12:
		day12.Run(part)
	case 13:
		day13.Run(part)
	case 14:
		day14.Run(part)
	case 15:
		day15.Run(part)
	case 16:
		day16.Run(part)
	case 17:
		day17.Run(part)
	case 18:
		day18.Run(part)
	case 19:
		day19.Run(part)
	case 20:
		day20.Run(part)
	case 21:
		day21.Run(part)
	case 22:
		day22.Run(part)
	case 23:
		day23.Run(part)
	case 24:
		day24.Run(part)
	case 25:
		day25.Run(part)
	}
	end := time.Now()
	duration := end.Sub(start)
	log.Printf("Runtime: %s", duration)
}
