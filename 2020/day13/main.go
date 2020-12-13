package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())

	var lines = aoc.FileLines("input")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(lines []string) (ans int) {
	var startTime = aoc.AtoiOrBust(lines[0])
	var buses []int
	for _, v := range strings.Split(lines[1], ",") {
		if v != "x" {
			buses = append(buses, aoc.AtoiOrBust(v))
		}
	}

	for t := startTime; ; t++ {
		for _, bus := range buses {
			if t%bus == 0 {
				return bus * (t - startTime)
			}
		}
	}
}

type Bus struct {
	id     int
	offset int
}

func part2(lines []string) int {
	var buses []Bus

	for k, v := range strings.Split(lines[1], ",") {
		if v != "x" {
			buses = append(buses, Bus{aoc.AtoiOrBust(v), k})
		}
	}

	sort.Slice(buses, func(i, j int) bool {
		return buses[i].id > buses[j].id
	})

	var correction = buses[0].offset
	var lowestOffset = 999
	for k := range buses {
		buses[k].offset -= correction
		if buses[k].offset < lowestOffset {
			lowestOffset = buses[k].offset
		}
	}

	for _, bus := range buses {
		fmt.Printf("bus %d offset %d\n", bus.id, bus.offset)
	}

	var cont = 0
	var interval = buses[0].id

	for t := interval; ; t += interval {

		// First bus always matches. Check if the second bus also matches
		if (t+buses[1].offset)%buses[1].id == 0 {
			cont = 0

			for k := range buses {
				if (t+buses[k].offset)%buses[k].id == 0 {
					cont++
					continue
				}

				break
			}

			if cont == len(buses) {
				return t + lowestOffset
			}
		}
	}
}
