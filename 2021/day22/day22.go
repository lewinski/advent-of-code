package main

import (
	"fmt"
	"regexp"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	steps := parseInput(lines)

	fmt.Println("part1:", part1(steps))
	fmt.Println("part2:", part2(steps))
}

type rebootStep struct {
	operation int
	x1, x2    int
	y1, y2    int
	z1, z2    int
}

func parseInput(lines []string) []rebootStep {
	steps := make([]rebootStep, 0, len(lines))

	re := regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
	for _, line := range lines {
		if !re.MatchString(line) {
			panic(line)
		}
		match := re.FindStringSubmatch(line)

		op := 0
		if match[1] == "on" {
			op = 1
		}
		step := rebootStep{
			operation: op,
			x1:        util.MustAtoi(match[2]),
			x2:        util.MustAtoi(match[3]),
			y1:        util.MustAtoi(match[4]),
			y2:        util.MustAtoi(match[5]),
			z1:        util.MustAtoi(match[6]),
			z2:        util.MustAtoi(match[7]),
		}
		steps = append(steps, step)
	}
	return steps
}

func part1(steps []rebootStep) int {
	g := util.IntGrid3{}

	for _, step := range steps {
		if step.x1 < -50 || step.x2 > 50 || step.y1 < -50 || step.y2 > 50 || step.z1 < -50 || step.z2 > 50 {
			continue
		}

		for x := step.x1; x <= step.x2; x++ {
			for y := step.y1; y <= step.y2; y++ {
				for z := step.z1; z <= step.z2; z++ {
					g.SetCoords(x, y, z, step.operation)
				}
			}
		}
	}

	cnt := 0
	g.Each(func(p util.Point3, x int) {
		cnt += x
	})

	return cnt
}

func intervalIntersection(a1, a2, b1, b2 int) (int, int, bool) {
	if a2 < b1 {
		return 0, 0, false
	}
	if a1 > b2 {
		return 0, 0, false
	}
	c1 := util.IMin(util.IMax(a1, b1), b2)
	c2 := util.IMin(util.IMax(a2, b1), b2)
	return c1, c2, true
}

func boxCount(step rebootStep, remaining []rebootStep) int {
	// turn the whole box on
	xc := step.x2 - step.x1 + 1
	yc := step.y2 - step.y1 + 1
	zc := step.z2 - step.z1 + 1
	on := xc * yc * zc

	// find other boxes that intersect with this one
	intersections := []rebootStep{}
	for _, other := range remaining {
		x1, x2, xoverlaps := intervalIntersection(step.x1, step.x2, other.x1, other.x2)
		y1, y2, yoverlaps := intervalIntersection(step.y1, step.y2, other.y1, other.y2)
		z1, z2, zoverlaps := intervalIntersection(step.z1, step.z2, other.z1, other.z2)

		// if there is an overlap in all three directions, then the boxes do intersect
		// and the ranges above define the box of the intersection. add it to our list
		// of items to reprocess later
		if xoverlaps && yoverlaps && zoverlaps {
			intersections = append(intersections, rebootStep{
				operation: other.operation,
				x1:        x1,
				x2:        x2,
				y1:        y1,
				y2:        y2,
				z1:        z1,
				z2:        z2,
			})
		}
	}

	// now go over all of the intersections we found and subtract them out of this count
	// they will get added back in correctly when the "original" gets handled
	for i, step := range intersections {
		on -= boxCount(step, intersections[i+1:])
	}

	return on
}

func part2(steps []rebootStep) int {
	total := 0

	for i := 0; i < len(steps); i++ {
		// skip calculations for "off" steps
		if steps[i].operation == 0 {
			continue
		}

		// add the number of cells turned on by this box that won't be considered by future boxes
		total += boxCount(steps[i], steps[i+1:])
	}

	return total
}
