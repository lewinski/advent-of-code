package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type crates map[int][]string

func main() {
	records := util.Records("input.txt")

	layout := strings.Split(records[0], "\n")
	commands := strings.Split(records[1], "\n")

	part1 := parseLayout(layout)
	part2 := parseLayout(layout)

	for _, line := range commands {
		var count, src, dst int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &src, &dst)
		part1.moveSingle(count, src, dst)
		part2.moveBulk(count, src, dst)
	}

	fmt.Println("part1:", part1.top())
	fmt.Println("part2:", part2.top())
}

func parseLayout(lines []string) crates {
	layout := crates{}

	for _, line := range lines {
		for i := 0; i < len(line); i += 4 {
			if line[i] == '[' {
				stack := 1 + i/4
				layout[stack] = append([]string{line[i+1 : i+2]}, layout[stack]...)
			}
		}
	}

	return layout
}

func (c crates) moveSingle(count, src, dst int) {
	for i := 0; i < count; i++ {
		c.moveBulk(1, src, dst)
	}
}

func (c crates) moveBulk(count, src, dst int) {
	c[dst] = append(c[dst], c[src][len(c[src])-count:]...)
	c[src] = c[src][:len(c[src])-count]
}

func (c crates) top() string {
	message := ""
	for i := 1; i <= len(c); i++ {
		message += c[i][len(c[i])-1]
	}
	return message
}
