package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", findPacketStart(lines[0], 4))
	fmt.Println("part2:", findPacketStart(lines[0], 14))
}

func findPacketStart(packet string, markerSize int) int {
start:
	for i := 0; i < len(packet)-markerSize; i++ {
		for j := 0; j < markerSize; j++ {
			for k := j + 1; k < markerSize; k++ {
				if packet[i+j] == packet[i+k] {
					continue start
				}
			}
		}
		return i + markerSize
	}

	panic("oh no")
}
