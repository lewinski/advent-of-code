package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	memory := util.IntFields(lines[0])
	fmt.Println("part1:", part1(memory))

	memory = util.IntFields(lines[0])
	fmt.Println("part2:", part2(memory))
}

func redistribute(memory []int) []int {
	// find largest
	max := memory[0]
	pos := 0
	for i, x := range memory {
		if x > max {
			pos = i
			max = x
		}
	}

	// zero
	memory[pos] = 0

	// redistribute
	for i := 0; i < max; i++ {
		pos++
		if pos == len(memory) {
			pos = 0
		}
		memory[pos]++
	}

	return memory
}

func part1(memory []int) int {
	seen := map[string]bool{}
	steps := 0

	for {
		memory = redistribute(memory)
		steps++

		configuration := fmt.Sprintf("%v", memory)
		if _, ok := seen[configuration]; ok {
			return steps
		}
		seen[configuration] = true
	}
}

func part2(memory []int) int {
	seen := map[string]int{}
	steps := 0

	for {
		memory = redistribute(memory)
		steps++

		configuration := fmt.Sprintf("%v", memory)
		if last, ok := seen[configuration]; ok {
			return steps - last
		}
		seen[configuration] = steps
	}
}
