package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	template := lines[0]
	rules := map[string]string{}
	for _, line := range lines[2:] {
		f := strings.Split(line, " -> ")
		rules[f[0]] = f[1]
	}

	fmt.Println("part1:", part1(template, rules))
	fmt.Println("part2:", part2(template, rules))
}

func part1(template string, rules map[string]string) int {
	cur := template
	for i := 0; i < 10; i++ {
		var next strings.Builder
		next.WriteByte(cur[0])
		for j := 0; j < len(cur)-1; j++ {
			next.WriteString(rules[cur[j:j+2]])
			next.WriteByte(cur[j+1])
		}
		cur = next.String()
	}

	counts := map[rune]int{}
	for _, r := range cur {
		counts[r]++
	}

	min, max := math.MaxInt, math.MinInt
	for _, c := range counts {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	return max - min
}

func part2(template string, rules map[string]string) int {
	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	betterRules := map[string][2]string{}
	for pair, middle := range rules {
		betterRules[pair] = [2]string{
			fmt.Sprintf("%c%s", pair[0], middle),
			fmt.Sprintf("%s%c", middle, pair[1]),
		}
	}

	for i := 0; i < 40; i++ {
		next := map[string]int{}
		for pair, count := range pairs {
			for _, p := range betterRules[pair] {
				next[p] += count
			}
		}
		pairs = next
	}

	counts := map[byte]int{}
	counts[template[0]] = 1
	for pair, count := range pairs {
		counts[pair[1]] += count
	}

	min, max := math.MaxInt, math.MinInt
	for _, c := range counts {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	return max - min
}
