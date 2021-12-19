package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	// example: target area: x=20..30, y=-10..-5
	//target := box{util.Point2{20, -10}, util.Point2{30, -5}}

	// input: target area: x=85..145, y=-163..-108
	target := box{util.Point2{85, -163}, util.Point2{145, -108}}

	// shoot any faster than this and we probably overshoot on the first steps
	max := util.IMax(
		util.IMax(util.IAbs(target.min[0]), util.IAbs(target.min[1])),
		util.IMax(util.IAbs(target.max[0]), util.IAbs(target.max[1])),
	)

	fmt.Println("part1:", part1(target, max))
	fmt.Println("part1:", part2(target, max))
}

type box struct {
	min util.Point2
	max util.Point2
}

func maxHeightForShot(v util.Point2, target box) int {
	height := math.MinInt

	pos := util.Origin2()
	for i := 0; i < 1000; i++ {
		if pos[0] >= target.min[0] &&
			pos[0] <= target.max[0] &&
			pos[1] >= target.min[1] &&
			pos[1] <= target.max[1] {
			// target hit
			return height
		}
		if pos[0] > target.max[0] {
			// overshoot
			return math.MinInt
		}

		pos = pos.Offset(v)
		if pos[1] > height {
			height = pos[1]
		}

		if v[0] > 0 {
			v[0]--
		} else if v[0] < 0 {
			v[0]++
		}
		v[1]--
	}

	return math.MinInt
}

func part1(target box, maxVelocity int) int {
	best := math.MinInt

	for x := 0; x < maxVelocity; x++ {
		for y := -maxVelocity; y < maxVelocity; y++ {
			h := maxHeightForShot(util.Point2{x, y}, target)
			if h > best {
				best = h
			}
		}
	}

	return best
}

func part2(target box, maxVelocity int) int {
	hits := 0

	for x := 0; x < maxVelocity; x++ {
		for y := -maxVelocity; y < maxVelocity; y++ {
			h := maxHeightForShot(util.Point2{x, y}, target)
			if h > math.MinInt {
				hits++
			}
		}
	}

	return hits
}
