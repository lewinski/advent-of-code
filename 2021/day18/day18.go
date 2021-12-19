package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

type snailfish struct {
	x, y       int
	xsub, ysub *snailfish
}

func newSnailfish(s string) snailfish {
	pos := 0
	f := parseSnailfish(s, &pos)
	f.reduce()
	return *f
}

func parseSnailfish(s string, pos *int) *snailfish {
	f := snailfish{}

	if s[*pos] != '[' {
		panic("expected [")
	}
	*pos++

	if s[*pos] == '[' {
		f.xsub = parseSnailfish(s, pos)
	} else {
		f.x = util.MustAtoi(s[*pos : *pos+1])
		*pos++
	}

	if s[*pos] != ',' {
		panic("expected ,")
	}
	*pos++

	if s[*pos] == '[' {
		f.ysub = parseSnailfish(s, pos)
	} else {
		f.y = util.MustAtoi(s[*pos : *pos+1])
		*pos++
	}

	if s[*pos] != ']' {
		panic("expected ]")
	}
	*pos++

	return &f
}

func (s snailfish) String() string {
	var x, y string
	if s.xsub != nil {
		x = s.xsub.String()
	} else {
		x = fmt.Sprintf("%d", s.x)
	}
	if s.ysub != nil {
		y = s.ysub.String()
	} else {
		y = fmt.Sprintf("%d", s.y)
	}
	return fmt.Sprintf("[%s,%s]", x, y)
}

func add(x, y snailfish) snailfish {
	s := snailfish{0, 0, &x, &y}
	s.reduce()
	return s
}

func (s *snailfish) visit(f func(s *snailfish, depth int) bool, depth int) bool {
	if s.xsub != nil && s.xsub.visit(f, depth+1) {
		return true
	}
	if s.ysub != nil && s.ysub.visit(f, depth+1) {
		return true
	}
	if f(s, depth) {
		return true
	}
	return false
}

func (s *snailfish) reduce() {
	for {
		//fmt.Println("reducing", s)
		if !s.explode() {
			if !s.split() {
				return
			}
		}
	}
}

func (s *snailfish) squashNodes() []*int {
	r := []*int{}
	if s.xsub == nil {
		r = append(r, &s.x)
	} else {
		r = append(r, s.xsub.squashNodes()...)
	}
	if s.ysub == nil {
		r = append(r, &s.y)
	} else {
		r = append(r, s.ysub.squashNodes()...)
	}
	return r
}

func (s *snailfish) explode() bool {
	var candidate *snailfish

	// find the thing we want to explode
	s.visit(func(s *snailfish, depth int) bool {
		if depth >= 4 && s.xsub == nil && s.ysub == nil {
			candidate = s
			return true
		}
		return false
	}, 0)

	if candidate == nil {
		return false
	}

	//fmt.Println("exploding", candidate)

	values := s.squashNodes()
	for i, p := range values {
		if p == &candidate.x {
			if i-1 >= 0 {
				*values[i-1] += candidate.x
			}
			if i+2 < len(values) {
				*values[i+2] += candidate.y
			}
			break
		}
	}

	s.visit(func(s *snailfish, depth int) bool {
		if s.xsub == candidate {
			s.x = 0
			s.xsub = nil
			return true
		}

		if s.ysub == candidate {
			s.y = 0
			s.ysub = nil
			return true
		}

		return false
	}, 0)

	return true
}

func (s *snailfish) split() bool {
	return s.visit(func(s *snailfish, depth int) bool {
		if s.xsub == nil && s.x >= 10 {
			//fmt.Println("splitting", s.x)
			f := snailfish{s.x / 2, s.x/2 + s.x%2, nil, nil}
			s.x = 0
			s.xsub = &f
			return true
		}
		if s.ysub == nil && s.y >= 10 {
			//fmt.Println("splitting", s.y)
			f := snailfish{s.y / 2, s.y/2 + s.y%2, nil, nil}
			s.y = 0
			s.ysub = &f
			return true
		}
		return false
	}, 0)
}

func (s snailfish) magnitude() int {
	x := s.x
	if s.xsub != nil {
		x = s.xsub.magnitude()
	}
	y := s.y
	if s.ysub != nil {
		y = s.ysub.magnitude()
	}
	return 3*x + 2*y
}

func part1(lines []string) int {
	accum := newSnailfish(lines[0])
	for _, line := range lines[1:] {
		//fmt.Println(" ", accum)

		s := newSnailfish(line)
		//fmt.Println("+", s)

		accum = add(accum, s)
		//fmt.Println("=", accum)
		//fmt.Println("")
	}
	return accum.magnitude()
}

func part2(lines []string) int {
	best := math.MinInt
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			c := add(newSnailfish(lines[i]), newSnailfish(lines[j]))
			mag := c.magnitude()
			best = util.IMax(mag, best)
		}
	}
	return best
}
