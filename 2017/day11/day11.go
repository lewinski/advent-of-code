package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	q, r, s := 0, 0, 0
	maxDist := 0

	for _, dir := range strings.Split(lines[0], ",") {
		switch dir {
		case "n":
			s += 1
			r -= 1
		case "ne":
			q += 1
			r -= 1
		case "nw":
			q -= 1
			s += 1
		case "s":
			s -= 1
			r += 1
		case "se":
			s -= 1
			q += 1
		case "sw":
			q -= 1
			r += 1
		}
		d := dist(q, r, s)
		if d > maxDist {
			maxDist = d
		}
	}

	fmt.Println("part1:", dist(q, r, s))
	fmt.Println("part2:", maxDist)
}

func dist(q, r, s int) int {
	return util.IMax(util.IAbs(q), util.IMax(util.IAbs(r), util.IAbs(s)))
}
