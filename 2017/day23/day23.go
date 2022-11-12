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

func part1(program []string) int {
	computer := newState()
	count := 0

	for {
		mul, err := computer.step(program)
		if errors.Is(err, ErrTerminated) {
			return count
		} else if err != nil {
			fmt.Println("error:", err)
			return -1
		}

		if mul {
			count++
		}
	}
}

func part2(program []string) int {
	computer := newState()
	computer.set("a", 1)
	for {
		computer.step(program)
		if computer.pc == 8 {
			break
		}
	}

	h := 0
	for b := computer.get("b"); b <= computer.get("c"); b += 17 {
		for d := 2; d < b; d++ {
			if b%d == 0 {
				h++
				break
			}
		}
	}
	return h
}

type state struct {
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

var ErrTerminated = errors.New("terminated")
var ErrInvalidInstruction = errors.New("invalid instruction")

func (s *state) step(program []string) (mul bool, err error) {
	if !s.running || s.pc > len(program)-1 {
		s.running = false
		return false, ErrTerminated
	}

	ins := program[s.pc]
	f := strings.Split(ins, " ")
	switch f[0] {
	case "set":
		s.set(f[1], s.get(f[2]))
		s.pc++
	case "sub":
		v := s.get(f[1]) - s.get(f[2])
		s.set(f[1], v)
		s.pc++
	case "mul":
		v := s.get(f[1]) * s.get(f[2])
		s.set(f[1], v)
		s.pc++
		mul = true
	case "jnz":
		v := s.get(f[1])
		if v != 0 {
			s.pc += s.get(f[2])
		} else {
			s.pc++
		}
	default:
		return false, ErrInvalidInstruction
	}

	return
}
