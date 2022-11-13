package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(dimensions []string) int {
	total := 0
	for _, d := range dimensions {
		m := strings.SplitN(d, "x", 3)
		w, h, l := util.MustAtoi(m[0]), util.MustAtoi(m[1]), util.MustAtoi(m[2])
		side1 := w * h
		side2 := h * l
		side3 := l * w
		min := util.IMin(side1, util.IMin(side2, side3))
		total += 2*side1 + 2*side2 + 2*side3 + min
	}
	return total
}

func part2(dimensions []string) int {
	total := 0
	for _, d := range dimensions {
		m := strings.SplitN(d, "x", 3)
		w, h, l := util.MustAtoi(m[0]), util.MustAtoi(m[1]), util.MustAtoi(m[2])
		perim1 := 2 * (w + h)
		perim2 := 2 * (h + l)
		perim3 := 2 * (l + w)
		min := util.IMin(perim1, util.IMin(perim2, perim3))
		total += min + w*h*l
	}
	return total
}
