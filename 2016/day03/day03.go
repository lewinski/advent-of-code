package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	valid := 0

	for _, line := range lines {
		sides := strings.Fields(line)

		a := util.MustAtoi(sides[0])
		b := util.MustAtoi(sides[1])
		c := util.MustAtoi(sides[2])

		if a+b > c && a+c > b && b+c > a {
			valid++
		}
	}
	fmt.Println("part1:", valid)

	valid = 0
	for i := 0; i < len(lines); i += 3 {
		a := strings.Fields(lines[i])
		b := strings.Fields(lines[i+1])
		c := strings.Fields(lines[i+2])

		for j := 0; j < 3; j++ {
			aj := util.MustAtoi(a[j])
			bj := util.MustAtoi(b[j])
			cj := util.MustAtoi(c[j])

			if aj+bj > cj && aj+cj > bj && bj+cj > aj {
				valid++
			}
		}
	}

	fmt.Println("part1:", valid)
}
