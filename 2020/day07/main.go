package main

import (
	"fmt"
	"strings"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := aoc.FileLines("input")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

type Bag1 struct {
	Color   string
	Parents []string
}

func part1(lines []string) int {
	var split []string
	var bags = make(map[string]Bag1)

	// Get all the bags
	for _, line := range lines {
		split = strings.SplitN(line, " bags contain ", 2)
		bags[split[0]] = Bag1{Color: split[0]}
	}
	// Fill in the parents
	for _, line := range lines {
		split = strings.SplitN(line, " bags contain ", 2)
		parentColor := split[0]

		if !strings.HasSuffix(line, "no other bags.") {
			for _, bag := range strings.Split(split[1], ", ") {
				split = strings.SplitN(bag, " ", 2)

				color := strings.TrimSuffix(
					strings.TrimRight(split[1], "s,."),
					" bag",
				)
				// fmt.Printf("%-26s %-26s %-26s\n", "'"+parentColor+"'", "'"+split[1]+"'", "'"+color+"'")

				bagCopy := bags[color]
				bagCopy.Parents = append(bagCopy.Parents, parentColor)
				bags[color] = bagCopy
			}
		}
	}

	var bagCount = make(map[string]int)
	countBagParents(bags, bags["shiny gold"], bagCount)
	return len(bagCount) - 1
}

func countBagParents(bags map[string]Bag1, bag Bag1, bagCount map[string]int) {
	bagCount[bag.Color] = bagCount[bag.Color] + 1
	for _, v := range bag.Parents {
		countBagParents(bags, bags[v], bagCount)
	}
}

type Bag2 struct {
	Color string

	Children []string
	Counts   []int
}

func part2(lines []string) int {
	var split []string
	var bags = make(map[string]Bag2)

	// Get all the bags
	for _, line := range lines {
		split = strings.SplitN(line, " bags contain ", 2)
		bag := Bag2{Color: split[0]}

		if !strings.HasSuffix(line, "no other bags.") {
			for _, bagStr := range strings.Split(split[1], ", ") {
				split = strings.SplitN(bagStr, " ", 2)

				count := aoc.AtoiOrBust(split[0])
				color := strings.TrimSuffix(
					strings.TrimRight(split[1], "s,."),
					" bag",
				)

				bag.Children = append(bag.Children, color)
				bag.Counts = append(bag.Counts, count)
			}
		}

		bags[bag.Color] = bag
	}

	return countBagChildren(bags, bags["shiny gold"])
}

func countBagChildren(bags map[string]Bag2, bag Bag2) (total int) {
	for k, color := range bag.Children {
		total += bag.Counts[k]
		total += countBagChildren(bags, bags[color]) * bag.Counts[k]
	}
	return total
}
