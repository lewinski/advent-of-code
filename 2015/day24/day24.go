package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	p := util.IntLines("input.txt")
	fmt.Println("part1:", guess(p, 3))
	fmt.Println("part1:", guess(p, 4))
}

// not general purpose, but seems to technically work for my input
// find all subsets of p that sum to 1/groups of the total weight
// and return the best quantum entanglement amongst the smallest size subsets
// which assumes:
// - the smallest set has 4, 5 or 6 packages
// - that the remaining packages can also be equally split
func guess(p []int, groups int) int {
	total := 0
	for _, p := range p {
		total += p
	}
	goal := total / groups

	bestSize := math.MaxInt
	bestQE := math.MaxInt

	for a := 0; a < len(p); a++ {
		for b := a + 1; b < len(p); b++ {
			for c := b + 1; c < len(p); c++ {
				for d := c + 1; d < len(p); d++ {

					if p[a]+p[b]+p[c]+p[d] == goal {
						size := 4
						qe := p[a] * p[b] * p[c] * p[d]
						if size < bestSize {
							bestSize = size
							bestQE = qe
						} else if size == bestSize && qe < bestQE {
							bestQE = qe
						}
					}

					for e := d + 1; e < len(p); e++ {
						if p[a]+p[b]+p[c]+p[d]+p[e] == goal {
							size := 5
							qe := p[a] * p[b] * p[c] * p[d] * p[e]
							if size < bestSize {
								bestSize = size
								bestQE = qe
							} else if size == bestSize && qe < bestQE {
								bestQE = qe
							}
						}

						for f := e + 1; f < len(p); f++ {
							if p[a]+p[b]+p[c]+p[d]+p[e]+p[f] == goal {
								size := 6
								qe := p[a] * p[b] * p[c] * p[d] * p[e] * p[f]
								if size < bestSize {
									bestSize = size
									bestQE = qe
								} else if size == bestSize && qe < bestQE {
									bestQE = qe
								}
							}
						}
					}
				}
			}
		}
	}

	return bestQE
}
