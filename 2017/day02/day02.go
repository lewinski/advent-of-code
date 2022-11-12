package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	ss := [][]int{}
	for _, line := range lines {
		ss = append(ss, util.IntFields(line))
	}

	fmt.Println("part1:", part1(ss))
	fmt.Println("part2:", part2(ss))
}

func part1(ss [][]int) int {
	sum := 0

	for _, y := range ss {
		min, max := y[0], y[0]
		for _, x := range y {
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
		}
		sum += max - min
	}

	return sum
}

func part2(ss [][]int) int {
	sum := 0

line:
	for _, y := range ss {
		for i := range y {
			for j := range y {
				if i == j {
					continue
				}
				if y[i]%y[j] == 0 {
					sum += y[i] / y[j]
					continue line
				}
			}
		}
	}

	return sum
}
