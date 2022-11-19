package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(lines []string) string {
	message := ""
	for i := 0; i < len(lines[0]); i++ {
		letters := map[byte]int{}
		for _, line := range lines {
			letters[line[i]]++
		}
		var mostCommon byte
		for letter, count := range letters {
			if mostCommon == 0 || letters[mostCommon] < count {
				mostCommon = letter
			}
		}
		message += string(mostCommon)
	}
	return message
}

func part2(lines []string) string {
	message := ""
	for i := 0; i < len(lines[0]); i++ {
		letters := map[byte]int{}
		for _, line := range lines {
			letters[line[i]]++
		}
		var leastCommon byte
		for letter, count := range letters {
			if leastCommon == 0 || letters[leastCommon] > count {
				leastCommon = letter
			}
		}
		message += string(leastCommon)
	}
	return message
}
