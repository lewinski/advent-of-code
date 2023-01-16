package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
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
	fmt.Println("be patient... takes about 5 minutes...")
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

func solve(states []state, rooms map[string]*room, cb func(state)) {
	dist := distances(rooms)

	for len(states) > 0 {
		s := states[0]
		states = states[1:]

		cb(s)

		newStates := 0
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
			newStates++
		}
	}
}

func part1(rooms map[string]*room) int {
	states := []state{}
	states = append(states, state{
		pos:       "AA",
		visited:   map[string]bool{"AA": true},
		remaining: 30,
		released:  0,
	})

	best := 0

	solve(states, rooms, func(s state) {
		if s.released > best {
			best = s.released
		}
	})

	return best
}

func p2key(s state) string {
	visited := []string{}
	for k := range s.visited {
		visited = append(visited, k)
	}
	sort.StringSlice(visited).Sort()
	return fmt.Sprintf("%s", visited)
}

func part2(rooms map[string]*room) int {
	dist := distances(rooms)

	states := []state{}
	states = append(states, state{
		pos:       "AA",
		visited:   map[string]bool{"AA": true},
		remaining: 26,
		released:  0,
	})

	// collect the best we can do for each position and visited set
	best := map[string]state{}
	solve(states, rooms, func(s state) {
		key := p2key(s)

		if best[key].released < s.released {
			best[key] = s.dup()
		}
	})

	// now restart the solve from each of the best states to simulate the other actor
	best2 := 0
	for _, s := range best {
		// what is the best we could possibly do if we could simultaneously visit all the other rooms asap?
		possible := s.released
		for _, r := range rooms {
			if r.rate > 0 && !s.visited[r.name] {
				possible += r.rate * (26 - 1 - dist["AA"][r.name])
			}
		}
		// if it is worse than what we already found, skip it
		if possible < best2 {
			continue
		}

		s.pos = "AA"
		s.remaining = 26
		states := []state{s}
		solve(states, rooms, func(s state) {
			if s.released > best2 {
				best2 = s.released
			}
		})
	}

	return best2
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
				dist[i][j] = math.MaxInt32
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
