package util

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

// Lines returns the contents of a file as a slice of lines
func Lines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([]string, 0, 20)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// IntLines returns the contents of a file as a slice of ints
func IntLines(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ints := make([]int, 0, 20)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, x)
	}

	return ints
}

// ScanRecords is a bufio.SplitFunc to split on two newlines
func ScanRecords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), bytes.TrimRight(data, "\n"), nil
	}
	return 0, nil, nil
}

// Records returns the contents of a file split by instances of two newlines
func Records(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	records := make([]string, 0, 20)
	scanner := bufio.NewScanner(f)
	scanner.Split(ScanRecords)
	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	return records
}

// FirstRune returns the first rune in the string
func FirstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}
