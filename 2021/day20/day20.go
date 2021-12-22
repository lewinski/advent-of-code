package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	recs := util.Records("input.txt")

	algorithm := recs[0]
	state := util.IntGrid2{}

	for y, line := range strings.Split(recs[1], "\n") {
		for x, c := range line {
			if c == '#' {
				state.SetCoords(x, y, 1)
			} else {
				state.SetCoords(x, y, 0)
			}
		}
	}

	fmt.Println("part1:", enhance(algorithm, state, 2))
	fmt.Println("part2:", enhance(algorithm, state, 50))
}

func lookup(algorithm string, idx int) int {
	if algorithm[idx] == '#' {
		return 1
	}
	return 0
}

func step(algorithm string, state util.IntGrid2, min, max, outside int) util.IntGrid2 {
	next := util.IntGrid2{}

	for x := min; x <= max; x++ {
		for y := min; y <= max; y++ {
			p := util.Point2{x, y}

			idx := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					dp := p.OffsetCoords(dx, dy)
					if state.Contains(dp) {
						idx = idx*2 + state.Get(dp)
					} else {
						idx = idx*2 + outside
					}
				}
			}

			next.SetCoords(x, y, lookup(algorithm, idx))
		}
	}

	return next
}

func enhance(algorithm string, state util.IntGrid2, times int) int {
	min, max := 0, 0
	for p := range state {
		max = util.IMax(max, util.IMax(p[0], p[1]))
	}

	outside := 0
	for i := 0; i < times; i++ {
		min--
		max++
		state = step(algorithm, state, min, max, outside)
		outside = lookup(algorithm, outside*0b111111111)
	}

	cnt := 0
	for _, c := range state {
		cnt += c
	}
	return cnt
}
