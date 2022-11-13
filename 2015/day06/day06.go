package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(commands []string) int {
	g := util.IntGrid2{}

	for _, command := range commands {
		command = strings.Replace(command, "turn ", "turn", 1)

		var op string
		var x1, y1, x2, y2 int
		fmt.Sscanf(command, "%s %d,%d through %d,%d", &op, &x1, &y1, &x2, &y2)

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch op {
				case "toggle":
					g.SetCoords(x, y, 1-g.GetCoords(x, y))
				case "turnon":
					g.SetCoords(x, y, 1)
				case "turnoff":
					g.SetCoords(x, y, 0)
				}
			}
		}
	}

	var count int
	g.Each(func(p util.Point2, v int) {
		if v == 1 {
			count++
		}
	})
	return count
}

func part2(commands []string) int {
	g := util.IntGrid2{}

	for _, command := range commands {
		command = strings.Replace(command, "turn ", "turn", 1)

		var op string
		var x1, y1, x2, y2 int
		fmt.Sscanf(command, "%s %d,%d through %d,%d", &op, &x1, &y1, &x2, &y2)

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch op {
				case "toggle":
					g.SetCoords(x, y, g.GetCoords(x, y)+2)
				case "turnon":
					g.SetCoords(x, y, g.GetCoords(x, y)+1)
				case "turnoff":
					if g.GetCoords(x, y) > 0 {
						g.SetCoords(x, y, g.GetCoords(x, y)-1)

					}
				}
			}
		}
	}

	var brightness int
	g.Each(func(p util.Point2, v int) {
		brightness += v
	})

	return brightness
}
