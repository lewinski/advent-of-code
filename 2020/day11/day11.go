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
	util.IntGrid2
	h, w int
}

func newFerry(h, w int) ferry {
	return ferry{
		IntGrid2: util.IntGrid2{},
		h:        h,
		w:        w,
	}
}

func parseInput(lines []string) ferry {
	ferry := newFerry(len(lines), len(lines[0]))
	for i, line := range lines {
		for j, x := range line {
			switch x {
			case '.':
				ferry.SetCoords(i, j, floor)
			case 'L':
				ferry.SetCoords(i, j, empty)
			case '#':
				ferry.SetCoords(i, j, occupied)
			default:
				ferry.SetCoords(i, j, unknown)
			}
		}
	}
	return ferry
}

func (f ferry) String() string {
	var sb strings.Builder
	for i := 0; i < f.h; i++ {
		for j := 0; j < f.w; j++ {
			switch f.GetCoords(i, j) {
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
	f.Each(func(p util.Point2, x int) {
		o := f.occupiedAround(p)
		if x == empty && o == 0 {
			n.Set(p, occupied)
		} else if x == occupied && o >= 4 {
			n.Set(p, empty)
		} else {
			n.Set(p, x)
		}
	})
	return n
}

func (f ferry) iter2() ferry {
	n := newFerry(f.h, f.w)
	f.Each(func(p util.Point2, x int) {
		o := f.occupiedDirectional(p)
		if x == empty && o == 0 {
			n.Set(p, occupied)
		} else if x == occupied && o >= 5 {
			n.Set(p, empty)
		} else {
			n.Set(p, x)
		}
	})
	return n
}

func (f ferry) occupiedSeats() (count int) {
	f.Each(func(p util.Point2, x int) {
		if x == occupied {
			count++
		}
	})
	return
}

func (f ferry) occupiedAround(p util.Point2) (count int) {
	for _, a := range p.Around() {
		if f.Get(a) == occupied {
			count++
		}
	}
	return
}

func (f ferry) occupiedDirectional(p util.Point2) (count int) {
	directions := util.Origin2().Around()
	for _, offset := range directions {
		for mult := 1; ; mult++ {
			a := p.Offset(offset.Scale(mult))
			if !f.Contains(a) {
				break
			}
			if f.Get(a) == floor {
				continue
			}
			if f.Get(a) == occupied {
				count++
			}
			break
		}
	}
	return
}
