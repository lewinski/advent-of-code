package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	fmt.Println("part1:", part1(10, 2, 1000))
	fmt.Println("part2:", part2(10, 2, 21))
}

func part1(p1start, p2start, goal int) int {
	g := game{p1pos: p1start, p2pos: p2start, cur: 1}
	d := die{}

	for {
		roll := d.roll() + d.roll() + d.roll()
		g.move(roll)
		score := g.done(goal)
		if score > 0 {
			return score
		}
	}
}

type die struct {
	cur   int
	rolls int
}

func (d *die) roll() int {
	d.cur++
	if d.cur == 101 {
		d.cur = 1
	}
	d.rolls++
	return d.cur
}

type game struct {
	p1pos, p1score int
	p2pos, p2score int
	cur            int
	die            die
}

func (g *game) move(amount int) {
	roll := g.die.roll() + g.die.roll() + g.die.roll()
	if g.cur == 1 {
		g.p1pos += roll
		for g.p1pos > 10 {
			g.p1pos -= 10
		}
		g.p1score += g.p1pos
		g.cur = 2
	} else {
		g.p2pos += roll
		for g.p2pos > 10 {
			g.p2pos -= 10
		}
		g.p2score += g.p2pos
		g.cur = 1
	}
}

func (g *game) done(points int) int {
	if g.p1score >= points {
		return g.p2score * g.die.rolls
	} else if g.p2score >= points {
		return g.p1score * g.die.rolls
	}
	return 0
}

func part2(p1start, p2start, goal int) int {
	cache := make(map[[5]int][2]int)

	var quantumGame func(int, int, int, int, int) (int, int)
	quantumGame = func(p1pos, p1score, p2pos, p2score, active int) (p1wins, p2wins int) {
		key := [5]int{p1pos, p1score, p2pos, p2score, active}
		if val, found := cache[key]; found {
			return val[0], val[1]
		}

		for r1 := 1; r1 <= 3; r1++ {
			for r2 := 1; r2 <= 3; r2++ {
				for r3 := 1; r3 <= 3; r3++ {
					if active == 1 {
						newpos := p1pos + r1 + r2 + r3
						for newpos > 10 {
							newpos -= 10
						}
						newscore := p1score + newpos
						if newscore >= goal {
							p1wins++
						} else {
							subp1wins, subp2wins := quantumGame(newpos, newscore, p2pos, p2score, 2)
							p1wins += subp1wins
							p2wins += subp2wins
						}
					} else {
						newpos := p2pos + r1 + r2 + r3
						for newpos > 10 {
							newpos -= 10
						}
						newscore := p2score + newpos
						if newscore >= goal {
							p2wins++
						} else {
							subp1wins, subp2wins := quantumGame(p1pos, p1score, newpos, newscore, 1)
							p1wins += subp1wins
							p2wins += subp2wins
						}
					}
				}
			}
		}

		cache[key] = [2]int{p1wins, p2wins}
		return
	}

	p1wins, p2wins := quantumGame(p1start, 0, p2start, 0, 1)

	return util.IMax(p1wins, p2wins)
}
