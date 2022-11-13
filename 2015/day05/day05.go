package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
	"go.arsenm.dev/pcre"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(lines []string) int {
	niceVowels := pcre.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	niceDouble := pcre.MustCompile(`([a-z])\1`)
	naughtyDoubles := pcre.MustCompile(`ab|cd|pq|xy`)

	total := 0
	for _, s := range lines {
		if niceVowels.MatchString(s) && niceDouble.MatchString(s) && !naughtyDoubles.MatchString(s) {
			total++
		}
	}

	return total
}

func part2(lines []string) int {
	nicePair := pcre.MustCompile(`([a-z]{2}).*\1`)
	niceRepeat := pcre.MustCompile(`([a-z]).\1`)

	total := 0
	for _, s := range lines {
		if nicePair.MatchString(s) && niceRepeat.MatchString(s) {
			total++
		}
	}
	return total
}
