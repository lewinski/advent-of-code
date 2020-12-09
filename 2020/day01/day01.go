package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	vals := make(map[int]bool)

	for _, val := range util.IntLines("input.txt") {
		if _, ok := vals[val]; ok {
			panic("duplicate value")
		}
		vals[val] = true
	}

	goal := 2020

part1:
	for k := range vals {
		if _, ok := vals[goal-k]; ok {
			fmt.Println("part1:", k*(goal-k))
			break part1
		}
	}

part2:
	for k := range vals {
		for j := range vals {
			if k == j {
				continue
			}
			if _, ok := vals[goal-k-j]; ok {
				fmt.Println("part2:", k*j*(goal-k-j))
				break part2
			}
		}
	}
}
