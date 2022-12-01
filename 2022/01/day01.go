package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	records := util.Records("input.txt")
	totals := make([]int, len(records))
	for _, r := range records {
		cur := 0
		for _, s := range strings.Split(r, "\n") {
			cur += util.MustAtoi(s)
		}
		totals = append(totals, cur)
	}
	sort.IntSlice(totals).Sort()
	fmt.Println("part1:", totals[len(totals)-1])
	fmt.Println("part2:", totals[len(totals)-1]+totals[len(totals)-2]+totals[len(totals)-3])
}
