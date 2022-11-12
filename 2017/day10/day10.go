package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(256, lines[0]))
	fmt.Println("part2:", part2(lines[0]))
}

func part1(size int, input string) int {
	lengths := []int{}
	for _, x := range strings.Split(input, ",") {
		lengths = append(lengths, util.MustAtoi(x))
	}

	knot := newKnotState(size)
	for _, x := range lengths {
		knot.twist(x)

	}
	return knot.ring[0] * knot.ring[1]
}

func part2(input string) string {
	return knotHash(input)
}

func knotHash(input string) string {
	lengths := []int{}
	for i := 0; i < len(input); i++ {
		lengths = append(lengths, int(input[i]))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	knot := newKnotState(256)
	for i := 0; i < 64; i++ {
		for _, x := range lengths {
			knot.twist(x)
		}
	}

	return knot.String()
}

type knotState struct {
	ring []int
	size int
	pos  int
	skip int
}

func newKnotState(size int) knotState {
	ks := knotState{
		size: size,
		pos:  0,
		skip: 0,
	}
	ks.ring = make([]int, size)
	for i := 0; i < size; i++ {
		ks.ring[i] = i
	}
	return ks
}

func (ks *knotState) twist(length int) {
	for i := 0; i < length/2; i++ {
		p1 := (ks.pos + i) % ks.size
		p2 := (ks.pos + length - i - 1) % ks.size
		ks.ring[p1], ks.ring[p2] = ks.ring[p2], ks.ring[p1]
	}

	ks.pos += length + ks.skip
	ks.skip++
}

func (ks knotState) String() string {
	var sb strings.Builder

	for i := 0; i < ks.size; i += 16 {
		x := 0
		for j := 0; j < 16; j++ {
			x ^= ks.ring[i+j]
		}
		sb.WriteString(fmt.Sprintf("%02x", x))
	}

	return sb.String()
}
