package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Records("input.txt")

	deck1 := parseDeck(input[0])
	deck2 := parseDeck(input[1])

	fmt.Println("part1:", combatGame(deck1, deck2).score())

	deck, _ := recursiveCombatGame(deck1, deck2)
	fmt.Println("part2:", deck.score())
}

type deck []int

func (d deck) score() int {
	score := 0
	for i := 0; i < len(d); i++ {
		score += d[i] * (len(d) - i)
	}
	return score
}

func (d deck) String() string {
	var sb strings.Builder
	for i, x := range d {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.Itoa(x))
	}
	return sb.String()
}

func parseDeck(input string) []int {
	deck := []int{}
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		deck = append(deck, util.MustAtoi(line))
	}
	return deck
}

func combatGame(p1, p2 deck) deck {
	for len(p1) > 0 && len(p2) > 0 {
		p1, p2 = combatTurn(p1, p2)
	}
	if len(p1) == 0 {
		return p2
	}
	return p1
}

func combatTurn(p1, p2 deck) (deck, deck) {
	c1, p1 := p1[0], p1[1:]
	c2, p2 := p2[0], p2[1:]
	if c1 > c2 {
		p1 = append(p1, c1, c2)
	} else {
		p2 = append(p2, c2, c1)
	}
	return p1, p2
}

func recursiveCombatGame(p1, p2 deck) (deck, bool) {
	state := map[string]bool{}
	for len(p1) > 0 && len(p2) > 0 {
		var infiniteLoop bool
		p1, p2, infiniteLoop = recursiveCombatTurn(p1, p2, &state)
		if infiniteLoop {
			return p1, true
		}
	}
	if len(p1) == 0 {
		return p2, false
	}
	return p1, true
}

func recursiveCombatTurn(p1, p2 deck, states *map[string]bool) (deck, deck, bool) {
	cur := p1.String() + ":" + p2.String()
	if (*states)[cur] {
		return p1, p2, true
	}
	(*states)[cur] = true

	c1, p1 := p1[0], p1[1:]
	c2, p2 := p2[0], p2[1:]

	p1win := false

	if len(p1) >= c1 && len(p2) >= c2 {
		sub1 := make([]int, c1)
		copy(sub1, p1[0:c1])
		sub2 := make([]int, c2)
		copy(sub2, p2[0:c2])
		_, p1win = recursiveCombatGame(sub1, sub2)
	} else {
		p1win = c1 > c2
	}

	if p1win {
		p1 = append(p1, c1, c2)
	} else {
		p2 = append(p2, c2, c1)
	}
	return p1, p2, false
}
