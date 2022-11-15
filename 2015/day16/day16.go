package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	tests := map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas":     "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}
	fmt.Println("part1:", match(lines, tests))

	tests["cats"] = ">7"
	tests["trees"] = ">3"
	tests["pomeranians"] = "<3"
	tests["goldfish"] = "<5"
	fmt.Println("part2:", match(lines, tests))
}

func match(lines []string, tests map[string]string) string {
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		things := strings.Split(parts[1], ", ")

		ok := true
		for _, thing := range things {
			kv := strings.Split(thing, ": ")
			if strings.HasPrefix(tests[kv[0]], ">") {
				if util.MustAtoi(kv[1]) <= util.MustAtoi(tests[kv[0]][1:]) {
					ok = false
					break
				}
			} else if strings.HasPrefix(tests[kv[0]], "<") {
				if util.MustAtoi(kv[1]) >= util.MustAtoi(tests[kv[0]][1:]) {
					ok = false
					break
				}
			} else {
				if kv[1] != tests[kv[0]] {
					ok = false
					break
				}
			}
		}

		if ok {
			return parts[0]
		}
	}
	panic("oh no")
}
