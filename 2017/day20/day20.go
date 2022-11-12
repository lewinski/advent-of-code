package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type particle struct {
	pos util.Point3
	vel util.Point3
	acc util.Point3
}

func parseChunk(s string) util.Point3 {
	p := util.Point3{}

	_, err := fmt.Sscanf(s[2:], "<%d,%d,%d>", &p[0], &p[1], &p[2])
	if err != nil {
		fmt.Println("error:", err)
	}
	return p
}

func mdist(p util.Point3) int {
	return util.IAbs(p[0]) + util.IAbs(p[1]) + util.IAbs(p[2])
}

func main() {
	lines := util.Lines("input.txt")

	particles := []particle{}

	for _, line := range lines {
		p := particle{}

		f := strings.Split(line, ", ")

		p.pos = parseChunk(f[0])
		p.vel = parseChunk(f[1])
		p.acc = parseChunk(f[2])

		particles = append(particles, p)
	}

	fmt.Println("part1:", part1(particles))
	fmt.Println("part2:", part2(particles))

}

func part1(particles []particle) int {
	// smallest acceleration would stay closest
	minacc := math.MaxInt
	for _, p := range particles {
		d := mdist(p.acc)
		if d <= minacc {
			minacc = d
		}
	}

	// if there is more than one smallest, then smallest started velocity would win
	minvel := math.MaxInt
	answer := 0
	for i, p := range particles {
		if mdist(p.acc) != minacc {
			continue
		}
		d := mdist(p.vel)
		if d <= minvel {
			minvel = d
			answer = i
		}
	}

	return answer
}

func part2(particles []particle) int {
	for i := 0; i < 1000; i++ {
		update := map[util.Point3][]particle{}

		// update every position
		for _, p := range particles {
			p.vel[0] += p.acc[0]
			p.vel[1] += p.acc[1]
			p.vel[2] += p.acc[2]
			p.pos[0] += p.vel[0]
			p.pos[1] += p.vel[1]
			p.pos[2] += p.vel[2]
			if _, ok := update[p.pos]; !ok {
				update[p.pos] = []particle{}
			}
			update[p.pos] = append(update[p.pos], p)
		}

		// create next list of particles
		next := []particle{}
		for _, ps := range update {
			if len(ps) == 1 {
				// didn't collide
				next = append(next, ps[0])
			}
		}

		particles = next
	}

	return len(particles)
}
