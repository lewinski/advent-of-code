package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

type state struct {
	snd       *int
	rcv       []int
	pc        int
	registers map[string]int
	running   bool
}

func (s state) get(x string) int {
	if x[0] >= 'a' && x[0] <= 'z' {
		return s.registers[x]
	} else {
		return util.MustAtoi(x)
	}
}

func (s state) set(x string, y int) {
	s.registers[x] = y
}

func newState() state {
	return state{
		registers: map[string]int{},
		running:   true,
	}
}

var ErrBlocked = errors.New("blocked")
var ErrTerminated = errors.New("terminated")

func (s *state) step(program []string, part int) (rcv *int, err error) {
	if !s.running || s.pc > len(program)-1 {
		s.running = false
		return nil, ErrTerminated
	}

	rcv = nil
	ins := program[s.pc]
	f := strings.Split(ins, " ")
	switch f[0] {
	case "snd":
		v := s.get(f[1])
		s.snd = &v
		s.pc++
	case "set":
		s.set(f[1], s.get(f[2]))
		s.pc++
	case "add":
		v := s.get(f[1]) + s.get(f[2])
		s.set(f[1], v)
		s.pc++
	case "mul":
		v := s.get(f[1]) * s.get(f[2])
		s.set(f[1], v)
		s.pc++
	case "mod":
		v := s.get(f[1]) % s.get(f[2])
		s.set(f[1], v)
		s.pc++
	case "rcv":
		if part == 1 {
			if s.get(f[1]) != 0 {
				rcv = &s.rcv[len(s.rcv)-1]
			}
		} else if part == 2 {
			if len(s.rcv) > 0 {
				rcv = &s.rcv[0]
				s.rcv = s.rcv[1:]
				s.set(f[1], *rcv)
			} else {
				return nil, ErrBlocked
			}
		}
		s.pc++
	case "jgz":
		if s.get(f[1]) > 0 {
			s.pc += s.get(f[2])
		} else {
			s.pc++
		}
	}

	return
}

func part1(program []string) int {
	computer := newState()

	for {
		rcv, err := computer.step(program, 1)
		if err != nil {
			fmt.Println("error:", err)
			return -1
		}

		if rcv != nil {
			return *rcv
		}

		if computer.snd != nil {
			computer.rcv = append(computer.rcv, *computer.snd)
			computer.snd = nil
		}
	}
}

func part2(program []string) int {
	c0 := newState()
	c0.set("p", 0)
	c1 := newState()
	c1.set("p", 1)

	c1sends := 0

	for {
		_, c0err := c0.step(program, 2)
		_, c1err := c1.step(program, 2)

		if c0err != nil && c1err != nil {
			return c1sends
		}

		if c0.snd != nil {
			c1.rcv = append(c1.rcv, *c0.snd)
			c0.snd = nil
		}

		if c1.snd != nil {
			c0.rcv = append(c0.rcv, *c1.snd)
			c1.snd = nil
			c1sends++
		}
	}
}
