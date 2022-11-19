package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := "uqwqemis"
	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func hex(n int) rune {
	if n < 10 {
		return rune('0' + n)
	}
	return rune('a' + n - 10)
}

func part1(key string) string {
	password := []rune{}

	for i := 0; ; i++ {
		b := []byte{}
		h := md5.New()
		io.WriteString(h, key)
		io.WriteString(h, fmt.Sprintf("%d", i))
		b = h.Sum(b)
		if b[0] == 0 && b[1] == 0 && b[2] < 16 {
			password = append(password, hex(int(b[2])))
			if len(password) == 8 {
				break
			}
		}
	}

	return string(password)
}

func part2(key string) string {
	password := []rune("--------")

	for i := 0; ; i++ {
		b := []byte{}
		h := md5.New()
		io.WriteString(h, key)
		io.WriteString(h, fmt.Sprintf("%d", i))
		b = h.Sum(b)
		if b[0] == 0 && b[1] == 0 && b[2] < 16 {
			pos := int(b[2])
			if pos < 8 && password[pos] == '-' {
				password[pos] = hex(int(b[3] >> 4))
				if !strings.Contains(string(password), "-") {
					break
				}
			}
		}
	}

	return string(password)
}
