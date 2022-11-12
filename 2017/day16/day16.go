package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

const SIZE = 16

func main() {
	input := util.Lines("input.txt")
	instructions := strings.Split(input[0], ",")
	fmt.Println("part1:", part1(instructions))
	fmt.Println("part2:", part2(instructions))

}

func programString(programs []byte) string {
	var sb strings.Builder
	for _, x := range programs {
		sb.WriteByte(x)
	}
	return sb.String()
}

func dance(instructions []string, iter int) string {
	type instruction struct {
		op     rune
		spin   int
		x1, x2 int
		p1, p2 byte
	}
	ins := []instruction{}

	for _, i := range instructions {
		x := instruction{}
		switch i[0] {
		case 's':
			x.op = 's'
			x.spin = util.MustAtoi(i[1:])
		case 'x':
			x.op = 'x'
			f := strings.SplitN(i[1:], "/", 2)
			x.x1 = util.MustAtoi(f[0])
			x.x2 = util.MustAtoi(f[1])
		case 'p':
			x.op = 'p'
			x.p1 = i[1]
			x.p2 = i[3]
		}
		ins = append(ins, x)
	}

	seen := map[string]bool{}
	pats := map[int]string{}

	programs := make([]byte, SIZE)
	for i := 0; i < SIZE; i++ {
		programs[i] = byte(97 + i)
	}

	for i := 0; i < iter; i++ {
		for _, x := range ins {
			switch x.op {
			case 's':
				front, back := programs[0:len(programs)-x.spin], programs[len(programs)-x.spin:]
				programs = append(back, front...)
			case 'x':
				programs[x.x1], programs[x.x2] = programs[x.x2], programs[x.x1]
			case 'p':
				var s1, s2 int
				for j := 0; j < SIZE; j++ {
					if programs[j] == x.p1 {
						s1 = j
					}
					if programs[j] == x.p2 {
						s2 = j
					}
				}
				programs[s1], programs[s2] = programs[s2], programs[s1]
			}
		}

		s := programString(programs)
		pats[i] = s
		if _, ok := seen[s]; ok {
			return pats[(iter%i)-1]
		}
		seen[s] = true
	}

	return programString(programs)
}

func part1(instructions []string) string {
	return dance(instructions, 1)
}

func part2(instructions []string) string {
	return dance(instructions, 1000000000)
}
