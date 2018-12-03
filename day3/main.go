package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main1() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	reg := regexp.MustCompile(`(\d+),(\d+): (\d+)x(\d+)`)
	var matrix [1000][1000]int
	var dupcount = 0

	for s.Scan() {
		nums := reg.FindStringSubmatch(s.Text())
		x, _ := strconv.Atoi(nums[1])
		y, _ := strconv.Atoi(nums[2])
		width, _ := strconv.Atoi(nums[3])
		height, _ := strconv.Atoi(nums[4])

		for ycursor := y; ycursor < y+height; ycursor++ {
			for xcursor := x; xcursor < x+width; xcursor++ {
				if matrix[xcursor][ycursor] == 1 {
					dupcount++
				}
				matrix[xcursor][ycursor]++
			}
		}
	}

	println(dupcount)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	reg := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	var matrix [1000][1000]int
	var overlapping = make(map[int]bool)

	for s.Scan() {
		nums := reg.FindStringSubmatch(s.Text())
		id, _ := strconv.Atoi(nums[1])
		x, _ := strconv.Atoi(nums[2])
		y, _ := strconv.Atoi(nums[3])
		width, _ := strconv.Atoi(nums[4])
		height, _ := strconv.Atoi(nums[5])

		overlapping[id] = false // Add the ID to the map

		for ycursor := y; ycursor < y+height; ycursor++ {
			for xcursor := x; xcursor < x+width; xcursor++ {
				if otherid := matrix[xcursor][ycursor]; otherid != 0 {
					overlapping[otherid] = true // the other id was overlapped by this id
					overlapping[id] = true      // this id overlaps another id
				}
				matrix[xcursor][ycursor] = id
			}
		}
	}

	for k, v := range overlapping {
		if v == false {
			println(k)
		}
	}
}
