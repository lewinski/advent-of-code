package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type unit struct {
	hp     int
	damage int
	armor  int
}

func main() {
	weapons := []item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}
	armor := []item{
		{"None", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}
	rings := []item{
		{"Right Empty", 0, 0, 0},
		{"Left Empty", 0, 0, 0},
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}

	best := math.MaxInt
	worst := 0
	boss := unit{109, 8, 2}

	for _, w := range weapons {
		for _, a := range armor {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1.name == r2.name {
						continue
					}
					player := unit{
						hp:     100,
						damage: w.damage + a.damage + r1.damage + r2.damage,
						armor:  w.armor + a.armor + r1.armor + r2.armor,
					}
					cost := w.cost + a.cost + r1.cost + r2.cost
					win := battle(player, boss)
					if win && cost < best {
						best = cost
					} else if !win && cost > worst {
						worst = cost
					}
				}
			}
		}
	}

	fmt.Println("part1:", best)
	fmt.Println("part2:", worst)
}

func battle(p1, p2 unit) bool {
	for {
		p2.hp -= util.IMax(1, p1.damage-p2.armor)
		if p2.hp <= 0 {
			return true
		}
		p1.hp -= util.IMax(1, p2.damage-p1.armor)
		if p1.hp <= 0 {
			return false
		}
	}
}
