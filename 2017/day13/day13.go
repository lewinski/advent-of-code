package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(parseFirewall(lines)))
	// ~75 seconds
	fmt.Println("part2:", part2(parseFirewall(lines)))
}

type firewall struct {
	scanners   map[int]int
	directions map[int]int
	ranges     map[int]int
	layers     int
}

func newFirewall() firewall {
	return firewall{
		scanners:   map[int]int{},
		directions: map[int]int{},
		ranges:     map[int]int{},
		layers:     0,
	}
}

func parseFirewall(lines []string) firewall {
	f := newFirewall()

	for _, line := range lines {
		x := strings.SplitN(line, ": ", 2)
		f.addScanner(util.MustAtoi(x[0]), util.MustAtoi(x[1]))
	}

	return f
}

func (f firewall) dup() firewall {
	g := newFirewall()
	for k, v := range f.scanners {
		g.scanners[k] = v
	}
	for k, v := range f.directions {
		g.directions[k] = v
	}
	for k, v := range f.ranges {
		g.ranges[k] = v
	}
	g.layers = f.layers
	return g
}

func (f *firewall) addScanner(d, r int) {
	f.scanners[d] = 0
	f.directions[d] = 1
	f.ranges[d] = r
	f.layers = util.IMax(f.layers, d)
}

func (f *firewall) advance() {
	for j := range f.scanners {
		if f.scanners[j] == 0 {
			f.directions[j] = 1
		}
		if f.scanners[j] == f.ranges[j]-1 {
			f.directions[j] = -1
		}
		f.scanners[j] = f.scanners[j] + f.directions[j]
	}
}

func part1(f firewall) int {
	severity := 0

	for i := 0; i <= f.layers; i++ {
		if f.ranges[i] != 0 && f.scanners[i] == 0 {
			severity += i * f.ranges[i]
		}

		f.advance()
	}

	return severity
}

func part2(f firewall) int {
	delay := 0
delay:
	for {
		g := f.dup()
		for i := 0; i <= f.layers; i++ {
			if g.ranges[i] != 0 && g.scanners[i] == 0 {
				f.advance()
				delay++
				continue delay
			}

			g.advance()
		}
		return delay
	}
}
