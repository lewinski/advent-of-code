package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	values := util.IntLines("input.txt")
	preamble := 25

	invalidValue := findInvalidValue(values, preamble)
	fmt.Println("part1:", invalidValue)
	fmt.Println("part2:", findEncryptionWeakness(values, invalidValue))
}

func findInvalidValue(values []int, preamble int) int {
search:
	for i := preamble; i < len(values); i++ {
		found := false
		for j := i - preamble; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if values[i] == values[j]+values[k] {
					found = true
					continue search
				}
			}
		}
		if !found {
			return values[i]
		}
	}
	panic("values are all valid")
}

func findEncryptionWeakness(values []int, goal int) int {
	for i := 0; i < len(values); i++ {
		sum := 0
		min, max := math.MaxInt32, math.MinInt32
		for j := i; sum < goal; j++ {
			sum += values[j]
			if values[j] < min {
				min = values[j]
			}
			if values[j] > max {
				max = values[j]
			}
			if sum == goal {
				return min + max
			}
		}
	}
	panic("can't find weakness")
}
