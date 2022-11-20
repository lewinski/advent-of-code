package main

import (
	"container/heap"
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	s := state{
		objects: map[string]int{
			"ee": 1,
		},
		depth: 0,
	}

	//	The first floor contains a promethium generator and a promethium-compatible microchip.
	s.objects["mg"] = 1
	s.objects["mm"] = 1
	//	The second floor contains a cobalt generator, a curium generator, a ruthenium generator, and a plutonium generator.
	s.objects["cg"] = 2
	s.objects["ug"] = 2
	s.objects["rg"] = 2
	s.objects["pg"] = 2
	//	The third floor contains a cobalt-compatible microchip, a curium-compatible microchip, a ruthenium-compatible microchip, and a plutonium-compatible microchip.
	s.objects["cm"] = 3
	s.objects["um"] = 3
	s.objects["rm"] = 3
	s.objects["pm"] = 3
	//	The fourth floor contains nothing relevant.

	fmt.Println("part1:", solve(s))

	// Extra parts:
	// An elerium generator.
	// An elerium-compatible microchip.
	// A dilithium generator.
	// A dilithium-compatible microchip.
	s.objects["eg"] = 1
	s.objects["em"] = 1
	s.objects["dg"] = 1
	s.objects["dm"] = 1
	fmt.Println("part2:", solve(s))
}

func solve(s state) int {
	seen := map[string]int{}
	seen[s.key()] = s.depth

	sh := &stateHeap{s}
	heap.Init(sh)

	for sh.Len() > 0 {
		s := heap.Pop(sh).(state)
		if s.done() {
			return s.depth
		}

		for _, next := range s.next() {
			key := next.key()
			d, found := seen[key]
			if !found || next.depth < d {
				seen[key] = s.depth
				heap.Push(sh, next)
			}
		}
	}

	panic("oh no")
}

type state struct {
	objects map[string]int
	depth   int
}

func (s state) key() string {
	chips := [4]int{}
	gens := [4]int{}
	for obj, floor := range s.objects {
		if obj[1] == 'm' {
			chips[floor-1]++
		} else {
			gens[floor-1]++
		}
	}
	return fmt.Sprintf("%d-%v-%v", s.objects["ee"], chips, gens)
}

type stateHeap []state

func (h stateHeap) Len() int            { return len(h) }
func (h stateHeap) Less(i, j int) bool  { return h[i].depth < h[j].depth }
func (h stateHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *stateHeap) Push(x interface{}) { *h = append(*h, x.(state)) }
func (h *stateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (s state) dup() state {
	dup := state{
		objects: map[string]int{},
		depth:   s.depth + 1,
	}
	for obj, floor := range s.objects {
		dup.objects[obj] = floor
	}
	return dup
}

func (s state) valid() bool {
	for obj, floor := range s.objects {
		// microchip
		if obj[1] == 'm' {
			// safe: generator is on floor
			if s.objects[obj[0:1]+"g"] == floor {
				continue
			}
			// unsafe: other generator is on floor
			for other, otherFloor := range s.objects {
				if floor == otherFloor && other[1] == 'g' {
					return false
				}
			}

		}
	}
	return true
}

func (s state) done() bool {
	for _, floor := range s.objects {
		if floor != 4 {
			// something is not on the fourth floor
			return false
		}
	}
	return true
}

func (s state) next() []state {
	var next []state

	// don't move anything below where the lowest item currently is
	bottom := 4

	elevator := s.objects["ee"]
	moving := []string{}

	for obj, floor := range s.objects {
		if obj != "ee" && floor == elevator {
			moving = append(moving, obj)
		}
		bottom = util.IMin(bottom, floor)
	}

	for i := 0; i < len(moving); i++ {
		// move elevator and object i up/down one
		if elevator > bottom {
			dup := s.dup()
			dup.objects["ee"] = elevator - 1
			dup.objects[moving[i]] = elevator - 1
			if dup.valid() {
				next = append(next, dup)
			}
		}
		if elevator < 4 {
			dup := s.dup()
			dup.objects["ee"] = elevator + 1
			dup.objects[moving[i]] = elevator + 1
			if dup.valid() {
				next = append(next, dup)
			}
		}

		for j := i + 1; j < len(moving); j++ {
			// move elevator and object i and j up/down one
			if elevator > bottom {
				dup := s.dup()
				dup.objects["ee"] = elevator - 1
				dup.objects[moving[i]] = elevator - 1
				dup.objects[moving[j]] = elevator - 1
				if dup.valid() {
					next = append(next, dup)
				}
			}
			if elevator < 4 {
				dup := s.dup()
				dup.objects["ee"] = elevator + 1
				dup.objects[moving[i]] = elevator + 1
				dup.objects[moving[j]] = elevator + 1
				if dup.valid() {
					next = append(next, dup)
				}
			}
		}
	}

	return next
}
