package main

import (
	"fmt"
)

func main() {
	input := "1321131112"

	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	fmt.Println("part1:", len(input))

	for i := 0; i < 10; i++ {
		input = lookAndSay(input)
	}
	fmt.Println("part2:", len(input))
}

func lookAndSay(input string) string {
	output := make([]byte, 0, len(input))

	for i := 0; i < len(input); {
		count := 1
		for j := i + 1; j < len(input); j++ {
			if input[i] != input[j] {
				break
			}
			count++
		}
		if count > 9 {
			panic("oh no")
		}
		output = append(output, byte('0'+count), input[i])
		i += count
	}

	return string(output)
}
