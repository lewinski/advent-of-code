package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

type search struct {
	pos   util.Point2
	depth int
}

func main() {
	input := 1364

	fmt.Println("part1:", part1(input, util.Point2{31, 39}))
	fmt.Println("part2:", part2(input, 50))

}

func part1(input int, goal util.Point2) int {
	reachable := util.IntGrid2{}
	reachable.SetCoords(1, 1, 0)

	queue := []search{{util.Point2{1, 1}, 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		d := s.depth + 1

		for _, o := range s.pos.Touching() {
			if o[0] < 0 || o[1] < 0 {
				continue
			}

			if !reachable.Contains(o) {
				if wall(o[0], o[1], input) {
					reachable.Set(o, -1)
				} else {
					reachable.Set(o, d)
					queue = append(queue, search{o, d})
				}
			} else {
				c := reachable.Get(o)
				if c < 0 {
					continue
				}
				if d < c {
					reachable.Set(o, d)
					queue = append(queue, search{o, d})
				}
			}
		}

		if reachable.Contains(goal) {
			return reachable.Get(goal)
		}
	}
	panic("oh no")
}

func part2(input int, maxDepth int) int {
	reachable := util.IntGrid2{}
	reachable.SetCoords(1, 1, 0)

	queue := []search{{util.Point2{1, 1}, 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		d := s.depth + 1
		if d > maxDepth {
			continue
		}

		for _, o := range s.pos.Touching() {
			if o[0] < 0 || o[1] < 0 {
				continue
			}

			if !reachable.Contains(o) {
				if wall(o[0], o[1], input) {
					reachable.Set(o, -1)
				} else {
					reachable.Set(o, d)
					queue = append(queue, search{o, d})
				}
			} else {
				c := reachable.Get(o)
				if c < 0 {
					continue
				}
				if d < c {
					reachable.Set(o, d)
					queue = append(queue, search{o, d})
				}
			}
		}
	}

	count := 0
	reachable.Each(func(p util.Point2, v int) {
		if v >= 0 {
			count++
		}
	})
	return count
}

func wall(x, y, input int) bool {
	n := x*x + 3*x + 2*x*y + y + y*y + input
	c := 0
	for n > 0 {
		c += n & 1
		n >>= 1
	}
	return c%2 == 1
}
