package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type blueprint struct {
	id                     int
	oreRobotOreCost        int
	clayRobotOreCost       int
	obsidianRobotOreCost   int
	obsidianRobotClayCost  int
	geodeRobotOreCost      int
	geodeRobotObsidianCost int
}

type state struct {
	minutes                                            int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
	ore, clay, obsidian, geodes                        int
}

func main() {
	lines := util.Lines("input.txt")

	blueprints := []blueprint{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, ":", "")
		blueprint := blueprint{}
		fmt.Sscanf(
			line,
			"Blueprint %d Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&blueprint.id,
			&blueprint.oreRobotOreCost,
			&blueprint.clayRobotOreCost,
			&blueprint.obsidianRobotOreCost,
			&blueprint.obsidianRobotClayCost,
			&blueprint.geodeRobotOreCost,
			&blueprint.geodeRobotObsidianCost,
		)
		blueprints = append(blueprints, blueprint)
	}

	fmt.Println("part1:", part1(blueprints, 24))
	fmt.Println("part2:", part2(blueprints[0:util.IMin(len(blueprints), 3)], 32))
}

func part1(blueprints []blueprint, minutes int) int {
	totalQuality := 0
	for _, blueprint := range blueprints {
		best := map[int]int{}

		states := []state{
			{oreRobots: 1},
		}

		for len(states) > 0 {
			state := states[len(states)-1]
			states = states[0 : len(states)-1]

			if state.geodes > best[state.minutes] {
				best[state.minutes] = state.geodes
			} else if estimateGeodes(state, minutes) < best[minutes] {
				continue
			}

			if state.minutes == minutes {
				continue
			}

			next := nextStates(blueprint, state)
			for _, s := range next {
				if estimateGeodes(s, minutes) < best[minutes] {
					continue
				}
				states = append(states, s)
			}
		}
		fmt.Printf("Blueprint %d: %d geodes\n", blueprint.id, best[minutes])
		totalQuality += blueprint.id * best[minutes]
	}

	return totalQuality
}

func part2(blueprints []blueprint, minutes int) int {
	answer := 1
	for _, blueprint := range blueprints {
		best := map[int]int{}

		states := []state{
			{oreRobots: 1},
		}

		for len(states) > 0 {
			state := states[len(states)-1]
			states = states[0 : len(states)-1]

			if state.geodes > best[state.minutes] {
				best[state.minutes] = state.geodes
			} else if estimateGeodes(state, minutes) < best[minutes] {
				continue
			}

			if state.minutes == minutes {
				continue
			}

			next := nextStates(blueprint, state)
			for _, s := range next {
				if estimateGeodes(s, minutes) < best[minutes] {
					continue
				}
				states = append(states, s)
			}

		}
		fmt.Printf("Blueprint %d: %d geodes\n", blueprint.id, best[minutes])
		answer *= best[minutes]
	}

	return answer
}

func estimateGeodes(s state, minutes int) int {
	remaining := minutes - s.minutes
	return s.geodes +
		(remaining * s.geodeRobots) +
		(remaining*(remaining-1))/2
}

func nextStates(bp blueprint, cur state) []state {
	states := []state{}

	var buildOre, buildClay, buildObsidian, buildGeode, doNothing *state

	// can i build an ore robot?
	oreRobotNeeded := cur.oreRobots < bp.oreRobotOreCost ||
		cur.oreRobots < bp.clayRobotOreCost ||
		cur.oreRobots < bp.obsidianRobotOreCost ||
		cur.oreRobots < bp.geodeRobotOreCost
	if oreRobotNeeded && cur.ore >= bp.oreRobotOreCost {
		buildOre = &state{
			minutes:        cur.minutes + 1,
			oreRobots:      cur.oreRobots + 1,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + cur.oreRobots - bp.oreRobotOreCost,
			clay:           cur.clay + cur.clayRobots,
			obsidian:       cur.obsidian + cur.obsidianRobots,
			geodes:         cur.geodes + cur.geodeRobots,
		}
	}

	// can i build a clay robot?
	clayRobotNeeded := cur.clayRobots < bp.obsidianRobotClayCost
	if clayRobotNeeded && cur.ore >= bp.clayRobotOreCost {
		buildClay = &state{
			minutes:        cur.minutes + 1,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots + 1,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + cur.oreRobots - bp.clayRobotOreCost,
			clay:           cur.clay + cur.clayRobots,
			obsidian:       cur.obsidian + cur.obsidianRobots,
			geodes:         cur.geodes + cur.geodeRobots,
		}
	}

	// can i build an obsidian robot?
	obsidianRobotNeeded := cur.obsidianRobots < bp.geodeRobotObsidianCost
	if obsidianRobotNeeded && cur.ore >= bp.obsidianRobotOreCost && cur.clay >= bp.obsidianRobotClayCost {
		buildObsidian = &state{
			minutes:        cur.minutes + 1,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots + 1,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + cur.oreRobots - bp.obsidianRobotOreCost,
			clay:           cur.clay + cur.clayRobots - bp.obsidianRobotClayCost,
			obsidian:       cur.obsidian + cur.obsidianRobots,
			geodes:         cur.geodes + cur.geodeRobots,
		}
	}

	// can i build a geode robot?
	if cur.ore >= bp.geodeRobotOreCost && cur.obsidian >= bp.geodeRobotObsidianCost {
		buildGeode = &state{
			minutes:        cur.minutes + 1,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots + 1,
			ore:            cur.ore + cur.oreRobots - bp.geodeRobotOreCost,
			clay:           cur.clay + cur.clayRobots,
			obsidian:       cur.obsidian + cur.obsidianRobots - bp.geodeRobotObsidianCost,
			geodes:         cur.geodes + cur.geodeRobots,
		}
	}

	// do nothing
	if oreRobotNeeded || clayRobotNeeded || obsidianRobotNeeded || buildGeode != nil {
		doNothing = &state{
			minutes:        cur.minutes + 1,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + cur.oreRobots,
			clay:           cur.clay + cur.clayRobots,
			obsidian:       cur.obsidian + cur.obsidianRobots,
			geodes:         cur.geodes + cur.geodeRobots,
		}
	}

	if doNothing != nil {
		states = append(states, *doNothing)
	}
	if buildOre != nil {
		states = append(states, *buildOre)
	}
	if buildClay != nil {
		states = append(states, *buildClay)
	}
	if buildObsidian != nil {
		states = append(states, *buildObsidian)
	}
	if buildGeode != nil {
		states = append(states, *buildGeode)
	}

	return states
}
