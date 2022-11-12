package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type program struct {
	name     string
	weight   int
	children []string
}

func main() {
	programs := map[string]program{}
	parents := map[string]string{}

	lines := util.Lines("input.txt")
	for _, line := range lines {
		fields := regexp.MustCompile(" -> ").Split(line, 2)

		p := program{}

		fmt.Sscanf(fields[0], "%s (%d)", &p.name, &p.weight)
		if len(fields) > 1 {
			p.children = strings.Split(fields[1], ", ")
			for _, name := range p.children {
				parents[name] = p.name
			}
		}

		programs[p.name] = p
	}

	root := part1(programs, parents)
	fmt.Println("part1:", root)
	fmt.Println("part2:", part2(root, weight(root, programs), programs))
}

func part1(programs map[string]program, parents map[string]string) string {
	for name := range programs {
		if _, ok := parents[name]; !ok {
			return name
		}
	}
	panic("didn't solve part 1")
}

func weight(program string, programs map[string]program) int {
	w := programs[program].weight
	for _, name := range programs[program].children {
		w += weight(name, programs)
	}
	return w
}

func part2(root string, goal int, programs map[string]program) int {
	type tweight struct {
		name   string
		weight int
	}

	weights := []tweight{}

	for _, name := range programs[root].children {
		w := weight(name, programs)
		goal -= w
		weights = append(weights, tweight{name, w})
	}

	if len(weights) < 2 {
		panic("unexpected")
	}

	if weights[0].weight == weights[1].weight && weights[0].weight != weights[2].weight {
		return part2(weights[2].name, weights[1].weight, programs)
	}
	if weights[0].weight == weights[2].weight && weights[0].weight != weights[1].weight {
		return part2(weights[1].name, weights[0].weight, programs)
	}
	if weights[1].weight == weights[2].weight && weights[0].weight != weights[1].weight {
		return part2(weights[0].name, weights[1].weight, programs)
	}

	for i := 3; i < len(weights); i++ {
		if weights[i].weight != weights[i-1].weight {
			return part2(weights[i].name, weights[i-1].weight, programs)
		}
	}

	return goal
}
