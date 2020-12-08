package main

import (
	"fmt"
	"strings"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

type instruction struct {
	op  string
	arg int
}

func main() {
	defer aoc.PrintDuration(time.Now())

	lines := aoc.FileLines("input")
	var i = make([]instruction, len(lines))
	var split []string
	for k, v := range lines {
		split = strings.SplitN(v, " ", 2)
		i[k].op, i[k].arg = split[0], aoc.AtoiOrBust(split[1])
	}

	fmt.Println("Answer 1:", part1(i))
	fmt.Println("Answer 2:", part2(i))
}

func part1(lines []instruction) int {
	var (
		acc int
		i   int // Instruction index
		log []int
	)
	for {
		for _, n := range log {
			if n == i {
				return acc
			}
		}
		log = append(log, i)

		switch lines[i].op {
		case "acc":
			acc += lines[i].arg
		case "jmp":
			i += lines[i].arg
			continue
		case "nop":
		default:
			panic("unknown instruction")
		}
		i++
	}
}

func part2(lines []instruction) int {
	var linesCopy = make([]instruction, len(lines))
	acc, execCount, ok := exec(lines)

	// Get all the nop and jmp instructions in the loop, then swap them out and
	// check if it solves the problem
	for ins, count := range execCount {
		if count == 100 && (lines[ins].op == "jmp" || lines[ins].op == "nop") {
			// Make a copy of the instruction slice
			copy(linesCopy, lines)

			// Change one instruction
			if linesCopy[ins].op == "jmp" {
				linesCopy[ins].op = "nop"
			} else {
				linesCopy[ins].op = "jmp"
			}

			// Check if this change fixed the program
			if acc, _, ok = exec(linesCopy); ok {
				return acc
			}
		}
	}
	panic("correct value not found")
}

func exec(ins []instruction) (acc int, execCount map[int]int, ok bool) {
	var i int // Instruction index
	execCount = make(map[int]int)

	for {
		if i == len(ins) {
			return acc, execCount, true
		}

		// Detect the loop and return the number of times each instruction was
		// executed
		if execCount[i] == 100 {
			return acc, execCount, false
		}
		execCount[i]++

		switch ins[i].op {
		case "acc":
			acc += ins[i].arg
		case "jmp":
			i += ins[i].arg
			continue
		case "nop":
		default:
			panic("unknown instruction")
		}
		i++
	}
}
