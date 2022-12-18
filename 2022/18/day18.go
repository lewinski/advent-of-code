package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

const EMPTY = 0
const LAVA = 1
const WATER = 2

func main() {
	lines := util.Lines("input.txt")

	g := util.IntGrid3{}
	min, max := math.MaxInt, math.MinInt
	for _, line := range lines {
		f := strings.Split(line, ",")
		x := util.MustAtoi(f[0])
		y := util.MustAtoi(f[1])
		z := util.MustAtoi(f[2])
		min = util.IMin(min, util.IMin(x, util.IMin(y, z)))
		max = util.IMax(max, util.IMax(x, util.IMax(y, z)))
		g.SetCoords(x, y, z, LAVA)
	}
	min--
	max++

	surface := 0
	g.Each(func(p util.Point3, v int) {
		for _, t := range p.Touching() {
			if g.Get(t) == EMPTY {
				surface++
			}
		}
	})
	fmt.Println("part1:", surface)

	todo := []util.Point3{}
	todo = append(todo, util.Point3{min, min, min})
	seen := map[util.Point3]bool{}
	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		if seen[p] {
			continue
		}
		seen[p] = true
		g.Set(p, WATER)
		for _, t := range p.Touching() {
			if t[0] < min || t[0] > max {
				continue
			}
			if t[1] < min || t[1] > max {
				continue
			}
			if t[2] < min || t[2] > max {
				continue
			}
			if !seen[t] && g.Get(t) == 0 {
				todo = append(todo, t)
			}
		}
	}

	outer := 0
	g.Each(func(p util.Point3, v int) {
		if v == LAVA {
			for _, t := range p.Touching() {
				if g.Get(t) == WATER {
					outer++
				}
			}
		}
	})
	fmt.Println("part2:", outer)
}
