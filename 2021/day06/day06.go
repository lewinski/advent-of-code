package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fish := [9]int{}
	for _, days := range strings.Split(lines[0], ",") {
		fish[util.MustAtoi(days)]++
	}

	fmt.Println("part1:", part1(fish))
	fmt.Println("part2:", part2(fish))
}

func tick(fish [9]int) [9]int {
	next := [9]int{}
	for days, count := range fish {
		if days == 0 {
			next[6] += count
			next[8] += count
		} else {
			next[days-1] += count
		}
	}
	return next
}

func part1(fish [9]int) int {
	for i := 0; i < 80; i++ {
		fish = tick(fish)
	}

	sum := 0
	for _, count := range fish {
		sum += count
	}
	return sum
}

func part2(fish [9]int) int {
	for i := 0; i < 256; i++ {
		fish = tick(fish)
	}

	sum := 0
	for _, count := range fish {
		sum += count
	}
	return sum
}
