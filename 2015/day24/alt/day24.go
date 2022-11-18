package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/lewinski/advent-of-code/util"
)

// might eventually terminate but takes too long
func main() {
	packages := util.IntLines("../input.txt")
	sort.Slice(packages, func(i, j int) bool {
		return packages[i] > packages[j]
	})
	fmt.Println("part1:", solve(packages, 3))
	fmt.Println("part2:", solve(packages, 4))
}

func solve(packages []int, groups int) int {
	total := 0
	for _, p := range packages {
		total += p
	}

	best := math.MaxInt
	bestSize := math.MaxInt

	weights := make([]int, groups)
	counts := make([]int, groups)
	qe := make([]int, groups)

	assignment := newAssignment(len(packages), groups)
next:
	for !assignment.done() {
		for i := 0; i < groups; i++ {
			weights[i] = 0
			counts[i] = 0
			qe[i] = 1
		}

		for i, a := range assignment.assignments {
			weights[a] += packages[i]
			counts[a]++
			qe[a] *= packages[i]
		}
		assignment.inc()

		for i := 1; i < groups; i++ {
			if weights[i] != weights[0] {
				continue next
			}
		}

		for i := 1; i < groups; i++ {
			if counts[i] < bestSize {
				best = qe[i]
				bestSize = counts[i]
			} else if counts[i] == bestSize && qe[i] < best {
				best = qe[i]
			}
		}
	}

	return best
}

type assignment struct {
	assignments []int
	max         int
}

func newAssignment(items, groups int) *assignment {
	return &assignment{
		assignments: make([]int, items),
		max:         groups,
	}
}

func (a *assignment) inc() {
	for i := range a.assignments {
		a.assignments[i]++
		if a.assignments[i] < a.max {
			return
		}
		if i != len(a.assignments)-1 {
			a.assignments[i] = 0
		}
	}
}

func (a *assignment) done() bool {
	return a.assignments[len(a.assignments)-1] == a.max
}
