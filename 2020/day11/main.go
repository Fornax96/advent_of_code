package main

import (
	"fmt"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())

	var lines = aoc.FileLines("input")
	var floormap = make([][]rune, len(lines))
	for k, v := range lines {
		floormap[k] = []rune(v)
	}

	fmt.Println("Answer 1:", part1(floormap))
	fmt.Println("Answer 2:", part2(floormap))
}

func part1(origMap [][]rune) (totalOccupied int) {
	var newMap [][]rune
	for {
		changes := 0
		newMap = mapCopy(origMap)
		for y, row := range origMap {
			for x, seat := range row {
				_, occupied, _ := neighbours(origMap, x, y)

				if seat == 'L' && occupied == 0 {
					newMap[y][x] = '#'
					changes++
				} else if seat == '#' && occupied > 3 {
					newMap[y][x] = 'L'
					changes++
				}
			}
		}

		origMap = newMap

		if changes == 0 {
			break
		}
	}

	for _, row := range newMap {
		for _, seat := range row {
			if seat == '#' {
				totalOccupied++
			}
		}
	}
	return totalOccupied
}

func part2(origMap [][]rune) (totalOccupied int) {
	var newMap [][]rune
	for {
		changes := 0
		newMap = mapCopy(origMap)
		for y, row := range origMap {
			for x, seat := range row {
				_, occupied, _ := neighbours2(origMap, x, y)

				if seat == 'L' && occupied == 0 {
					newMap[y][x] = '#'
					changes++
				} else if seat == '#' && occupied > 4 {
					newMap[y][x] = 'L'
					changes++
				}
			}
		}

		origMap = newMap

		if changes == 0 {
			break
		}
	}

	for _, row := range newMap {
		for _, seat := range row {
			if seat == '#' {
				totalOccupied++
			}
		}
	}
	return totalOccupied
}

func neighbours(floormap [][]rune, xPos, yPos int) (empty, occupied, floor int) {
	for y := yPos - 1; y <= yPos+1; y++ {
		for x := xPos - 1; x <= xPos+1; x++ {
			if x == xPos && y == yPos {
				continue // Don't count the starting point
			}
			if y < len(floormap) && y >= 0 && x < len(floormap[y]) && x >= 0 {
				switch floormap[y][x] {
				case 'L':
					empty++
				case '#':
					occupied++
				case '.':
					floor++
				}
			}
		}
	}
	return empty, occupied, floor
}

func neighbours2(floormap [][]rune, xPos, yPos int) (empty, occupied, floor int) {
	var modifiers = []struct {
		x int
		y int
	}{
		{-1, -1}, {-1, 0}, {-1, +1},
		{0, -1}, {0, +1},
		{+1, -1}, {+1, 0}, {+1, +1},
	}

	for _, mod := range modifiers {
		var x, y = xPos, yPos
		for {
			x += mod.x
			y += mod.y

			if y < len(floormap) && y >= 0 && x < len(floormap[y]) && x >= 0 {
				switch floormap[y][x] {
				case 'L':
					empty++
				case '#':
					occupied++
				case '.':
					floor++
					continue
				}
			}
			break
		}
	}
	return empty, occupied, floor
}

func mapCopy(origMap [][]rune) (newMap [][]rune) {
	newMap = make([][]rune, len(origMap))
	for k := range newMap {
		newMap[k] = make([]rune, len(origMap[k]))
		copy(newMap[k], origMap[k])
	}
	return newMap
}
