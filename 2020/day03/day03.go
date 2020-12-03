package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	forest := util.Lines("input.txt")
	fmt.Println("part1:", trees(forest, 3, 1))
	fmt.Println("part2:", trees(forest, 1, 1)*trees(forest, 3, 1)*trees(forest, 5, 1)*trees(forest, 7, 1)*trees(forest, 1, 2))
}

func trees(forest []string, slopeRight, slopeDown int) (trees int) {
	pos := 0
	for i := 0; i < len(forest); i += slopeDown {
		line := forest[i]
		if line[pos%len(line)] == '#' {
			trees++
		}
		pos += slopeRight
	}
	return
}
