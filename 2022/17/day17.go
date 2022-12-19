package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	jets := lines[0]
	curJet := 0

	playArea := util.IntGrid2{}
	highestY := 0
	heights := map[int]int{}

	cycleStart := 0
	cycleEnd := 0

	for i := 0; i < 50000; i++ {
		r := newFallingRock(i, highestY)

		done := false
		for !done {
			r.move(jets[curJet], playArea)
			curJet++
			if curJet == len(jets) {
				if cycleStart == 0 {
					cycleStart = i
				} else if (i%5) == (cycleStart%5) && cycleEnd == 0 {
					cycleEnd = i
				}
				curJet = 0
			}
			done = r.drop(playArea)
		}

		highestY = util.IMax(r.draw(playArea), highestY)
		heights[i] = highestY
		if i == 2021 {
			fmt.Println("part1:", highestY)
		}
	}

	cycleLength := cycleEnd - cycleStart
	cycleHeight := heights[cycleEnd-1] - heights[cycleStart]

	goal := 1000000000000
	extra := (goal - cycleStart) % cycleLength
	numCycles := (goal - cycleStart - extra) / cycleLength

	total := heights[cycleStart-1+extra] + numCycles*cycleHeight
	fmt.Println("part2:", total)
}

type fallingRock struct {
	pos   util.Point2
	shape []util.Point2
}

func newFallingRock(n, y int) fallingRock {
	f := fallingRock{pos: util.Point2{2, y + 4}}

	switch n % 5 {
	case 0:
		f.shape = []util.Point2{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		}
	case 1:
		f.shape = []util.Point2{
			{1, 0},
			{0, 1},
			{1, 1},
			{2, 1},
			{1, 2},
		}
	case 2:
		f.shape = []util.Point2{
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 1},
			{2, 2},
		}
	case 3:
		f.shape = []util.Point2{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		}
	case 4:
		f.shape = []util.Point2{
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		}
	}

	return f
}

func (r *fallingRock) move(dir byte, playArea util.IntGrid2) {
	pos := r.pos
	switch dir {
	case '<':
		pos[0]--
	case '>':
		pos[0]++
	default:
		panic("invalid dir: " + string(dir))
	}

	for _, s := range r.shape {
		p := s.Offset(pos)
		if p[0] < 0 || p[0] > 6 || playArea.Contains(p) {
			return
		}
	}

	r.pos = pos
}

func (r *fallingRock) drop(playArea util.IntGrid2) bool {
	pos := r.pos
	pos[1]--

	for _, s := range r.shape {
		if pos[1] == 0 || playArea.Contains(s.Offset(pos)) {
			return true
		}
	}

	r.pos = pos
	return false
}

func (r *fallingRock) draw(playArea util.IntGrid2) int {
	maxY := 0
	for _, s := range r.shape {
		p := s.Offset(r.pos)
		maxY = util.IMax(maxY, p[1])
		playArea.Set(p, 1)
	}
	return maxY
}
