package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	g := util.IntGrid2{}
	var start, goal util.Point2

	for y, line := range lines {
		for x, c := range line {
			p := util.Point2{x, y}
			switch c {
			case 'S':
				start = p
				g.Set(p, 1)
			case 'E':
				goal = p
				g.Set(p, 26)
			default:
				g.Set(p, 1+int(c-'a'))
			}
		}
	}

	fmt.Println("part1:", shortestPath(g, start, goal))

	part2 := math.MaxInt
	g.Each(func(p util.Point2, v int) {
		if v == 1 {
			part2 = util.IMin(shortestPath(g, p, goal), part2)
		}
	})
	fmt.Println("part2:", part2)
}

type searchState struct {
	pos   util.Point2
	steps int
}

type searchHeap []searchState

func (h searchHeap) Len() int           { return len(h) }
func (h searchHeap) Less(i, j int) bool { return h[i].steps < h[j].steps }
func (h searchHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *searchHeap) Push(x any) {
	*h = append(*h, x.(searchState))
}

func (h *searchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func shortestPath(g util.IntGrid2, start, goal util.Point2) int {
	search := &searchHeap{}
	heap.Init(search)
	heap.Push(search, searchState{start, 0})

	seen := map[util.Point2]bool{}

	for search.Len() > 0 {
		cur := heap.Pop(search).(searchState)
		if cur.pos == goal {
			return cur.steps
		}

		for _, p := range cur.pos.Touching() {
			if _, ok := seen[p]; ok {
				continue
			}

			h := g.Get(p)
			if h > 0 && h <= g.Get(cur.pos)+1 {
				seen[p] = true
				heap.Push(search, searchState{p, cur.steps + 1})
			}
		}
	}

	return math.MaxInt
}
