package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	grid := initGrid(input)
	fmt.Println("part1:", grid.flipped())
	for i := 0; i < 100; i++ {
		grid = step(grid)
	}
	fmt.Println("part2:", grid.flipped())
}

type hex struct {
	x, y, z int
}

func (h hex) around() []hex {
	return []hex{
		{x: h.x - 1, y: h.y + 1, z: h.z},
		{x: h.x + 1, y: h.y - 1, z: h.z},
		{x: h.x - 1, y: h.y, z: h.z + 1},
		{x: h.x + 1, y: h.y, z: h.z - 1},
		{x: h.x, y: h.y - 1, z: h.z + 1},
		{x: h.x, y: h.y + 1, z: h.z - 1},
	}
}

type hexGrid map[hex]bool

func (h hexGrid) flipped() int {
	flipped := 0
	for _, state := range h {
		if state {
			flipped++
		}
	}
	return flipped
}

func initGrid(input []string) hexGrid {
	tiles := hexGrid{}

	for _, line := range input {
		tiles[coords(line)] = !tiles[coords(line)]
	}

	return tiles
}

func coords(steps string) hex {
	pos := hex{}
	for i := 0; i < len(steps); i++ {
		var command string
		if steps[i] == 'n' || steps[i] == 's' {
			command = steps[i : i+2]
			i++
		} else {
			command = steps[i : i+1]
		}

		switch command {
		case "e":
			pos.x++
			pos.y--
		case "w":
			pos.x--
			pos.y++

		case "se":
			pos.y--
			pos.z++
		case "nw":
			pos.y++
			pos.z--

		case "sw":
			pos.x--
			pos.z++
		case "ne":
			pos.x++
			pos.z--
		}
	}
	return pos
}

func step(h hexGrid) hexGrid {
	hexCounts := map[hex]int{}
	for p, f := range h {
		if f == false {
			continue
		}
		for _, a := range p.around() {
			hexCounts[a]++
		}
	}

	next := hexGrid{}
	for p, c := range hexCounts {
		if h[p] == false && c == 2 {
			next[p] = true
		}
		if h[p] == true && (c == 1 || c == 2) {
			next[p] = true
		}
	}
	return next
}
