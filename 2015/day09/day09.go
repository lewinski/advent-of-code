package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	places := map[string]bool{}
	distances := map[string]int{}

	for _, line := range lines {
		var place1, place2 string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &place1, &place2, &distance)
		places[place1] = true
		places[place2] = true
		distances[place1+"-"+place2] = distance
		distances[place2+"-"+place1] = distance
	}

	fmt.Println("part1:", part1(places, distances))
	fmt.Println("part2:", part2(places, distances))
}

func part1(places map[string]bool, distances map[string]int) int {
	min := math.MaxInt64
	seen := map[string]bool{}
	for place := range places {
		seen[place] = true
		min = util.IMin(min, shortestPath(place, places, distances, seen, 0))
		delete(seen, place)
	}
	return min
}

func shortestPath(place string, places map[string]bool, distances map[string]int, seen map[string]bool, distance int) int {
	if len(places) == len(seen) {
		return distance
	}

	min := math.MaxInt64
	for next := range places {
		if seen[next] {
			continue
		}
		seen[next] = true
		min = util.IMin(min, shortestPath(next, places, distances, seen, distance+distances[place+"-"+next]))
		delete(seen, next)
	}
	return min
}

func part2(places map[string]bool, distances map[string]int) int {
	max := 0
	seen := map[string]bool{}
	for place := range places {
		seen[place] = true
		max = util.IMax(max, longestPath(place, places, distances, seen, 0))
		delete(seen, place)
	}
	return max
}

func longestPath(place string, places map[string]bool, distances map[string]int, seen map[string]bool, distance int) int {
	if len(places) == len(seen) {
		return distance
	}

	max := 0
	for next := range places {
		if seen[next] {
			continue
		}
		seen[next] = true
		max = util.IMax(max, longestPath(next, places, distances, seen, distance+distances[place+"-"+next]))
		delete(seen, next)
	}
	return max
}
