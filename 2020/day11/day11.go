package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	ferry := parseInput(util.Lines("input.txt"))
	for {
		gerry := ferry.iter1()
		if ferry.String() == gerry.String() {
			break
		}
		ferry = gerry
	}
	fmt.Println("part1:", ferry.occupiedSeats())

	ferry = parseInput(util.Lines("input.txt"))
	for {
		gerry := ferry.iter2()
		if ferry.String() == gerry.String() {
			break
		}
		ferry = gerry
	}
	fmt.Println("part2:", ferry.occupiedSeats())
}

const (
	floor    int = 0
	empty    int = 1
	occupied int = 2
	unknown  int = 3
)

type ferry struct {
	util.IntGrid
}

func newFerry(h, w int) ferry {
	ferry := ferry{}
	ferry.IntGrid = util.NewIntGrid(h, w)
	return ferry
}

func parseInput(lines []string) ferry {
	ferry := newFerry(len(lines), len(lines[0]))
	for i, line := range lines {
		for j, x := range line {
			switch x {
			case '.':
				ferry.Set(i, j, floor)
			case 'L':
				ferry.Set(i, j, empty)
			case '#':
				ferry.Set(i, j, occupied)
			default:
				ferry.Set(i, j, unknown)
			}
		}
	}
	return ferry
}

func (f ferry) String() string {
	var sb strings.Builder
	for i := 0; i < f.Height(); i++ {
		for j := 0; j < f.Width(); j++ {
			switch f.Get(i, j) {
			case floor:
				sb.WriteString(".")
			case empty:
				sb.WriteString("L")
			case occupied:
				sb.WriteString("#")
			default:
				sb.WriteString("?")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (f ferry) iter1() ferry {
	n := newFerry(f.Height(), f.Width())
	for i := 0; i < f.Height(); i++ {
		for j := 0; j < f.Width(); j++ {
			o := f.occupiedAround(i, j)
			if f.Get(i, j) == empty && o == 0 {
				n.Set(i, j, occupied)
			} else if f.Get(i, j) == occupied && o >= 4 {
				n.Set(i, j, empty)
			} else {
				n.Set(i, j, f.Get(i, j))
			}
		}
	}
	return n
}

func (f ferry) iter2() ferry {
	n := newFerry(f.Height(), f.Width())
	for i := 0; i < f.Height(); i++ {
		for j := 0; j < f.Width(); j++ {
			o := f.occupiedDirectional(i, j)
			if f.Get(i, j) == empty && o == 0 {
				n.Set(i, j, occupied)
			} else if f.Get(i, j) == occupied && o >= 5 {
				n.Set(i, j, empty)
			} else {
				n.Set(i, j, f.Get(i, j))
			}
		}
	}
	return n
}

func (f ferry) occupiedSeats() (count int) {
	for i := 0; i < f.Height(); i++ {
		for j := 0; j < f.Width(); j++ {
			if f.Get(i, j) == occupied {
				count++
			}
		}
	}
	return
}

func directions() [8][2]int {
	return [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
}

func (f ferry) occupiedAround(i, j int) int {
	count := 0
	for _, offset := range directions() {
		ix := i + offset[0]
		jy := j + offset[1]
		if !f.Contains(ix, jy) {
			continue
		}
		if f.Get(ix, jy) == occupied {
			count++
		}
	}
	return count
}

func (f ferry) occupiedDirectional(i, j int) int {
	count := 0
	for _, offset := range directions() {
		for mult := 1; ; mult++ {
			ix := i + (mult * offset[0])
			jy := j + (mult * offset[1])
			if !f.Contains(ix, jy) {
				break
			}
			if f.Get(ix, jy) == floor {
				continue
			}
			if f.Get(ix, jy) == occupied {
				count++
			}
			break
		}
	}
	return count
}
