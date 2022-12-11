package main

import (
	"fmt"
	"sort"
)

type monkeyOp func(int) int

type monkeyTest struct {
	mod, t, f int
}

type monkey struct {
	items     []int
	op        monkeyOp
	test      monkeyTest
	testCount int
}

func (m monkey) inspect(item int) int {
	if item%m.test.mod == 0 {
		return m.test.t
	}
	return m.test.f
}

func main() {
	input := monkeyExample

	fmt.Println("part1:", part1(input()))
	fmt.Println("part2:", part2(input()))
}

func part1(monkeys []monkey) int {
	manageWorry := func(i int) int {
		return i / 3
	}

	return simulate(monkeys, manageWorry, 20)
}

func part2(monkeys []monkey) int {
	group := 1
	for _, m := range monkeys {
		group *= m.test.mod
	}

	manageWorry := func(i int) int {
		return i % group
	}

	return simulate(monkeys, manageWorry, 10000)
}

func simulate(monkeys []monkey, manageWorry func(int) int, rounds int) int {
	for z := 0; z < rounds; z++ {
		for i, m := range monkeys {
			for _, j := range m.items {
				monkeys[i].testCount++

				j = manageWorry(m.op(j))
				dest := m.inspect(j)

				monkeys[dest].items = append(monkeys[dest].items, j)
			}

			monkeys[i].items = []int{}
		}
	}

	return monkeyBusiness(monkeys)
}

func monkeyBusiness(monkeys []monkey) int {
	counts := make([]int, len(monkeys))
	for i, m := range monkeys {
		counts[i] = m.testCount
	}
	sort.Ints(counts)
	return counts[len(counts)-1] * counts[len(counts)-2]
}

func monkeyExample() []monkey {
	return []monkey{
		{
			items: []int{79, 98},
			op:    func(i int) int { return i * 19 },
			test:  monkeyTest{23, 2, 3},
		},
		{
			items: []int{54, 65, 75, 74},
			op:    func(i int) int { return i + 6 },
			test:  monkeyTest{19, 2, 0},
		},
		{
			items: []int{79, 60, 97},
			op:    func(i int) int { return i * i },
			test:  monkeyTest{13, 1, 3},
		},
		{
			items: []int{74},
			op:    func(i int) int { return i + 3 },
			test:  monkeyTest{17, 0, 1},
		},
	}
}

func monkeyInput() []monkey {
	return []monkey{
		{
			// Monkey 0
			items: []int{76, 88, 96, 97, 58, 61, 67},
			op:    func(i int) int { return i * 19 },
			test:  monkeyTest{3, 2, 3},
		},
		{
			// Monkey 1
			items: []int{93, 71, 79, 83, 69, 70, 94, 98},
			op:    func(i int) int { return i + 8 },
			test:  monkeyTest{11, 5, 6},
		},
		{
			// Monkey 2
			items: []int{50, 74, 67, 92, 61, 76},
			op:    func(i int) int { return i * 13 },
			test:  monkeyTest{19, 3, 1},
		},
		{
			// Monkey 3
			items: []int{76, 92},
			op:    func(i int) int { return i + 6 },
			test:  monkeyTest{5, 1, 6},
		},
		{
			// Monkey 4
			items: []int{74, 94, 55, 87, 62},
			op:    func(i int) int { return i + 5 },
			test:  monkeyTest{2, 2, 0},
		},
		{
			// Monkey 5
			items: []int{59, 62, 53, 62},
			op:    func(i int) int { return i * i },
			test:  monkeyTest{7, 4, 7},
		},
		{
			// Monkey 6
			items: []int{62},
			op:    func(i int) int { return i + 2 },
			test:  monkeyTest{17, 5, 7},
		},
		{
			// Monkey 7
			items: []int{85, 54, 53},
			op:    func(i int) int { return i + 3 },
			test:  monkeyTest{13, 4, 0},
		},
	}
}
