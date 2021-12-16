package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input2.txt")

	size := len(lines)
	g := util.IntGrid2{}
	for y, line := range lines {
		for x, r := range line {
			g.SetCoords(x, y, util.MustAtoi(string(r)))
		}
	}

	fmt.Println("part1:", search(g, size))
	fmt.Println("part2:", search(extend(g, size), size*5))
}

func extend(g util.IntGrid2, size int) util.IntGrid2 {
	h := util.IntGrid2{}
	g.Each(func(p util.Point2, x int) {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				risk := x + i + j
				for risk > 9 {
					risk -= 9
				}
				h.SetCoords(p[0]+i*size, p[1]+j*size, risk)
			}
		}
	})
	return h
}

// heap implementation for graph search
type searchState struct {
	pos  util.Point2
	cost int
}

type searchHeap []searchState

func (h searchHeap) Len() int           { return len(h) }
func (h searchHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h searchHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *searchHeap) Push(x interface{}) {
	*h = append(*h, x.(searchState))
}

func (h *searchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func search(g util.IntGrid2, size int) int {
	start := util.Origin2()

	distance := util.IntGrid2{}
	distance[start] = 0

	visited := map[util.Point2]bool{}
	visited[start] = true

	unvisited := &searchHeap{}
	heap.Init(unvisited)
	heap.Push(unvisited, searchState{start, 0})

	for unvisited.Len() > 0 {
		next := heap.Pop(unvisited).(searchState)
		visited[next.pos] = true

		for _, p := range next.pos.Touching() {
			if !g.Contains(p) {
				continue
			}
			if _, ok := visited[p]; ok {
				continue
			}

			oldCost := math.MaxInt
			if d, ok := distance[p]; ok {
				oldCost = d
			}
			newCost := next.cost + g[p]
			if newCost < oldCost {
				distance[p] = newCost
				heap.Push(unvisited, searchState{p, newCost})
			}
		}
	}

	return distance[util.Point2{size - 1, size - 1}]
}
