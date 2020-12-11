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

func (f ferry) occupiedAround(i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if i+x >= 0 && i+x < f.h {
				if j+y >= 0 && j+y < f.w {
					if f.at(i+x, j+y) == occupied {
						count++
					}
				}
			}
		}
	}
	return count
}

func (f ferry) occupiedDirectional(i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for mult := 1; ; mult++ {
				ix := i + (mult * x)
				jy := j + (mult * y)
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
	}
	return count
}
