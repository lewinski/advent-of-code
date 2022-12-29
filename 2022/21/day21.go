package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	monkeys := map[string]string{}

	for _, line := range lines {
		f := strings.SplitN(line, ": ", 2)
		monkeys[f[0]] = f[1]
	}

	fmt.Println("part1:", int64(eval(monkeys, "root")))

	// search for value where difference becomes zero
	monkeys["root"] = strings.Replace(monkeys["root"], "+", "-", 1)

	min, max := 1, 1
	minSign := math.Signbit(evalHuman(monkeys, "root", min))
	for {
		if math.Signbit(evalHuman(monkeys, "root", max)) != minSign {
			break
		}
		max *= 2
	}

	for {
		if min == max {
			panic("oh no")
		}

		half := min + ((max - min) / 2)
		diff := evalHuman(monkeys, "root", half)

		if diff == 0 {
			fmt.Println("part2:", half)
			break
		}

		if math.Signbit(diff) != minSign {
			max = half
		} else {
			min = half
		}
	}
}

func eval(monkeys map[string]string, name string) float64 {
	job, found := monkeys[name]
	if !found {
		panic("not found: " + name)
	}

	f := strings.Split(job, " ")
	if len(f) == 1 {
		return float64(util.MustAtoi(job))
	}

	l := eval(monkeys, f[0])
	r := eval(monkeys, f[2])

	switch f[1] {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	default:
		panic("unknown op: " + f[1])
	}
}

func evalHuman(monkeys map[string]string, name string, human int) float64 {
	tmp := monkeys["humn"]
	monkeys["humn"] = fmt.Sprintf("%d", human)
	rv := eval(monkeys, name)
	monkeys["humn"] = tmp
	return rv
}
