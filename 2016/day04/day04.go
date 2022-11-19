package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	sum := 0
	northpole := 0

	for _, t := range lines {
		name, id, checksum := parseRoom(t)

		if checksumRoom(name) == checksum {
			sum += id
			if decrypt(name, id) == "northpole object storage" {
				northpole = id
			}
		}
	}

	fmt.Println("part1:", sum)
	fmt.Println("part2:", northpole)
}

func parseRoom(room string) (name string, id int, checksum string) {
	name = room[:strings.LastIndex(room, "-")]
	id = util.MustAtoi(room[strings.LastIndex(room, "-")+1 : strings.Index(room, "[")])
	checksum = room[strings.Index(room, "[")+1 : len(room)-1]
	return
}

func checksumRoom(name string) string {
	m := map[rune]int{}
	for _, c := range name {
		if c == '-' {
			continue
		}
		m[c]++
	}
	letters := []rune{}
	for c := range m {
		letters = append(letters, c)
	}
	sort.Slice(letters, func(i, j int) bool {
		if m[letters[i]] == m[letters[j]] {
			return letters[i] < letters[j]
		}
		return m[letters[i]] > m[letters[j]]
	})
	return string(letters[:5])
}

func decrypt(name string, id int) string {
	decrypted := ""
	for _, c := range name {
		if c == '-' {
			decrypted += " "
			continue
		}
		decrypted += string(rune((int(c)-'a'+id)%26 + 'a'))
	}
	return decrypted
}
