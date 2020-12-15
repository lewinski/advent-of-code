package main

import (
	"fmt"
)

func main() {
	input := []int{0, 13, 1, 16, 6, 17}
	fmt.Println("part1:", elfGame(input, 2020))
	fmt.Println("part2:", elfGame(input, 30000000))
}

func elfGame(seed []int, turns int) int {
	before1 := map[int]int{}
	before2 := map[int]int{}
	prev := 0

	for turn := 0; turn < turns; turn++ {
		var next int
		if turn < len(seed) {
			next = seed[turn]
		} else {
			x, _ := before1[prev]
			y, in2 := before2[prev]
			if !in2 {
				next = 0
			} else {
				next = x - y
			}
		}

		if x, ok := before1[next]; ok {
			before2[next] = x
		}
		before1[next] = turn
		prev = next
	}

	return prev
}
