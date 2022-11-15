package main

import (
	"fmt"
)

func main() {
	maxScore := 0

	for f := 0; f <= 100; f++ {
		for c := 0; c <= 100-f; c++ {
			for b := 0; b <= 100-f-c; b++ {
				s := 100 - f - c - b

				score := score(f, c, b, s)
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}
	fmt.Println("part1:", maxScore)

	maxScore = 0
	for f := 0; f <= 100; f++ {
		for c := 0; c <= 100-f; c++ {
			for b := 0; b <= 100-f-c; b++ {
				s := 100 - f - c - b

				if calories(f, c, b, s) != 500 {
					continue
				}

				score := score(f, c, b, s)
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}

	fmt.Println("part2:", maxScore)
}

/*
Frosting:     capacity 4,  durability -2, flavor 0,  texture 0, calories 5
Candy:        capacity 0,  durability 5,  flavor -1, texture 0, calories 8
Butterscotch: capacity -1, durability 0,  flavor 5,  texture 0, calories 6
Sugar:        capacity 0,  durability 0,  flavor -2, texture 2, calories 1
*/

func score(f, c, b, s int) int {
	cap := 4*f - b
	dur := -2*f + 5*c
	fla := -c + 5*b - 2*s
	tex := 2 * s
	if cap < 0 || dur < 0 || fla < 0 || tex < 0 {
		return 0
	}
	return cap * dur * fla * tex
}

func calories(f, c, b, s int) int {
	return 5*f + 8*c + 6*b + s
}
