package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	wires := newWires()
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		wires.instructions[parts[1]] = parts[0]
	}

	a := wires.eval("a")
	fmt.Println("part1:", a)

	wires.values = make(map[string]uint16)
	wires.instructions["b"] = fmt.Sprint(a)
	fmt.Println("part2:", wires.eval("a"))
}

type wires struct {
	values       map[string]uint16
	instructions map[string]string
}

func newWires() wires {
	return wires{
		values:       make(map[string]uint16),
		instructions: make(map[string]string),
	}
}

func (w *wires) eval(name string) uint16 {
	if m, _ := regexp.Match(`\d+`, []byte(name)); m {
		return uint16(util.MustAtoi(name))
	}

	if val, found := w.values[name]; found {
		return val
	}

	var val uint16

	ins, found := w.instructions[name]
	if !found {
		panic("unknown wire: " + name)
	}
	parts := strings.Split(ins, " ")

	switch len(parts) {
	case 1:
		val = w.eval(parts[0])
	case 2:
		if parts[0] == "NOT" {
			val = ^w.eval(parts[1])
		} else {
			panic("unknown op: " + parts[0])
		}
	case 3:
		switch parts[1] {
		case "AND":
			val = w.eval(parts[0]) & w.eval(parts[2])
		case "OR":
			val = w.eval(parts[0]) | w.eval(parts[2])
		case "LSHIFT":
			val = w.eval(parts[0]) << w.eval(parts[2])
		case "RSHIFT":
			val = w.eval(parts[0]) >> w.eval(parts[2])
		default:
			panic("unknown op: " + parts[1])
		}
	}

	w.values[name] = val
	return val
}
