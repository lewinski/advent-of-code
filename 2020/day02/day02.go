package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func firstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	part1Valid, part2Valid := 0, 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		policy := strings.Split(fields[0], "-")
		p1, _ := strconv.Atoi(policy[0])
		p2, _ := strconv.Atoi(policy[1])
		want := firstRune(fields[1])
		password := fields[2]

		if validatePasswordPart1(p1, p2, want, password) {
			part1Valid++
		}
		if validatePasswordPart2(p1, p2, want, password) {
			part2Valid++
		}
	}

	fmt.Println("part1:", part1Valid)
	fmt.Println("part2:", part2Valid)
}

func validatePasswordPart1(minCount, maxCount int, want rune, password string) bool {
	count := 0
	for _, c := range password {
		if c == want {
			count++
		}
	}
	return minCount <= count && count <= maxCount
}

func validatePasswordPart2(pos1, pos2 int, want rune, password string) bool {
	p1, p2 := false, false
	for i, c := range password {
		if c == want {
			if i+1 == pos1 {
				p1 = true
			}
			if i+1 == pos2 {
				p2 = true
			}
		}
	}
	return p1 != p2
}
