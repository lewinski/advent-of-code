package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")[0]

	println("part1:", decompressedLength(input, 1))
	println("part2:", decompressedLength(input, 2))
}

func decompressedLength(s string, v int) int {
	length := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			end := i + 1
			for s[end] != ')' {
				end++
			}

			var len, repeat int
			fmt.Sscanf(s[i+1:end], "%dx%d", &len, &repeat)

			if v == 1 {
				length += repeat * len
			} else if v == 2 {
				length += repeat * decompressedLength(s[end+1:end+1+len], 2)
			}

			i = end + len
		} else {
			length++
		}
	}

	return length
}
