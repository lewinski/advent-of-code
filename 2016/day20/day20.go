package main

import (
	"fmt"
	"sort"

	"github.com/lewinski/advent-of-code/util"
)

type block struct {
	low, high int
}

func (b block) adjacent(other block) bool {
	return b.high+1 == other.low || b.low-1 == other.high
}

func (b block) overlaps(other block) bool {
	return b.low <= other.high && b.high >= other.low
}

func (b *block) merge(other block) {
	if other.low < b.low {
		b.low = other.low
	}
	if other.high > b.high {
		b.high = other.high
	}
}

func main() {
	blocks := []block{}

	// read all the blocks
	lines := util.Lines("input.txt")
	for i := 0; i < len(lines); i++ {
		var b block
		fmt.Sscanf(lines[i], "%d-%d", &b.low, &b.high)
		blocks = append(blocks, b)
	}

	// sort by low address
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].low < blocks[j].low
	})

	// merge overlapping or adjacent blocks
	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {
			if blocks[i].overlaps(blocks[j]) || blocks[i].adjacent(blocks[j]) {
				blocks[i].merge(blocks[j])
				blocks = append(blocks[:j], blocks[j+1:]...)
				j--
			}
		}
	}

	// first address is after the first block of invalid addresses
	fmt.Println("part1:", blocks[0].high+1)

	// count the number of valid addresses
	valid := 1 << 32
	for _, b := range blocks {
		valid -= b.high - b.low + 1
	}
	fmt.Println("part2:", valid)
}
