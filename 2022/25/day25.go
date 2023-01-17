package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	total := 0
	for _, line := range lines {
		i := snafuToInt(line)
		total += i
	}

	fmt.Println("part1:", intToSnafu(total))
}

func snafuToInt(s string) int {
	v := 0

	r := []rune(s)
	m := 1
	for i := len(r) - 1; i >= 0; i-- {
		switch r[i] {
		case '2':
			v += 2 * m
		case '1':
			v += 1 * m
		case '0':
			v += 0 * m
		case '-':
			v += -1 * m
		case '=':
			v += -2 * m
		}
		m *= 5
	}

	return v
}

func intToSnafu(i int) string {
	// convert to base 5 (digits in backwards order)
	b5 := []int{}

	m := 1
	x := i
	for {
		r := x % (m * 5)
		b5 = append(b5, r/m)
		x -= r
		m *= 5
		if x == 0 {
			break
		}
	}

	// add a leading zero
	b5 = append(b5, 0)

	// convert to balanced base 5
	for i := 0; i < len(b5); i++ {
		if b5[i] > 2 {
			b5[i] -= 5
			b5[i+1]++
		}
	}

	// remove leading zero if possible
	if b5[len(b5)-1] == 0 {
		b5 = b5[:len(b5)-1]
	}

	// convert to string
	s := ""
	for _, v := range b5 {
		switch v {
		case 2:
			s = "2" + s
		case 1:
			s = "1" + s
		case 0:
			s = "0" + s
		case -1:
			s = "-" + s
		case -2:
			s = "=" + s
		}
	}

	return s
}
