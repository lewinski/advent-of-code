package main

import (
	"fmt"
	"regexp"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	contains := 0
	overlaps := 0
	for _, line := range lines {
		f := regexp.MustCompile(`[-,]`).Split(line, -1)
		a := util.MustAtoi(f[0])
		b := util.MustAtoi(f[1])
		c := util.MustAtoi(f[2])
		d := util.MustAtoi(f[3])
		if intervalContains(a, b, c, d) {
			contains++
		}
		if intervalsOverlap(a, b, c, d) {
			overlaps++
		}
	}

	fmt.Println("part1:", contains)
	fmt.Println("part2:", overlaps)
}

// true if a-b contains c-d or c-d contains a-b
func intervalContains(a, b, c, d int) bool {
	if a >= c && b <= d {
		// c..a..b..d
		return true
	} else if c >= a && d <= b {
		// a..c..d..b
		return true
	}
	return false
}

// true if a-b overlaps c-d
func intervalsOverlap(a, b, c, d int) bool {
	if a <= c && b >= c {
		// abcd, acbd, acdb
		return true
	} else if c <= a && d >= a {
		// cabd, cadb, cdab
		return true
	}
	return false
}

// all possibilities if a<=b, c<=d
// case 1: a <= c
// abcd - overlap if b == c
// abdc - invalid
// acbd - overlap if b >= c
// acdb - contains & overlap
// adbc - invalid
// adcb - invalid
// all starting with b are invalid
// case 2: c <= a
// cabd - contains & overlap
// cadb - overlap if a <= d
// cbad - invalid
// cbda - invalid
// cdab - overlap if a == d
// cdba - invalid
// all starting with d are invalid
