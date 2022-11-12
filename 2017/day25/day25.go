package main

import "fmt"

func main() {
	state := 'A'
	tape := map[int]int{}
	cursor := 0

	for i := 0; i < 12481997; i++ {
		val := tape[cursor]
		switch state {
		case 'A':
			if val == 0 {
				tape[cursor] = 1
				cursor++ // right
				state = 'B'
			} else {
				tape[cursor] = 0
				cursor-- // left
				state = 'C'
			}
		case 'B':
			if val == 0 {
				tape[cursor] = 1
				cursor-- // left
				state = 'A'
			} else {
				tape[cursor] = 1
				cursor++ // right
				state = 'D'
			}
		case 'C':
			if val == 0 {
				tape[cursor] = 0
				cursor-- // left
				state = 'B'
			} else {
				tape[cursor] = 0
				cursor-- // left
				state = 'E'
			}
		case 'D':
			if val == 0 {
				tape[cursor] = 1
				cursor++ // right
				state = 'A'
			} else {
				tape[cursor] = 0
				cursor++ // right
				state = 'B'
			}
		case 'E':
			if val == 0 {
				tape[cursor] = 1
				cursor-- // left
				state = 'F'
			} else {
				tape[cursor] = 1
				cursor-- // left
				state = 'C'
			}
		case 'F':
			if val == 0 {
				tape[cursor] = 1
				cursor++ // right
				state = 'D'
			} else {
				tape[cursor] = 1
				cursor++ // right
				state = 'A'
			}
		}
	}

	total := 0
	for _, v := range tape {
		if v == 1 {
			total++
		}
	}

	fmt.Println("part1:", total)
}
