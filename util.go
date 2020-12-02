package aoc

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
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

// AtoiOrBust converts a string to an integer, or it panics
func AtoiOrBust(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}
