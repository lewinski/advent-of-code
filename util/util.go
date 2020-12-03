package util

import (
	"bufio"
	"os"
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

// FirstRune returns the first rune in the string
func FirstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}
