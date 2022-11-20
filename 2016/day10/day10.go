package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type target interface {
	Give(value int)
}

type output struct {
	id     int
	values []int
}

func (o *output) Give(value int) {
	o.values = append(o.values, value)
}

func (o output) String() string {
	return fmt.Sprintf("output %d: %v", o.id, o.values)
}

type bot struct {
	id        int
	holding   []int
	low, high target
}

func (b *bot) Give(value int) {
	b.holding = append(b.holding, value)
}

func (b *bot) Tick() bool {
	if len(b.holding) == 2 {
		if b.holding[0] < b.holding[1] {
			b.low.Give(b.holding[0])
			b.high.Give(b.holding[1])
		} else {
			b.low.Give(b.holding[1])
			b.high.Give(b.holding[0])
		}
		b.holding = []int{}
		return true
	}
	return false
}

func (b bot) String() string {
	return fmt.Sprintf("bot %d: %v", b.id, b.holding)
}

type factory struct {
	bots    map[int]*bot
	outputs map[int]*output
}

func makeFactory() factory {
	return factory{
		bots:    make(map[int]*bot),
		outputs: make(map[int]*output),
	}
}

func (f *factory) bot(id int) *bot {
	if _, ok := f.bots[id]; !ok {
		f.bots[id] = &bot{id: id}
	}
	return f.bots[id]
}

func (f *factory) output(id int) *output {
	if _, ok := f.outputs[id]; !ok {
		f.outputs[id] = &output{id: id}
	}
	return f.outputs[id]
}

func (f *factory) target(t string, id int) target {
	switch t {
	case "bot":
		return f.bot(id)
	case "output":
		return f.output(id)
	default:
		panic("unknown target type")
	}
}

func main() {
	lines := util.Lines("input.txt")

	factory := makeFactory()

	for _, line := range lines {
		f := strings.Fields(line)
		if f[0] == "value" {
			value := util.MustAtoi(f[1])
			bot := factory.bot(util.MustAtoi(f[5]))
			bot.Give(value)
		} else if f[0] == "bot" {
			bot := factory.bot(util.MustAtoi(f[1]))
			bot.low = factory.target(f[5], util.MustAtoi(f[6]))
			bot.high = factory.target(f[10], util.MustAtoi(f[11]))
		}
	}

	part1 := 0
	for {
		changed := false
		for _, bot := range factory.bots {
			if len(bot.holding) == 2 {
				if bot.holding[0] == 17 && bot.holding[1] == 61 || bot.holding[0] == 61 && bot.holding[1] == 17 {
					part1 = bot.id
				}
			}
			if bot.Tick() {
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", factory.outputs[0].values[0]*factory.outputs[1].values[0]*factory.outputs[2].values[0])
}
