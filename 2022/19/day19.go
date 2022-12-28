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
		totalQuality += blueprint.id * maxGeodes(blueprint, minutes)
	}
	return totalQuality
}

func part2(blueprints []blueprint, minutes int) int {
	answer := 1
	for _, blueprint := range blueprints {
		answer *= maxGeodes(blueprint, minutes)
	}
	return answer
}

func maxGeodes(bp blueprint, minutes int) int {
	best := 0

	states := []state{
		{oreRobots: 1, minutes: minutes},
	}

	for len(states) > 0 {
		state := states[len(states)-1]
		states = states[0 : len(states)-1]

		if state.minutes == 0 {
			if state.geodes > best {
				best = state.geodes
			}
			continue
		}

		next := nextStates(bp, state)
		for _, s := range next {
			if estimateMaxGeodes(s) < best {
				continue
			}
			states = append(states, s)
		}
	}
	fmt.Printf("Blueprint %d: %d geodes\n", bp.id, best)
	return best
}

func estimateMaxGeodes(s state) int {
	return s.geodes +
		(s.minutes * s.geodeRobots) +
		((s.minutes * (s.minutes - 1)) / 2)
}

func nextStates(bp blueprint, cur state) []state {
	states := []state{}

	// should i build an ore robot?
	oreRobotNeeded := cur.oreRobots < bp.oreRobotOreCost ||
		cur.oreRobots < bp.clayRobotOreCost ||
		cur.oreRobots < bp.obsidianRobotOreCost ||
		cur.oreRobots < bp.geodeRobotOreCost
	if oreRobotNeeded {
		mins := 1
		if cur.ore < bp.oreRobotOreCost {
			mins += (bp.oreRobotOreCost - cur.ore + cur.oreRobots - 1) / cur.oreRobots
		}

		buildOre := state{
			minutes:        cur.minutes - mins,
			oreRobots:      cur.oreRobots + 1,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + (mins * cur.oreRobots) - bp.oreRobotOreCost,
			clay:           cur.clay + (mins * cur.clayRobots),
			obsidian:       cur.obsidian + (mins * cur.obsidianRobots),
			geodes:         cur.geodes + (mins * cur.geodeRobots),
		}
		if buildOre.minutes >= 0 {
			states = append(states, buildOre)
		}
	}

	// should i build a clay robot?
	clayRobotNeeded := cur.clayRobots < bp.obsidianRobotClayCost
	if clayRobotNeeded {
		mins := 1
		if cur.ore < bp.clayRobotOreCost {
			mins += (bp.clayRobotOreCost - cur.ore + cur.oreRobots - 1) / cur.oreRobots
		}

		buildClay := state{
			minutes:        cur.minutes - mins,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots + 1,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + (mins * cur.oreRobots) - bp.clayRobotOreCost,
			clay:           cur.clay + (mins * cur.clayRobots),
			obsidian:       cur.obsidian + (mins * cur.obsidianRobots),
			geodes:         cur.geodes + (mins * cur.geodeRobots),
		}
		if buildClay.minutes >= 0 {
			states = append(states, buildClay)
		}
	}

	// can i build an obsidian robot?
	obsidianRobotNeeded := cur.obsidianRobots < bp.geodeRobotObsidianCost
	if obsidianRobotNeeded && cur.clayRobots > 0 {
		mins := 1
		if cur.ore < bp.obsidianRobotOreCost || cur.clay < bp.obsidianRobotClayCost {
			mins += util.IMax(
				(bp.obsidianRobotOreCost-cur.ore+cur.oreRobots-1)/cur.oreRobots,
				(bp.obsidianRobotClayCost-cur.clay+cur.clayRobots-1)/cur.clayRobots,
			)
		}

		buildObsidian := state{
			minutes:        cur.minutes - mins,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots + 1,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + (mins * cur.oreRobots) - bp.obsidianRobotOreCost,
			clay:           cur.clay + (mins * cur.clayRobots) - bp.obsidianRobotClayCost,
			obsidian:       cur.obsidian + (mins * cur.obsidianRobots),
			geodes:         cur.geodes + (mins * cur.geodeRobots),
		}
		if buildObsidian.minutes >= 0 {
			states = append(states, buildObsidian)
		}
	}

	// can i build a geode robot?
	if cur.obsidianRobots > 0 {
		mins := 1
		if cur.ore < bp.geodeRobotOreCost || cur.obsidian < bp.geodeRobotObsidianCost {
			mins += util.IMax(
				(bp.geodeRobotOreCost-cur.ore+cur.oreRobots-1)/cur.oreRobots,
				(bp.geodeRobotObsidianCost-cur.obsidian+cur.obsidianRobots-1)/cur.obsidianRobots,
			)
		}

		buildGeode := state{
			minutes:        cur.minutes - mins,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots + 1,
			ore:            cur.ore + (mins * cur.oreRobots) - bp.geodeRobotOreCost,
			clay:           cur.clay + (mins * cur.clayRobots),
			obsidian:       cur.obsidian + (mins * cur.obsidianRobots) - bp.geodeRobotObsidianCost,
			geodes:         cur.geodes + (mins * cur.geodeRobots),
		}
		if buildGeode.minutes >= 0 {
			states = append(states, buildGeode)
		}
	}

	// do nothing
	if cur.geodeRobots > 1 {
		doNothing := state{
			minutes:        0,
			oreRobots:      cur.oreRobots,
			clayRobots:     cur.clayRobots,
			obsidianRobots: cur.obsidianRobots,
			geodeRobots:    cur.geodeRobots,
			ore:            cur.ore + (cur.minutes * cur.oreRobots),
			clay:           cur.clay + (cur.minutes * cur.clayRobots),
			obsidian:       cur.obsidian + (cur.minutes * cur.obsidianRobots),
			geodes:         cur.geodes + (cur.minutes * cur.geodeRobots),
		}
		states = append(states, doNothing)
	}

	return states
}
