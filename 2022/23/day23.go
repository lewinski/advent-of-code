package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	grid := util.IntGrid2{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				grid.SetCoords(x, y, 1)
			}
		}
	}

	rules := []struct {
		proposedMove util.Point2
		emptyMoves   []util.Point2
	}{
		{
			// north
			proposedMove: util.Point2{0, -1},
			emptyMoves:   []util.Point2{{-1, -1}, {0, -1}, {1, -1}},
		},
		{
			// south
			proposedMove: util.Point2{0, 1},
			emptyMoves:   []util.Point2{{-1, 1}, {0, 1}, {1, 1}},
		},
		{
			// west
			proposedMove: util.Point2{-1, 0},
			emptyMoves:   []util.Point2{{-1, -1}, {-1, 0}, {-1, 1}},
		},
		{
			// east
			proposedMove: util.Point2{1, 0},
			emptyMoves:   []util.Point2{{1, -1}, {1, 0}, {1, 1}},
		},
	}

	round := 0
	for {
		proposals := map[util.Point2]util.Point2{}
		counts := map[util.Point2]int{}

		grid.Each(func(p util.Point2, v int) {

			// if nobody is around current elf, stay still
			around := 0
			for _, d := range p.Around() {
				if grid.Contains(d) {
					around++
				}
			}
			if around == 0 {
				proposals[p] = p
				counts[p]++
				return
			}

			// see if any movement rules apply
		nextRule:
			for r := 0; r < len(rules); r++ {
				rule := rules[(round+r)%len(rules)]
				for _, d := range rule.emptyMoves {
					if grid.Contains(p.Offset(d)) {
						continue nextRule
					}
				}
				next := p.Offset(rule.proposedMove)
				proposals[p] = next
				counts[next]++
				return
			}

			// no place to move to, so stay still
			proposals[p] = p
			counts[p]++
		})

		// calculate next state
		next := util.IntGrid2{}

		moved := 0
		grid.Each(func(p util.Point2, v int) {
			if counts[proposals[p]] == 1 {
				// new location if we are the only one wanting to occupy it
				next.Set(proposals[p], 1)
				// if we actually moved, count it
				if proposals[p] != p {
					moved++
				}
			} else {
				// otherwise don't move
				next.Set(p, 1)
			}
		})

		// apply next state
		grid = next
		round++

		// print free spaces in round 10
		if round == 10 {
			min, max := grid.Bounds()
			fmt.Println("part1:", (max[0]-min[0]+1)*(max[1]-min[1]+1)-len(grid))
		}

		// if nobody moved, we are done with part 2
		if moved == 0 {
			fmt.Println("part2:", round)
			break
		}
	}
}

func printGrid(g util.IntGrid2) {
	min, max := g.Bounds()
	for y := min[1]; y <= max[1]; y++ {
		for x := min[0]; x <= max[0]; x++ {
			if g.ContainsCoords(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
