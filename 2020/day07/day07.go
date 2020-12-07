package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	bags := parseBags(input)

	mine := "shiny gold"

	fmt.Println("part1:", len(bagsContaining(mine, bags)))

	count := 0
	for _, c := range contentsOf(mine, bags) {
		count += c
	}
	fmt.Println("part2:", count)
}

type bag struct {
	name     string
	children *map[string]int
}

func parseBags(lines []string) map[string]bag {
	bags := map[string]bag{}

	for _, line := range lines {
		line = strings.TrimRight(line, ".")
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")

		pair := strings.SplitN(line, " contain ", 2)
		name := pair[0]

		if pair[1] == "no other" {
			bags[name] = bag{name: name}
		} else {
			children := map[string]int{}
			for _, child := range strings.Split(pair[1], ", ") {
				count, _ := strconv.Atoi(child[0:1])
				children[child[2:]] = count
			}
			bags[name] = bag{name: name, children: &children}
		}
	}

	return bags
}

func contentsOf(name string, bags map[string]bag) map[string]int {
	contents := map[string]int{}

	if bags[name].children != nil {
		for child, count := range *bags[name].children {
			if _, ok := contents[child]; ok {
				contents[child] += count
			} else {
				contents[child] = count
			}

			for childName, childCount := range contentsOf(child, bags) {
				if _, ok := contents[childName]; ok {
					contents[childName] += count * childCount
				} else {
					contents[childName] = count * childCount
				}
			}
		}
	}

	return contents
}

func bagsContaining(goal string, bags map[string]bag) []string {
	containing := []string{}

	for name := range bags {
		contents := contentsOf(name, bags)
		if _, ok := contents[goal]; ok {
			containing = append(containing, name)
		}
	}

	return containing
}
