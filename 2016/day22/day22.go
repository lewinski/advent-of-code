package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type node struct {
	name  string
	x, y  int
	size  int
	used  int
	avail int
	goal  bool
}

func parseNodes(df []string) []node {
	var nodes []node
	for _, line := range df {
		var n node

		f := strings.Fields(line)
		if !strings.HasPrefix(f[0], "/dev/grid") {
			continue
		}

		n.name = f[0]

		p := strings.Split(f[0], "-")
		n.x = util.MustAtoi(p[1][1:])
		n.y = util.MustAtoi(p[2][1:])

		n.size = util.MustAtoi(strings.TrimSuffix(f[1], "T"))
		n.used = util.MustAtoi(strings.TrimSuffix(f[2], "T"))
		n.avail = util.MustAtoi(strings.TrimSuffix(f[3], "T"))

		nodes = append(nodes, n)
	}
	return nodes
}

func main() {
	df := util.Lines("input.txt")

	nodes := parseNodes(df)
	fmt.Println("part1:", len(viablePairs(nodes)))

	nodes[962].goal = true
	fmt.Println("part2:", part2(nodes))
}

func viablePairs(nodes []node) [][2]int {
	var pairs [][2]int
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].used > 0 && nodes[i].used <= nodes[j].avail {
				pairs = append(pairs, [2]int{i, j})
			}
			if nodes[j].used > 0 && nodes[j].used <= nodes[i].avail {
				pairs = append(pairs, [2]int{j, i})
			}
		}
	}
	return pairs
}

func validMoves(nodes []node) [][2]int {
	moves := [][2]int{}
	pairs := viablePairs(nodes)
	for _, pair := range pairs {
		if validMove(nodes[pair[0]], nodes[pair[1]]) {
			moves = append(moves, pair)
		}
	}
	return moves
}

func validMove(from, to node) bool {
	dx := util.IAbs(from.x - to.x)
	dy := util.IAbs(from.y - to.y)
	if dx == 1 && dy == 0 || dx == 0 && dy == 1 {
		return from.used < to.avail
	}
	return false
}

func printGrid(nodes []node) {
	g := util.IntGrid2{}
	for i, n := range nodes {
		g.SetCoords(n.x, n.y, i)
	}

	for y := 0; y <= 25; y++ {
		for x := 0; x <= 37; x++ {
			idx := g.GetCoords(x, y)
			node := nodes[idx]
			if node.used == 0 { // empty
				fmt.Printf("\u001b[33m%03d:__\u001b[0m ", idx)
			} else if node.size > 100 { // never gonna move it
				fmt.Printf("\u001b[31m%03d:##\u001b[0m ", idx)
			} else if node.goal { // the one we're trying to move
				fmt.Printf("\u001b[32m%03d:%2d\u001b[0m ", idx, node.used)
			} else { // everything else
				fmt.Printf("%03d:\u001b[2m%2d\u001b[0m ", idx, node.used)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func move(nodes []node, from, to int, steps *int) {
	if !validMove(nodes[from], nodes[to]) {
		panic("invalid move")
	}

	data := nodes[from].used

	nodes[to].used += data
	nodes[to].avail -= data
	nodes[to].goal = nodes[from].goal

	nodes[from].used -= data

	nodes[from].avail += data
	nodes[from].goal = false

	*steps++
}

func part2(nodes []node) int {
	col := 26

	steps := 0

	// move over around wall
	for i := 439; i > 153; i -= col {
		move(nodes, i-col, i, &steps)
	}

	// move to top
	for i := 153; i > 130; i -= 1 {
		move(nodes, i-1, i, &steps)
	}

	// move near goal
	for i := 130; i < 936; i += col {
		move(nodes, i+col, i, &steps)
	}

	// 5 move pattern to cycle over top row to origin
	for i := 936; i > 0; i -= col {
		move(nodes, i+col, i, &steps)
		move(nodes, i+1+col, i+col, &steps)
		move(nodes, i+1, i+1+col, &steps)
		move(nodes, i+1-col, i+1, &steps)
		move(nodes, i-col, i+1-col, &steps)
	}

	// final move
	move(nodes, col, 0, &steps)

	// solved manually with this and inserting moves above
	printGrid(nodes)
	fmt.Println(validMoves(nodes))

	return steps
}
