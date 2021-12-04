package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func makeCounts(numbers []string) []int {
	cnt := len(numbers[0])

	counts := make([]int, cnt)

	for _, number := range numbers {
		for i := 0; i < cnt; i++ {
			if number[i] == '1' {
				counts[i]++
			}
		}
	}

	return counts
}

func part1(numbers []string) int64 {
	counts := makeCounts(numbers)

	var gamma, epsilon strings.Builder

	for i := 0; i < len(numbers[0]); i++ {
		if counts[i] < len(numbers)-counts[i] {
			epsilon.WriteRune('1')
			gamma.WriteRune('0')
		} else {
			epsilon.WriteRune('0')
			gamma.WriteRune('1')
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma.String(), 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon.String(), 2, 64)

	return gammaInt * epsilonInt
}

func part2(numbers []string) int64 {
	pos := 0
	candidates := numbers
	for len(candidates) > 1 {
		counts := makeCounts(candidates)

		want := byte('0')
		if counts[pos] >= len(candidates)-counts[pos] {
			want = byte('1')
		}
		newCandidates := []string{}
		for _, s := range candidates {
			if s[pos] == want {
				newCandidates = append(newCandidates, s)
			}
		}
		candidates = newCandidates
		pos++
	}
	oxygen, _ := strconv.ParseInt(candidates[0], 2, 64)

	pos = 0
	candidates = numbers
	for len(candidates) > 1 {
		counts := makeCounts(candidates)

		want := byte('0')
		if counts[pos] < len(candidates)-counts[pos] {
			want = byte('1')
		}
		newCandidates := []string{}
		for _, s := range candidates {
			if s[pos] == want {
				newCandidates = append(newCandidates, s)
			}
		}

		candidates = newCandidates
		pos++
	}
	co2, _ := strconv.ParseInt(candidates[0], 2, 64)

	return oxygen * co2
}
