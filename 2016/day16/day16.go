package main

import "fmt"

func main() {
	input := "00101000101111010"
	fmt.Println("part1:", checksum(dragon(input, 272)))
	fmt.Println("part2:", checksum(dragon(input, 35651584)))
}

func dragon(input string, length int) string {
	for len(input) < length {
		l := len(input)
		rev := make([]byte, l)
		for i, b := range input {
			if b == '0' {
				rev[l-i-1] = '1'
			} else {
				rev[l-i-1] = '0'
			}
		}
		input = input + "0" + string(rev)
	}
	return input[:length]
}

func checksum(input string) string {
redo:
	c := []byte{}
	for i := 0; i < len(input); i += 2 {
		if input[i] == input[i+1] {
			c = append(c, '1')
		} else {
			c = append(c, '0')
		}
	}
	input = string(c)
	if len(input)%2 == 0 {
		goto redo
	}
	return input
}
