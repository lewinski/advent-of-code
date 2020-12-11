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
	layout [][]int
	h, w   int
}

func newFerry(h, w int) ferry {
	f := ferry{h: h, w: w}
	f.layout = make([][]int, h)
	for i := 0; i < h; i++ {
		f.layout[i] = make([]int, w)
	}
	return f
}

func (f ferry) at(x, y int) int {
	return f.layout[x][y]
}

func (f *ferry) set(x, y, val int) {
	f.layout[x][y] = val
}

func parseInput(lines []string) ferry {
	ferry := newFerry(len(lines), len(lines[0]))
	for i, line := range lines {
		for j, x := range line {
			switch x {
			case '.':
				ferry.set(i, j, floor)
			case 'L':
				ferry.set(i, j, empty)
			case '#':
				ferry.set(i, j, occupied)
			default:
				ferry.set(i, j, unknown)
			}
		}
	}
	return ferry
}

func (f ferry) String() string {
	var sb strings.Builder
	for i := 0; i < f.h; i++ {
		for j := 0; j < f.w; j++ {
			switch f.at(i, j) {
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
	n := newFerry(f.h, f.w)
	for i := 0; i < f.h; i++ {
		for j := 0; j < f.w; j++ {
			o := f.occupiedAround(i, j)
			if f.at(i, j) == empty && o == 0 {
				n.set(i, j, occupied)
			} else if f.at(i, j) == occupied && o >= 4 {
				n.set(i, j, empty)
			} else {
				n.set(i, j, f.at(i, j))
			}
		}
	}
	return n
}

func (f ferry) iter2() ferry {
	n := newFerry(f.h, f.w)
	for i := 0; i < f.h; i++ {
		for j := 0; j < f.w; j++ {
			o := f.occupiedDirectional(i, j)
			if f.at(i, j) == empty && o == 0 {
				n.set(i, j, occupied)
			} else if f.at(i, j) == occupied && o >= 5 {
				n.set(i, j, empty)
			} else {
				n.set(i, j, f.at(i, j))
			}
		}
	}
	return n
}

func (f ferry) occupiedSeats() (count int) {
	for i := 0; i < f.h; i++ {
		for j := 0; j < f.w; j++ {
			if f.at(i, j) == occupied {
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
			if ix < 0 || ix >= f.h {
				continue
			}
			if jy < 0 || jy >= f.w {
				continue
			}
			if f.at(ix, jy) == occupied {
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
				if ix < 0 || ix >= f.h {
					break
				}
				if jy < 0 || jy >= f.w {
					break
				}
				if f.at(ix, jy) == floor {
					continue
				}
				if f.at(ix, jy) == occupied {
					count++
				}
				break
			}
		}
	return count
}
