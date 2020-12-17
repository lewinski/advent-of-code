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
		cube = step1(cube)
	}

	cubeCount := 0
	cube.Each(func(p util.Point3, x int) {
		if x > 0 {
			cubeCount++
		}
	})
	fmt.Println("part1:", cubeCount)

	hypercube := parseInput2(input, steps)
	for i := 0; i < steps; i++ {
		hypercube = step2(hypercube)
	}

	hypercubeCount := 0
	hypercube.Each(func(p util.Point4, x int) {
		if x > 0 {
			hypercubeCount++
		}
	})
	fmt.Println("part2:", hypercubeCount)

}

func parseInput1(lines []string, steps int) util.IntGrid3 {
	g := util.IntGrid3{}
	for x, line := range lines {
		for y, c := range line {
			val := 0
			if c == '#' {
				val = 1
			}
			g[util.Point3{x, y, 0}] = val
		}
	}
	return g
}

func step1(g util.IntGrid3) util.IntGrid3 {
	accum := util.IntGrid3{}
	g.Each(func(p util.Point3, x int) {
		if x > 0 {
			for _, a := range p.Around() {
				accum[a]++
			}
		}
	})
	result := util.IntGrid3{}
	accum.Each(func(p util.Point3, x int) {
		if (x == 3) || (g[p] > 0 && x == 2) {
			result[p] = 1
		}
	})
	return result
}

func parseInput2(lines []string, steps int) util.IntGrid4 {
	g := util.IntGrid4{}
	for x, line := range lines {
		for y, c := range line {
			val := 0
			if c == '#' {
				val = 1
			}
			g[util.Point4{x, y, 0}] = val
		}
	}
	return g
}

func step2(g util.IntGrid4) util.IntGrid4 {
	accum := util.IntGrid4{}
	g.Each(func(p util.Point4, x int) {
		if x > 0 {
			for _, a := range p.Around() {
				accum[a]++
			}
		}
	})
	result := util.IntGrid4{}
	accum.Each(func(p util.Point4, x int) {
		if (x == 3) || (g[p] > 0 && x == 2) {
			result[p] = 1
		}
	})
	return result
}
