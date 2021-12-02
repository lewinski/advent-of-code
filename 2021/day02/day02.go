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

func part1(commands []string) int {
	x, y := 0, 0

	for _, command := range commands {
		f := strings.Fields(command)
		arg := util.MustAtoi(f[1])

		switch f[0] {
		case "forward":
			x += arg
		case "down":
			y += arg
		case "up":
			y -= arg
		}
	}

	return x * y
}

func part2(commands []string) int {
	x, y, aim := 0, 0, 0

	for _, command := range commands {
		f := strings.Fields(command)
		arg := util.MustAtoi(f[1])

		switch f[0] {
		case "forward":
			x += arg
			y += arg * aim
		case "down":
			aim += arg
		case "up":
			aim -= arg
		}
	}

	return x * y
}
