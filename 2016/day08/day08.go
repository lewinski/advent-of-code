package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	s := newScreen(50, 6)

	lines := util.Lines("input.txt")
	for _, line := range lines {
		if line[0:4] == "rect" {
			var a, b int
			fmt.Sscanf(line, "rect %dx%d", &a, &b)
			s.rect(a, b)
		} else if line[0:10] == "rotate row" {
			var y, n int
			fmt.Sscanf(line, "rotate row y=%d by %d", &y, &n)
			s.rotateRow(y, n)
		} else if line[0:13] == "rotate column" {
			var x, n int
			fmt.Sscanf(line, "rotate column x=%d by %d", &x, &n)
			s.rotateColumn(x, n)
		}
	}

	on := 0
	s.pixels.Each(func(p util.Point2, v int) {
		on += v
	})

	fmt.Println("part1:", on)
	fmt.Println("part2:")
	fmt.Println(s)
}

type screen struct {
	width, height int
	pixels        util.IntGrid2
}

func newScreen(w, h int) screen {
	return screen{
		width:  w,
		height: h,
		pixels: util.IntGrid2{},
	}
}

func (s screen) String() string {
	output := ""
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			if s.pixels.GetCoords(x, y) == 1 {
				output += "█"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func (s *screen) rect(a, b int) {
	for x := 0; x < a; x++ {
		for y := 0; y < b; y++ {
			s.pixels.SetCoords(x, y, 1)
		}
	}
}

func (s *screen) rotateColumn(x, n int) {
	col := make([]int, s.height)
	for y := 0; y < s.height; y++ {
		col[y] = s.pixels.GetCoords(x, y)
	}
	for y := 0; y < s.height; y++ {
		s.pixels.SetCoords(x, y, col[(y-n+s.height)%s.height])
	}
}

func (s *screen) rotateRow(y, n int) {
	row := make([]int, s.width)
	for x := 0; x < s.width; x++ {
		row[x] = s.pixels.GetCoords(x, y)
	}
	for x := 0; x < s.width; x++ {
		s.pixels.SetCoords(x, y, row[(x-n+s.width)%s.width])
	}
}
