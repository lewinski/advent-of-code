package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Records("input.txt")

	fieldMap := parseClasses(input[0])
	myTickets := parseTickets(input[1])
	otherTickets := parseTickets(input[2])

	validTickets, errorRate := validTickets(otherTickets, fieldMap)
	fmt.Println("part1:", errorRate)

	fieldPos := solveFieldMapping(validTickets, fieldMap)
	part2 := 1
	for name, position := range fieldPos {
		if strings.HasPrefix(name, "departure") {
			part2 *= myTickets[0][position]
		}
	}
	fmt.Println("part2:", part2)
}

type field struct {
	name       string
	min1, max1 int
	min2, max2 int
}

func (f field) contains(x int) bool {
	return (f.min1 <= x && x <= f.max1) || (f.min2 <= x && x <= f.max2)
}

type fieldMap map[string]field

func (f fieldMap) contains(x int) bool {
	found := false
	for _, field := range f {
		if field.contains(x) {
			found = true
		}
	}
	return found
}

type ticket []int

func parseClasses(classes string) fieldMap {
	classMap := fieldMap{}

	for _, line := range strings.Split(classes, "\n") {
		parts := strings.SplitN(line, ": ", 2)
		field := field{name: parts[0]}
		fmt.Sscanf(parts[1], "%d-%d or %d-%d", &field.min1, &field.max1, &field.min2, &field.max2)
		classMap[field.name] = field
	}

	return classMap
}

func parseTicket(s string) ticket {
	t := ticket{}
	for _, v := range strings.Split(s, ",") {
		t = append(t, util.MustAtoi(v))
	}
	return t
}

func parseTickets(s string) []ticket {
	tickets := []ticket{}
	for i, line := range strings.Split(s, "\n") {
		if i == 0 {
			continue
		}
		tickets = append(tickets, parseTicket(line))
	}
	return tickets
}

func validTickets(tickets []ticket, rules fieldMap) ([]ticket, int) {
	validTickets := []ticket{}
	errorRate := 0
	for _, ticket := range tickets {
		valid := true
		for _, x := range ticket {
			if !rules.contains(x) {
				valid = false
				errorRate += x
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, errorRate
}

func solveFieldMapping(tickets []ticket, rules fieldMap) map[string]int {
	candidates := map[string][]int{}

	for name, field := range rules {
		c := []int{}
	outer:
		for i := 0; i < len(tickets[0]); i++ {
			for j := range tickets {
				if !field.contains(tickets[j][i]) {
					continue outer
				}
			}
			c = append(c, i)
		}
		candidates[name] = c
	}

	fieldPos := map[string]int{}
	solved := make([]bool, len(tickets[0]))

	for len(fieldPos) < len(candidates) {
		for name, c := range candidates {
			unsolved := []int{}
			for _, x := range c {
				if !solved[x] {
					unsolved = append(unsolved, x)
				}
			}

			if len(unsolved) == 1 {
				fieldPos[name] = unsolved[0]
				solved[unsolved[0]] = true
			}
		}
	}

	return fieldPos
}
