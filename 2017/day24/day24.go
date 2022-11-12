package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	parts := map[util.Point2]bool{}
	for _, line := range lines {
		f := strings.Split(line, "/")
		parts[util.Point2{util.MustAtoi(f[0]), util.MustAtoi(f[1])}] = false
	}

	fmt.Println("part1:", part1(parts))
	fmt.Println("part2:", part2(parts))
}

func part1(parts map[util.Point2]bool) int {
	return strongest(parts, 0, 0)
}

func strongest(parts map[util.Point2]bool, port, strength int) int {
	best := strength

	for p, used := range parts {
		if used {
			continue
		}

		var need int
		if p[0] == port {
			need = p[1]
		} else if p[1] == port {
			need = p[0]
		} else {
			continue
		}

		parts[p] = true
		s := strength + p[0] + p[1]
		sub := strongest(parts, need, s)
		parts[p] = false

		if sub > best {
			best = sub
		}
	}

	return best
}

func part2(parts map[util.Point2]bool) int {
	s, _ := longest(parts, 0, 0, 0)
	return s
}

func longest(parts map[util.Point2]bool, port, strength, depth int) (int, int) {
	bestStrength := strength
	bestDepth := depth

	for p, used := range parts {
		if used {
			continue
		}

		var need int
		if p[0] == port {
			need = p[1]
		} else if p[1] == port {
			need = p[0]
		} else {
			continue
		}

		parts[p] = true
		s := strength + p[0] + p[1]
		d := depth + 1
		subStr, subDep := longest(parts, need, s, d)
		parts[p] = false

		if subDep > bestDepth {
			bestDepth = subDep
			bestStrength = subStr
		} else if subDep == bestDepth && subStr > bestStrength {
			bestStrength = subStr
		}
	}

	return bestStrength, bestDepth
}
