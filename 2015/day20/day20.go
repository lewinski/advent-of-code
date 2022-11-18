package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := 34000000

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

// thought i might have a recursive quick form but it is better to just brute force it
// http://eulerarchive.maa.org/docs/translations/E175en.pdf

func part1(input int) int {
	max := input / 10
	houses := make([]int, max)
	for elf := 1; elf < max; elf++ {
		for house := elf; house < max; house += elf {
			houses[house] += elf * 10
		}
	}
	for house, presents := range houses {
		if presents >= input {
			return house
		}
	}
	panic("oh no")
}

func part2(input int) int {
	max := input / 10
	houses := make([]int, max)
	for elf := 1; elf < max; elf++ {
		stop := util.IMin(elf*50, max)
		for house := elf; house < stop; house += elf {
			houses[house] += elf * 11
		}
	}
	for house, presents := range houses {
		if presents >= input {
			return house
		}
	}
	panic("oh no")
}
