package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := "ljoxqyyw"
	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func prepareGrid(input string) util.IntGrid2 {
	bits := map[byte]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'a': "1010",
		'b': "1011",
		'c': "1100",
		'd': "1101",
		'e': "1110",
		'f': "1111",
	}

	g := util.IntGrid2{}

	for i := 0; i < 128; i++ {
		k := knotHash(fmt.Sprintf("%s-%d", input, i))
		var sb strings.Builder
		for j := 0; j < len(k); j++ {
			sb.WriteString(bits[k[j]])
		}
		s := sb.String()
		for j := 0; j < 128; j++ {
			if s[j] == '1' {
				g.SetCoords(j, i, 1)
			}
		}
	}

	return g
}

func part1(input string) int {
	g := prepareGrid(input)
	return len(g)
}

func part2(input string) int {
	g := prepareGrid(input)
	g.Each(func(p util.Point2, x int) {
		g.Set(p, -1)
	})

	var flood func(p util.Point2, group int)
	flood = func(p util.Point2, group int) {
		g.Set(p, group)
		for _, p2 := range p.Touching() {
			if g.Get(p2) < 0 {
				flood(p2, group)
			}
		}
	}

	group := 0
	g.Each(func(p util.Point2, x int) {
		if x > 0 {
			return
		}
		group++
		flood(p, group)
	})

	return group
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
