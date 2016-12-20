package day20

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rule struct {
	start int
	end   int
}

type rules []rule

func (rs rules) Len() int {
	return len(rs)
}

func (rs rules) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs rules) Less(i, j int) bool {
	return rs[i].start < rs[j].start
}

func (rs rules) Excluded(ip int) (bool, int) {
	for i, rule := range rs {
		if rule.start > ip {
			return false, 0
		} else if rule.start <= ip && rule.end >= ip {
			return true, i
		}
	}
	return false, 0
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
		ranges := rules(make([]rule, len(lines)))
		for i, line := range lines {
			parts := strings.Split(line, "-")
			ranges[i].start, _ = strconv.Atoi(parts[0])
			ranges[i].end, _ = strconv.Atoi(parts[1])
		}
		sort.Sort(ranges)
		ip := 0
		excluded, i := ranges.Excluded(ip)
		for excluded {
			ip = ranges[i].end + 1
			excluded, i = ranges.Excluded(ip)
		}
		log.Printf("Found not excluded IP %d", ip)
	case 2:
		lines := strings.Split(string(data), "\n")
		ranges := rules(make([]rule, len(lines)))
		for i, line := range lines {
			parts := strings.Split(line, "-")
			ranges[i].start, _ = strconv.Atoi(parts[0])
			ranges[i].end, _ = strconv.Atoi(parts[1])
		}
		sort.Sort(ranges)
		maxIp := 4294967295
		white := 0
		ip := 0
		for ip <= maxIp {
			ipExcluded, i := ranges.Excluded(ip)
			if ipExcluded {
				ip = ranges[i].end + 1
			} else {
				white++
				ip++
			}
		}
		log.Printf("Found %d white-listed IP addresses", white)
	}
}
