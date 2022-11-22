package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

type state struct {
	pos   util.Point2
	path  string
	depth int
}

func main() {
	goal := util.Point2{3, 3}

	passcode := "rrrbmfta"

	best := math.MaxInt
	bestPath := ""

	worst := 0
	worstPath := ""

	states := []state{}
	states = append(states, state{util.Origin2(), "", 0})
	for len(states) > 0 {
		s := states[len(states)-1]
		states = states[:len(states)-1]

		if s.pos == goal {
			if s.depth < best {
				best = s.depth
				bestPath = s.path
			}
			if s.depth > worst {
				worst = s.depth
				worstPath = s.path
			}
			continue
		}

		for d, p := range moves(s.pos, passcode, s.path) {
			states = append(states, state{p, s.path + d, s.depth + 1})
		}
	}

	fmt.Println("part1:", bestPath, best)
	fmt.Println("part2:", worstPath, worst)
}

func moves(pos util.Point2, passcode, path string) map[string]util.Point2 {
	md5 := md5path(passcode, path)
	moves := map[string]util.Point2{}

	// up
	if md5[0] >= 'b' && pos[0] > 0 {
		moves["U"] = util.Point2{pos[0] - 1, pos[1]}
	}

	// down
	if md5[1] >= 'b' && pos[0] < 3 {
		moves["D"] = util.Point2{pos[0] + 1, pos[1]}
	}

	// left
	if md5[2] >= 'b' && pos[1] > 0 {
		moves["L"] = util.Point2{pos[0], pos[1] - 1}
	}

	// right
	if md5[3] >= 'b' && pos[1] < 3 {
		moves["R"] = util.Point2{pos[0], pos[1] + 1}
	}

	return moves
}

func md5path(s, t string) string {
	h := md5.New()
	io.WriteString(h, s)
	io.WriteString(h, t)

	b := []byte{}
	b = h.Sum(b)
	return hex.EncodeToString(b)
}
