package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")

	timestamp := util.MustAtoi(input[0])
	schedule := input[1]

	bus, when := earliestBus(timestamp, schedule)
	fmt.Println("part1:", bus*(when-timestamp))

	fmt.Println("part2:", contestTimestamp(schedule))
}

func earliestBus(after int, schedule string) (bus int, when int) {
	buses := []int{}
	for _, f := range strings.Split(schedule, ",") {
		if f == "x" {
			continue
		}
		buses = append(buses, util.MustAtoi(f))
	}

	bus, when = 0, math.MaxInt64
	for _, id := range buses {
		nextDeparture := id * int(math.Ceil(float64(after)/float64(id)))
		if nextDeparture < when {
			bus, when = id, nextDeparture
		}
	}

	return
}

func contestTimestamp(schedule string) int {
	ids := []int{}       // route numbers
	residuals := []int{} // position of route in schedule
	limit := 1           // maximum timestamp to check
	for i, f := range strings.Split(schedule, ",") {
		if f == "x" {
			continue
		}

		// trying to find time where time mod id = residual
		id := util.MustAtoi(f)
		ids = append(ids, id)
		residuals = append(residuals, (id-(i%id))%id)

		// all of the busses appear to be prime/coprime so we should be able
		// to solve it before the product of them all (which would be when it
		// essentially loops back to zero)
		limit *= id
	}

	tm := ids[0]   // start when the first bus leaves
	step := ids[0] // increment for next timestamp to check
	busIdx := 1    // bus we're currently working on solving
	for {
		// have we solved the current bus?
		if tm%ids[busIdx] == residuals[busIdx] {
			// increase the step size to accomodate this bus
			step *= ids[busIdx]
			// now work on the next bus
			busIdx++
			// unless we're done
			if busIdx >= len(ids) {
				return tm
			}
		} else {
			// haven't solved yet, so lets try more
			tm += step
		}
		if tm > limit {
			panic("unable to solve contest")
		}
	}
}
