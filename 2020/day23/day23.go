package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := "789465123"

	current, cups := stringToCup(input)
	for i := 0; i < 100; i++ {
		current = move(current, cups)
	}
	one := &cups[0]
	fmt.Println("part1:", cupToString(one)[1:])

	current, cups = stringToCup2(input)
	for i := 0; i < 10000000; i++ {
		current = move(current, cups)
	}
	one = &cups[0]
	fmt.Println("part2:", one.next.label*one.next.next.label)
}

type cup struct {
	label int
	next  *cup
}

func newCups(size int) []cup {
	cups := make([]cup, size)
	for i := 0; i < size; i++ {
		cups[i].label = i + 1
		if i == len(cups)-1 {
			cups[i].next = &cups[0]
		} else {
			cups[i].next = &cups[i+1]
		}
	}
	return cups
}

func stringToCup(s string) (*cup, []cup) {
	f := []int{}
	for _, x := range s {
		f = append(f, util.MustAtoi(string(x)))
	}

	cups := newCups(len(f))
	current := &cups[f[0]-1]

	last := current
	for i, x := range f {
		if i == 0 {
			continue
		}
		last.next = &cups[x-1]
		last = last.next
	}
	last.next = current

	return current, cups
}

func stringToCup2(s string) (*cup, []cup) {
	f := []int{}
	for _, x := range s {
		f = append(f, util.MustAtoi(string(x)))
	}

	cups := newCups(1000000)
	current := &cups[f[0]-1]

	last := current
	for i, x := range f {
		if i == 0 {
			continue
		}
		last.next = &cups[x-1]
		last = last.next
	}
	last.next = &cups[len(f)]
	cups[len(cups)-1].next = current

	return current, cups
}

func cupToString(current *cup) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(current.label))
	for i := current.next; i != current; i = i.next {
		sb.WriteString(strconv.Itoa(i.label))
	}
	return sb.String()
}

func move(current *cup, cups []cup) *cup {
	pickUp := current.next
	current.next = current.next.next.next.next
	pickUp.next.next.next = pickUp

	destLabel := current.label
	for {
		destLabel--
		if destLabel == 0 {
			destLabel = len(cups)
		}
		if pickUp.label == destLabel || pickUp.next.label == destLabel || pickUp.next.next.label == destLabel {
			continue
		}
		break
	}
	dest := &cups[destLabel-1]

	dest.next, pickUp.next.next.next = pickUp, dest.next

	return current.next
}
