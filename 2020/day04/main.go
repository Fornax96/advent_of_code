package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aoc "github.com/Fornax96/advent_of_code"
)

func main() {
	defer aoc.PrintDuration(time.Now())
	lines := aoc.FileLines("input")
	lines = append(lines, "")
	fmt.Println("Answer 1:", part1(lines))
	fmt.Println("Answer 2:", part2(lines))
}

func part1(lines []string) (valid int) {
	var fields = make(map[string]bool)
	var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, line := range lines {
		if line == "" {
			allValid := true
			for _, field := range required {
				if !fields[field] {
					allValid = false
					break
				}
			}
			if allValid {
				valid++
			}
			fields = make(map[string]bool) // Reset the map
			continue
		}

		for _, field := range required {
			if strings.Contains(line, field+":") {
				fields[field] = true
			}
		}
	}

	return valid
}

func part2(lines []string) (valid int) {
	var fields = make(map[string]string)

	for _, line := range lines {
		if line == "" {
			if validateInt(fields["byr"], 1920, 2002) &&
				validateInt(fields["iyr"], 2010, 2020) &&
				validateInt(fields["eyr"], 2020, 2030) &&
				validateHeight(fields["hgt"]) &&
				validateHair(fields["hcl"]) &&
				validateEye(fields["ecl"]) &&
				validatePid(fields["pid"]) {
				valid++
			}
			fields = make(map[string]string) // Reset the map
			continue
		}

		couples := strings.Split(line, " ")
		for _, couple := range couples {
			kv := strings.Split(couple, ":")
			fields[kv[0]] = kv[1]
		}
	}

	return valid
}

func validateInt(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}
func validateHeight(s string) bool {
	if strings.HasSuffix(s, "cm") {
		i := aoc.AtoiOrBust(strings.TrimRight(s, "cm"))
		return i >= 150 && i <= 193
	}
	if strings.HasSuffix(s, "in") {
		i := aoc.AtoiOrBust(strings.TrimRight(s, "in"))
		return i >= 59 && i <= 76
	}
	return false
}
func validateHair(s string) bool {
	if len(s) != 7 || !strings.HasPrefix(s, "#") {
		return false
	}
	for _, r := range strings.TrimPrefix(s, "#") {
		if !strings.ContainsRune("0123456789abcdef", r) {
			return false
		}
	}
	return true
}
func validateEye(s string) bool {
	return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
}
func validatePid(s string) bool {
	if len(s) != 9 {
		return false
	}
	return validateInt(s, 0, 999999999)
}
