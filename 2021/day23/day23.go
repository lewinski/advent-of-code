package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	fmt.Println("part1:", part1())
	fmt.Println("part2:", part2())
}

type move struct {
	idx int
	dst util.Point2
}

type state struct {
	pods   []util.Point2 // where pods are
	grid   util.IntGrid2 // what things are occupied
	moves  []int         // how many moves each pod has taken
	cost   []int         // how much it costs to move each pod
	goal   []int         // which column the pod wants to be in
	energy int           // energy expended to get to this point
}

func newState(size int) state {
	s := state{}
	s.pods = make([]util.Point2, size)
	s.grid = util.IntGrid2{}
	s.moves = make([]int, size)
	s.cost = make([]int, size)
	s.goal = make([]int, size)
	q := size / 4
	for i := 0; i < q; i++ {
		s.cost[i] = 1
		s.goal[i] = 3
		s.cost[i+q] = 10
		s.goal[i+q] = 5
		s.cost[i+2*q] = 100
		s.goal[i+2*q] = 7
		s.cost[i+3*q] = 1000
		s.goal[i+3*q] = 9
	}
	return s
}

func dupState(s state) state {
	r := newState(len(s.pods))
	r.energy = s.energy
	copy(r.pods, s.pods)
	copy(r.moves, s.moves)
	s.grid.Each(func(p util.Point2, x int) {
		r.grid.Set(p, x)
	})
	return r
}

type searchHeap []state

func (h searchHeap) Len() int           { return len(h) }
func (h searchHeap) Less(i, j int) bool { return h[i].energy < h[j].energy }
func (h searchHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *searchHeap) Push(x interface{}) {
	*h = append(*h, x.(state))
}

func (h *searchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (s state) set(idx int, p util.Point2) {
	s.pods[idx] = p
	s.grid.Set(p, s.goal[idx])
}

func (s state) unset(p util.Point2) {
	s.grid.Set(p, 0)
}

func (s state) blocked(p util.Point2) bool {
	return s.grid.Get(p) != 0
}

func (s state) blockedPath(src, dst util.Point2) bool {
	if src[1] != 1 {
		for y := src[1] - 1; y != 1; y-- {
			if s.blocked(util.Point2{src[0], y}) {
				return true
			}
		}
	}

	sign := 1
	if dst[0] < src[0] {
		sign = -1
	}
	for x := src[0]; x != dst[0]+sign; x += sign {
		if src[0] == x && src[1] == 1 {
			continue
		}
		if s.blocked(util.Point2{x, 1}) {
			return true
		}
	}

	if dst[1] != 1 {
		for y := dst[1]; y != 1; y-- {
			if s.blocked(util.Point2{dst[0], y}) {
				return true
			}
		}
	}

	return false
}

func (s state) solved() bool {
	for i := range s.pods {
		if s.pods[i][0] != s.goal[i] {
			return false
		}
	}
	return true
}

func applyMove(s state, m move) state {
	r := dupState(s)

	for i, p := range s.pods {
		if i == m.idx {
			continue
		}
		if p[0] == m.dst[0] && p[1] == m.dst[1] {
			fmt.Println(s)
			fmt.Println(m)
			panic("invalid move")
		}
	}

	moves := util.IAbs(s.pods[m.idx][0]-m.dst[0]) +
		util.IAbs(s.pods[m.idx][1]-1) +
		util.IAbs(m.dst[1]-1)

	r.unset(s.pods[m.idx])
	r.set(m.idx, m.dst)
	r.energy += moves * s.cost[m.idx]
	r.moves[m.idx]++

	return r
}

func (s state) validMoves(maxY int) []move {
	moves := []move{}

	for i := range s.pods {
		if s.pods[i][0] == s.goal[i] && s.moves[i] > 0 {
			// in room and already moved
			continue
		}

		for y := maxY; y >= 2; y-- {
			// move into deepest part of room
			p := util.Point2{s.goal[i], y}
			if s.grid.Get(p) == s.goal[i] {
				continue
			} else if s.grid.Get(p) != 0 {
				break
			} else if !s.blockedPath(s.pods[i], p) {
				moves = append(moves, move{i, p})
				break
			}
		}

		if s.pods[i][1] > 1 {
			// move to hallways
			hallway := []util.Point2{
				{1, 1},
				{2, 1},
				// {3, 1}, room entrance
				{4, 1},
				// {5, 1}, room entrance
				{6, 1},
				// {7, 1}, room entrance
				{8, 1},
				// {9, 1}, room entrance
				{10, 1},
				{11, 1},
			}
			for _, p := range hallway {
				if !s.blockedPath(s.pods[i], p) {
					moves = append(moves, move{i, p})
				}
			}
		}
	}
	return moves
}

func solve(start state, depth int) int {
	best := math.MaxInt

	search := &searchHeap{}
	heap.Init(search)
	heap.Push(search, start)

	distance := map[string]int{}

	for search.Len() > 0 {
		cur := heap.Pop(search).(state)
		if cur.energy > best {
			continue
		}

		for _, m := range cur.validMoves(depth) {
			s := applyMove(cur, m)
			if s.solved() {
				if s.energy < best {
					best = s.energy
				}
			} else if s.energy > best {
				continue
			} else {
				key := fmt.Sprintf("%v", s.pods)
				oldDistance := math.MaxInt
				if d, ok := distance[key]; ok {
					oldDistance = d
				}
				if s.energy < oldDistance {
					distance[key] = s.energy
					heap.Push(search, s)
				}
			}
		}
	}

	return best
}

func part1() int {
	// #############
	// #...........#
	// ###D#A#D#C###
	//   #B#C#B#A#
	//   #########
	start := newState(8)

	start.set(0, util.Point2{5, 2})
	start.set(1, util.Point2{9, 3})

	start.set(2, util.Point2{3, 3})
	start.set(3, util.Point2{7, 3})

	start.set(4, util.Point2{5, 3})
	start.set(5, util.Point2{9, 2})

	start.set(6, util.Point2{3, 2})
	start.set(7, util.Point2{7, 2})

	return solve(start, 3)

	// start = applyMove(start, move{5, util.Point2{6, 1}})
	// start = applyMove(start, move{1, util.Point2{10, 1}})
	// start = applyMove(start, move{0, util.Point2{2, 1}})
	// start = applyMove(start, move{7, util.Point2{9, 3}})
	// start = applyMove(start, move{3, util.Point2{8, 1}})
	// start = applyMove(start, move{5, util.Point2{7, 3}})
	// start = applyMove(start, move{4, util.Point2{7, 2}})
	// start = applyMove(start, move{3, util.Point2{5, 3}})
	// start = applyMove(start, move{6, util.Point2{9, 2}})
	// start = applyMove(start, move{2, util.Point2{5, 2}})
	// start = applyMove(start, move{0, util.Point2{3, 3}})
	// start = applyMove(start, move{1, util.Point2{3, 2}})
	// fmt.Println(start)
	// fmt.Println(start.solved())
}

func part2() int {
	// #############
	// #...........#
	// ###D#A#D#C###
	//   #D#C#B#A#
	//   #D#B#A#C#
	//   #B#C#B#A#
	//   #########
	start := newState(16)

	start.set(0, util.Point2{5, 2})
	start.set(1, util.Point2{9, 3})
	start.set(2, util.Point2{7, 4})
	start.set(3, util.Point2{9, 5})

	start.set(4, util.Point2{7, 3})
	start.set(5, util.Point2{5, 4})
	start.set(6, util.Point2{3, 5})
	start.set(7, util.Point2{7, 5})

	start.set(8, util.Point2{9, 2})
	start.set(9, util.Point2{5, 3})
	start.set(10, util.Point2{9, 4})
	start.set(11, util.Point2{5, 5})

	start.set(12, util.Point2{3, 2})
	start.set(13, util.Point2{7, 2})
	start.set(14, util.Point2{3, 3})
	start.set(15, util.Point2{3, 4})

	return solve(start, 5)
}
