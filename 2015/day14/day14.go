package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

type reindeer struct {
	name             string
	speed, fly, rest int
}

func main() {
	lines := util.Lines("input.txt")

	deer := []reindeer{}

	for _, line := range lines {
		var d reindeer
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &d.name, &d.speed, &d.fly, &d.rest)
		deer = append(deer, d)
	}

	fmt.Println("part1:", part1(deer, 2503))
	fmt.Println("part2:", part2(deer, 2503))
}

func part1(deer []reindeer, time int) int {
	bestScore := 0

	for _, d := range deer {
		remain := time % (d.fly + d.rest)
		cycles := (time - remain) / (d.fly + d.rest)

		score := cycles * d.speed * d.fly
		if remain > d.fly {
			score += d.speed * d.fly
		} else {
			score += d.speed * remain
		}

		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
}

func part2(deer []reindeer, time int) int {
	type state struct {
		distance int
		points   int
		flying   bool
		counter  int
	}
	scores := map[string]state{}

	for _, d := range deer {
		scores[d.name] = state{flying: true}
	}

	for i := 0; i < time; i++ {
		best := 0

		for _, d := range deer {
			s := scores[d.name]

			// advance
			if s.flying {
				s.distance += d.speed
			}

			// are we leading?
			if s.distance > best {
				best = s.distance
			}

			// check if we're tired or rested
			s.counter++
			if s.flying && s.counter == d.fly {
				s.flying = false
				s.counter = 0
			} else if !s.flying && s.counter == d.rest {
				s.flying = true
				s.counter = 0
			}

			// save updates
			scores[d.name] = s
		}

		// update everyone who is leading
		for _, d := range deer {
			s := scores[d.name]
			if s.distance == best {
				s.points++
				scores[d.name] = s
			}
		}
	}

	best := 0
	for _, s := range scores {
		if s.points > best {
			best = s.points
		}
	}

	return best
}
