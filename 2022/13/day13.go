package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	packets := []atom{}

	part1 := 0

	fields := util.Records("input.txt")
	for i, field := range fields {
		lines := strings.Split(field, "\n")
		left := parseAtom(lines[0])
		right := parseAtom(lines[1])
		packets = append(packets, left, right)

		if left.Compare(right) <= 0 {
			part1 += i + 1
		}
	}

	fmt.Println("part1:", part1)

	divider1 := listAtom(listAtom(intAtom(2)))
	divider2 := listAtom(listAtom(intAtom(6)))
	packets = append(packets, divider1, divider2)

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) < 0
	})

	part2 := 1
	for i := 0; i < len(packets); i++ {
		if divider1.Compare(packets[i]) == 0 {
			part2 *= i + 1
		} else if divider2.Compare(packets[i]) == 0 {
			part2 *= i + 1
		}
	}
	fmt.Println("part2:", part2)
}

type atom struct {
	value *int
	list  []atom
}

func intAtom(value int) atom {
	return atom{value: &value}
}

func listAtom(list ...atom) atom {
	return atom{list: list}
}

func (a atom) String() string {
	if a.value != nil {
		return fmt.Sprintf("%d", *a.value)
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i, item := range a.list {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(item.String())
	}
	sb.WriteString("]")
	return sb.String()
}

func (left atom) Compare(right atom) int {
	if left.value != nil && right.value != nil {
		if *left.value < *right.value {
			return -1
		} else if *left.value > *right.value {
			return 1
		} else {
			return 0
		}
	}

	var leftList, rightList []atom
	if left.value == nil {
		leftList = left.list
	} else {
		leftList = []atom{intAtom(*left.value)}
	}

	if right.value == nil {
		rightList = right.list
	} else {
		rightList = []atom{intAtom(*right.value)}
	}

	for i := 0; i < len(leftList); i++ {
		if i >= len(rightList) {
			return 1
		}
		cmp := leftList[i].Compare(rightList[i])
		if cmp != 0 {
			return cmp
		}
	}

	if len(leftList) < len(rightList) {
		return -1
	}

	return 0
}

func parseAtom(line string) atom {
	pos := 0
	return parseList(line, &pos)
}

func parseNumber(line string, pos *int) atom {
	end := strings.IndexAny(line[*pos:], ",]")
	value := util.MustAtoi(line[*pos : *pos+end])
	*pos += end
	return intAtom(value)
}

func parseList(line string, pos *int) atom {
	if line[*pos] != '[' {
		panic("expected '['")
	}
	*pos++

	item := atom{list: []atom{}}

	for *pos < len(line) {
		if line[*pos] == '[' {
			item.list = append(item.list, parseList(line, pos))
		} else if line[*pos] == ']' {
			*pos++
			return item
		} else if line[*pos] == ',' {
			*pos++
		} else {
			item.list = append(item.list, parseNumber(line, pos))
		}
	}

	panic("invalid data")
}
