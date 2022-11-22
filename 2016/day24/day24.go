package main

import (
	"container/list"
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	grid, pois := parseGrid("input.txt")

	// pairwise distances between every point
	dists := map[[2]int]int{}
	for i, p1 := range pois {
		d := bfsDistances(p1, grid)
		for j, p2 := range pois {
			if i == j {
				continue
			}
			dists[[2]int{i, j}] = d[p2]
		}
	}

	fmt.Println("part1:", travel('0', 0, dists, map[int]bool{'0': true}))
	fmt.Println("part2:", travelLoop('0', '0', 0, dists, map[int]bool{'0': true}))

}

func parseGrid(filename string) (util.IntGrid2, map[int]util.Point2) {
	input := util.Lines(filename)

	g := util.IntGrid2{}
	pois := map[int]util.Point2{}

	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				g.SetCoords(x, y, -1)
			} else if c == '.' {
				g.SetCoords(x, y, 0)
			} else {
				pois[int(c)] = util.Point2{x, y}
				g.SetCoords(x, y, int(c))
			}
		}
	}

	return g, pois
}

func bfsDistances(pos util.Point2, grid util.IntGrid2) map[util.Point2]int {
	candidates := list.New()
	reachable := map[util.Point2]int{}

	reachable[pos] = 0
	candidates.PushBack(pos)

	for candidates.Len() > 0 {
		p := candidates.Front().Value.(util.Point2)
		candidates.Remove(candidates.Front())
		for _, o := range p.Touching() {
			// already been here
			if _, ok := reachable[o]; ok {
				continue
			}
			// can't go here
			if grid.Get(o) == -1 {
				continue
			}
			// we can reach here with an additional step from where we are now
			reachable[o] = reachable[p] + 1
			candidates.PushBack(o)
		}
	}

	return reachable
}

// best path for visiting every point once
func travel(pos, traveled int, dists map[[2]int]int, seen map[int]bool) int {
	best := math.MaxInt
	done := true
	for k, d := range dists {
		if k[0] != pos {
			continue
		}
		if seen[k[1]] {
			continue
		}
		done = false
		seen[k[1]] = true
		best = util.IMin(best, travel(k[1], traveled+d, dists, seen))
		seen[k[1]] = false
	}

	if done {
		return traveled
	}
	return best
}

// best path for visiting every point once and returning to initial
func travelLoop(initial, current, traveled int, dists map[[2]int]int, seen map[int]bool) int {
	best := math.MaxInt
	done := true
	for k, d := range dists {
		if k[0] != current {
			continue
		}
		if seen[k[1]] {
			continue
		}
		done = false
		seen[k[1]] = true
		best = util.IMin(best, travelLoop(initial, k[1], traveled+d, dists, seen))
		seen[k[1]] = false
	}

	if done {
		return traveled + dists[[2]int{current, initial}]
	}
	return best
}
