package aoc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// ScannerCloser is a scanner with a Close method to close the underlying file
type ScannerCloser struct {
	*bufio.Scanner
	io.Closer
}

// FileScanner opens a file and wraps it in a scanner. Close the file after you
// are done with the returned io.Closer
func FileScanner(p string) ScannerCloser {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	return ScannerCloser{bufio.NewScanner(f), f}
}

// FileLines returns all the lines in a file in a slice
func FileLines(p string) (lines []string) {
	f := FileScanner(p)
	for f.Scan() {
		lines = append(lines, f.Text())
	}
	f.Close()
	return lines
}

// FileInts reads all the lines of a file, parses them into ints and returns
// them as a slice
func FileInts(p string) (ints []int) {
	f := FileScanner(p)
	for f.Scan() {
		ints = append(ints, AtoiOrBust(f.Text()))
	}
	f.Close()
	return ints
}

// AtoiOrBust converts a string to an integer, or it panics
func AtoiOrBust(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

// PrintDuration prints the time since t. Call this function like this and it
// will print the run time of a function after it ends:
//
//  defer aoc.PrintDuration(time.Now())
//
// This works because when you use 'defer' the arguments of a function will be
// evaluated immediately, while the function itself will only be evaluated after
// the parent function ends.
func PrintDuration(t time.Time) {
	fmt.Println(time.Since(t))
}
