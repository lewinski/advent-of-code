package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	width := 40
	height := 6
	size := width * height

	pc := 0
	cycles := 0
	x := 1
	history := make([]int, size)

	for cycles < len(history) && pc < len(lines) {
		insn := strings.Fields(lines[pc])
		switch insn[0] {
		case "noop":
			history[cycles] = x
			cycles += 1
		case "addx":
			history[cycles] = x
			history[cycles+1] = x
			cycles += 2
			x += util.MustAtoi(insn[1])
		}
		pc++
	}

	sum := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		sum += i * history[i-1]
	}
	fmt.Println("part1:", sum)

	fmt.Println("part2:")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x-1 <= history[y*width+x] && history[y*width+x] <= x+1 {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
