package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type room struct {
	name     string
	rate     int
	adjacent map[string]int
}

func main() {
	rx := regexp.MustCompile(`^Valve (..) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)

	rooms := map[string]*room{}

	lines := util.Lines("input.txt")
	for _, line := range lines {
		f := rx.FindStringSubmatch(line)
		r := room{
			name:     f[1],
			rate:     util.MustAtoi(f[2]),
			adjacent: map[string]int{},
		}
		for _, other := range strings.Split(f[3], ", ") {
			r.adjacent[other] = 1
		}
		rooms[r.name] = &r
	}

	fmt.Println("part1:", part1(rooms))
	fmt.Println("part2:", part2(rooms))
}

type state struct {
	pos       string
	visited   map[string]bool
	remaining int
	released  int
}

func (s state) dup() state {
	visited := map[string]bool{}
	for k, v := range s.visited {
		visited[k] = v
	}
	return state{
		pos:       s.pos,
		visited:   visited,
		remaining: s.remaining,
		released:  s.released,
	}
}

func part1(rooms map[string]*room) int {
	dist := distances(rooms)

	states := []state{}
	states = append(states, state{
		pos:       "AA",
		visited:   map[string]bool{"AA": true},
		remaining: 30,
		released:  0,
	})
	best := 0

	for len(states) > 0 {
		s := states[0]
		states = states[1:]

		if s.released > best {
			best = s.released
		}

		for other := range rooms {
			if rooms[other].rate == 0 {
				continue
			}

			if s.visited[other] {
				continue
			}

			next := s.dup()
			next.pos = other
			next.visited[other] = true
			next.remaining -= dist[s.pos][other] + 1
			if next.remaining < 0 {
				continue
			}
			next.released += next.remaining * rooms[other].rate
			states = append(states, next)
		}
	}

	return best
}

func part2(rooms map[string]*room) int {
	return 0
}

func distances(rooms map[string]*room) map[string]map[string]int {
	dist := map[string]map[string]int{}
	for i := range rooms {
		dist[i] = map[string]int{}
		for j := range rooms {
			if i == j {
				dist[i][j] = 0
			} else if rooms[i].adjacent[j] == 1 {
				dist[i][j] = 1
			} else {
				dist[i][j] = 100000
			}
		}
	}

	for k := range rooms {
		for i := range rooms {
			for j := range rooms {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}
