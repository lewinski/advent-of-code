package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type replacement struct {
	from, to string
}

func main() {
	lines := util.Lines("input.txt")

	var replacements []replacement
	var medicine string
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "=>") {
			parts := strings.Split(line, " => ")
			replacements = append(replacements, replacement{parts[0], parts[1]})
		} else {
			medicine = line
		}
	}

	fmt.Println("part1:", part1(medicine, replacements))
	fmt.Println("part2:", part2(medicine, replacements))

}

func nextSet(seeds map[string]bool, replacements []replacement) map[string]bool {
	next := map[string]bool{}
	for seed := range seeds {
		for i := range seed {
			for _, r := range replacements {
				if strings.HasPrefix(seed[i:], r.from) {
					next[seed[:i]+r.to+seed[i+len(r.from):]] = true
				}
			}
		}
	}
	return next
}

func part1(medicine string, replacements []replacement) int {
	return len(nextSet(map[string]bool{medicine: true}, replacements))
}

func part2(medicine string, replacements []replacement) int {
	// work backwards
	reverseReplacements := []replacement{}
	for _, r := range replacements {
		reverseReplacements = append(reverseReplacements, replacement{r.to, r.from})
	}

	depth := 1
	seeds := nextSet(map[string]bool{medicine: true}, reverseReplacements)
	for {
		// done
		if _, ok := seeds["e"]; ok {
			return depth
		}

		// calculate complete next set
		next := nextSet(seeds, reverseReplacements)
		depth++

		// prune to the set of shortest molecules (reducing most quickly)
		// 100 seems reasonable, but it works as low as 3 for my input
		limit := 100
		if len(next) == 0 {
			panic("oh no")
		} else if len(next) < limit {
			seeds = next
		} else {
			all := []string{}
			for s := range next {
				all = append(all, s)
			}

			sort.Slice(all, func(i, j int) bool {
				return len(all[i]) < len(all[j])
			})

			seeds = map[string]bool{}
			for i := 0; i < limit; i++ {
				seeds[all[i]] = true
			}
		}
	}
}
