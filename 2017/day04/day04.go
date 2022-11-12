package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(lines []string) int {
	valid := 0
line:
	for _, l := range lines {
		w := strings.Fields(l)
		for i := 0; i < len(w); i++ {
			for j := i + 1; j < len(w); j++ {
				if w[i] == w[j] {
					continue line
				}
			}
		}
		valid++
	}
	return valid
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func part2(lines []string) int {
	valid := 0
line:
	for _, l := range lines {
		w := strings.Fields(l)
		sw := make([]string, len(w))
		for i, s := range w {
			sw[i] = SortString(s)
		}
		for i := 0; i < len(sw); i++ {
			for j := i + 1; j < len(sw); j++ {
				if sw[i] == sw[j] {
					continue line
				}
			}
		}
		valid++
	}
	return valid
}
