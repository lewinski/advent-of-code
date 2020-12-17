package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	steps := 6

	cube := parseInput1(input, steps)
	for i := 0; i < steps; i++ {
		cube = cube.step()
	}
	count := 0
	for _, x := range cube.cells {
		count += x
	}
	fmt.Println("part1:", count)

	hypercube := parseInput2(input, steps)
	for i := 0; i < steps; i++ {
		hypercube = hypercube.step()
	}
	count = 0
	for _, x := range hypercube.cells {
		count += x
	}
	fmt.Println("part2:", count)

}

func parseInput1(lines []string, steps int) cube {
	size := util.IMax(len(lines), len(lines[0]))
	half := size / 2
	cube := newCube(size + steps)

	for x, line := range lines {
		for y, c := range line {
			val := 0
			if c == '#' {
				val = 1
			}
			cube.set(x-half, y-half, 0, val)
		}
	}

	return cube
}

func parseInput2(lines []string, steps int) hypercube {
	size := util.IMax(len(lines), len(lines[0]))
	half := size / 2
	cube := newHypercube(size + steps)

	for x, line := range lines {
		for y, c := range line {
			val := 0
			if c == '#' {
				val = 1
			}
			cube.set(0, x-half, y-half, 0, val)
		}
	}

	return cube
}

type point3 struct {
	x, y, z int
}

type cube struct {
	max   int
	size  int
	cells []int
}

func newCube(max int) cube {
	c := cube{max: max, size: max*2 + 1}
	c.cells = make([]int, c.size*c.size*c.size)
	return c
}

func (c cube) at(x, y, z int) int {
	if util.IAbs(x) > c.max || util.IAbs(y) > c.max || util.IAbs(z) > c.max {
		return 0
	}
	offset := ((x + c.max) * c.size * c.size) + ((y + c.max) * c.size) + (z + c.max)
	return c.cells[offset]
}

func (c *cube) set(x, y, z, val int) {
	offset := ((x + c.max) * c.size * c.size) + ((y + c.max) * c.size) + (z + c.max)
	c.cells[offset] = val
}

func (c cube) step() cube {
	d := newCube(c.max)
	for x := -c.max; x < c.max; x++ {
		for y := -c.max; y < c.max; y++ {
			for z := -c.max; z < c.max; z++ {
				neighbors := 0
				for _, p := range around3(x, y, z) {
					if c.at(p.x, p.y, p.z) == 1 {
						neighbors++
					}
				}
				if c.at(x, y, z) == 1 {
					if neighbors == 2 || neighbors == 3 {
						d.set(x, y, z, 1)
					}
				} else if neighbors == 3 {
					d.set(x, y, z, 1)
				}
			}
		}
	}
	return d
}

func around3(x, y, z int) []point3 {
	return []point3{
		{x - 1, y - 1, z - 1},
		{x - 1, y - 1, z},
		{x - 1, y - 1, z + 1},
		{x - 1, y, z - 1},
		{x - 1, y, z},
		{x - 1, y, z + 1},
		{x - 1, y + 1, z - 1},
		{x - 1, y + 1, z},
		{x - 1, y + 1, z + 1},
		{x, y - 1, z - 1},
		{x, y - 1, z},
		{x, y - 1, z + 1},
		{x, y, z - 1},
		// {x, y, z},
		{x, y, z + 1},
		{x, y + 1, z - 1},
		{x, y + 1, z},
		{x, y + 1, z + 1},
		{x + 1, y - 1, z - 1},
		{x + 1, y - 1, z},
		{x + 1, y - 1, z + 1},
		{x + 1, y, z - 1},
		{x + 1, y, z},
		{x + 1, y, z + 1},
		{x + 1, y + 1, z - 1},
		{x + 1, y + 1, z},
		{x + 1, y + 1, z + 1},
	}
}

type point4 struct {
	w, x, y, z int
}

type hypercube struct {
	max   int
	size  int
	cells []int
}

func newHypercube(max int) hypercube {
	c := hypercube{max: max, size: max*2 + 1}
	c.cells = make([]int, c.size*c.size*c.size*c.size)
	return c
}

func (c hypercube) at(w, x, y, z int) int {
	if util.IAbs(w) > c.max || util.IAbs(x) > c.max || util.IAbs(y) > c.max || util.IAbs(z) > c.max {
		return 0
	}
	offset := ((w + c.max) * c.size * c.size * c.size) + ((x + c.max) * c.size * c.size) + ((y + c.max) * c.size) + (z + c.max)
	return c.cells[offset]
}

func (c *hypercube) set(w, x, y, z, val int) {
	offset := ((w + c.max) * c.size * c.size * c.size) + ((x + c.max) * c.size * c.size) + ((y + c.max) * c.size) + (z + c.max)
	c.cells[offset] = val
}

func (c hypercube) step() hypercube {
	d := newHypercube(c.max)
	for w := -c.max; w < c.max; w++ {
		for x := -c.max; x < c.max; x++ {
			for y := -c.max; y < c.max; y++ {
				for z := -c.max; z < c.max; z++ {
					neighbors := 0
					for _, p := range around4(w, x, y, z) {
						if c.at(p.w, p.x, p.y, p.z) == 1 {
							neighbors++
						}
					}
					if c.at(w, x, y, z) == 1 {
						if neighbors == 2 || neighbors == 3 {
							d.set(w, x, y, z, 1)
						}
					} else if neighbors == 3 {
						d.set(w, x, y, z, 1)
					}
				}
			}
		}
	}
	return d
}

func around4(w, x, y, z int) []point4 {
	return []point4{
		{w - 1, x - 1, y - 1, z - 1},
		{w - 1, x - 1, y - 1, z},
		{w - 1, x - 1, y - 1, z + 1},
		{w - 1, x - 1, y, z - 1},
		{w - 1, x - 1, y, z},
		{w - 1, x - 1, y, z + 1},
		{w - 1, x - 1, y + 1, z - 1},
		{w - 1, x - 1, y + 1, z},
		{w - 1, x - 1, y + 1, z + 1},
		{w - 1, x, y - 1, z - 1},
		{w - 1, x, y - 1, z},
		{w - 1, x, y - 1, z + 1},
		{w - 1, x, y, z - 1},
		{w - 1, x, y, z},
		{w - 1, x, y, z + 1},
		{w - 1, x, y + 1, z - 1},
		{w - 1, x, y + 1, z},
		{w - 1, x, y + 1, z + 1},
		{w - 1, x + 1, y - 1, z - 1},
		{w - 1, x + 1, y - 1, z},
		{w - 1, x + 1, y - 1, z + 1},
		{w - 1, x + 1, y, z - 1},
		{w - 1, x + 1, y, z},
		{w - 1, x + 1, y, z + 1},
		{w - 1, x + 1, y + 1, z - 1},
		{w - 1, x + 1, y + 1, z},
		{w - 1, x + 1, y + 1, z + 1},

		{w, x - 1, y - 1, z - 1},
		{w, x - 1, y - 1, z},
		{w, x - 1, y - 1, z + 1},
		{w, x - 1, y, z - 1},
		{w, x - 1, y, z},
		{w, x - 1, y, z + 1},
		{w, x - 1, y + 1, z - 1},
		{w, x - 1, y + 1, z},
		{w, x - 1, y + 1, z + 1},
		{w, x, y - 1, z - 1},
		{w, x, y - 1, z},
		{w, x, y - 1, z + 1},
		{w, x, y, z - 1},
		// {w, x, y, z},
		{w, x, y, z + 1},
		{w, x, y + 1, z - 1},
		{w, x, y + 1, z},
		{w, x, y + 1, z + 1},
		{w, x + 1, y - 1, z - 1},
		{w, x + 1, y - 1, z},
		{w, x + 1, y - 1, z + 1},
		{w, x + 1, y, z - 1},
		{w, x + 1, y, z},
		{w, x + 1, y, z + 1},
		{w, x + 1, y + 1, z - 1},
		{w, x + 1, y + 1, z},
		{w, x + 1, y + 1, z + 1},

		{w + 1, x - 1, y - 1, z - 1},
		{w + 1, x - 1, y - 1, z},
		{w + 1, x - 1, y - 1, z + 1},
		{w + 1, x - 1, y, z - 1},
		{w + 1, x - 1, y, z},
		{w + 1, x - 1, y, z + 1},
		{w + 1, x - 1, y + 1, z - 1},
		{w + 1, x - 1, y + 1, z},
		{w + 1, x - 1, y + 1, z + 1},
		{w + 1, x, y - 1, z - 1},
		{w + 1, x, y - 1, z},
		{w + 1, x, y - 1, z + 1},
		{w + 1, x, y, z - 1},
		{w + 1, x, y, z},
		{w + 1, x, y, z + 1},
		{w + 1, x, y + 1, z - 1},
		{w + 1, x, y + 1, z},
		{w + 1, x, y + 1, z + 1},
		{w + 1, x + 1, y - 1, z - 1},
		{w + 1, x + 1, y - 1, z},
		{w + 1, x + 1, y - 1, z + 1},
		{w + 1, x + 1, y, z - 1},
		{w + 1, x + 1, y, z},
		{w + 1, x + 1, y, z + 1},
		{w + 1, x + 1, y + 1, z - 1},
		{w + 1, x + 1, y + 1, z},
		{w + 1, x + 1, y + 1, z + 1},
	}
}
