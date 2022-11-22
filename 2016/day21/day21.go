package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	fmt.Println("part1:", scramble("abcdefgh", input))
	fmt.Println("part2:", descramble("fbgdceah", input))
}

func scramble(password string, ops []string) string {
	chars := []byte(password)
	for _, line := range ops {
		f := strings.Fields(line)
		if f[0] == "swap" && f[1] == "position" {
			x := util.MustAtoi(f[2])
			y := util.MustAtoi(f[5])
			chars[x], chars[y] = chars[y], chars[x]
		} else if f[0] == "swap" && f[1] == "letter" {
			x := strings.IndexByte(string(chars), f[2][0])
			y := strings.IndexByte(string(chars), f[5][0])
			chars[x], chars[y] = chars[y], chars[x]
		} else if f[0] == "rotate" && f[1] == "left" {
			x := util.MustAtoi(f[2])
			chars = append(chars[x:], chars[:x]...)
		} else if f[0] == "rotate" && f[1] == "right" {
			x := util.MustAtoi(f[2])
			chars = append(chars[len(chars)-x:], chars[:len(chars)-x]...)
		} else if f[0] == "rotate" && f[1] == "based" {
			x := strings.IndexByte(string(chars), f[6][0])
			if x >= 4 {
				x += 2
			} else {
				x += 1
			}
			x %= len(chars)
			chars = append(chars[len(chars)-x:], chars[:len(chars)-x]...)
		} else if f[0] == "reverse" {
			x := util.MustAtoi(f[2])
			y := util.MustAtoi(f[4])
			for i := 0; i < (y-x+1)/2; i++ {
				chars[x+i], chars[y-i] = chars[y-i], chars[x+i]
			}
		} else if f[0] == "move" {
			x := util.MustAtoi(f[2])
			y := util.MustAtoi(f[5])
			c := chars[x]
			chars = append(chars[:x], chars[x+1:]...)
			chars = append(chars[:y], append([]byte{c}, chars[y:]...)...)
		} else {
			panic(line)
		}
	}
	return string(chars)
}

func descramble(password string, ops []string) string {
	// for rotate based on letter
	// if letter is at before, it will get shifted right and end up at after
	// if letter is now at after, then it must have been shifted right and we can undo it by going left the same amount
	// before | shift | after
	// 0      | 1     | 1
	// 1      | 2     | 3
	// 2      | 3     | 5
	// 3      | 4     | 7
	// 4      | 6     | 2
	// 5      | 7     | 4
	// 6      | 8 (0) | 6
	// 7      | 9 (1) | 0
	shifts := map[int]int{
		1: 1,
		3: 2,
		5: 3,
		7: 4,
		2: 6,
		4: 7,
		6: 0,
		0: 1,
	}

	chars := []byte(password)
	for i := len(ops) - 1; i >= 0; i-- {
		f := strings.Fields(ops[i])
		if f[0] == "swap" && f[1] == "position" {
			// same as scramble
			x := util.MustAtoi(f[2])
			y := util.MustAtoi(f[5])
			chars[x], chars[y] = chars[y], chars[x]
		} else if f[0] == "swap" && f[1] == "letter" {
			// same as scramble
			x := strings.IndexByte(string(chars), f[2][0])
			y := strings.IndexByte(string(chars), f[5][0])
			chars[x], chars[y] = chars[y], chars[x]
		} else if f[0] == "rotate" && f[1] == "left" {
			// opposite of scramble (shifts right)
			x := util.MustAtoi(f[2])
			chars = append(chars[len(chars)-x:], chars[:len(chars)-x]...)
		} else if f[0] == "rotate" && f[1] == "right" {
			// opposite of scramble (shifts left)
			x := util.MustAtoi(f[2])
			chars = append(chars[x:], chars[:x]...)
		} else if f[0] == "rotate" && f[1] == "based" {
			// special see above
			x := strings.IndexByte(string(chars), f[6][0])
			shift := shifts[x]
			chars = append(chars[shift:], chars[:shift]...)
		} else if f[0] == "reverse" {
			// same as scramble
			x := util.MustAtoi(f[2])
			y := util.MustAtoi(f[4])
			for i := 0; i < (y-x+1)/2; i++ {
				chars[x+i], chars[y-i] = chars[y-i], chars[x+i]
			}
		} else if f[0] == "move" {
			// opposite of scramble (x, y reversed)
			y := util.MustAtoi(f[2])
			x := util.MustAtoi(f[5])
			c := chars[x]
			chars = append(chars[:x], chars[x+1:]...)
			chars = append(chars[:y], append([]byte{c}, chars[y:]...)...)
		} else {
			panic(f)
		}
	}
	return string(chars)
}
