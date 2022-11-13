package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	key := "yzbqklnj"
	fmt.Println("part1:", part1(key))
	fmt.Println("part2:", part2(key))
}

func part1(key string) int {
	for i := 1; ; i++ {
		b := []byte{}
		h := md5.New()
		io.WriteString(h, key)
		io.WriteString(h, fmt.Sprintf("%d", i))
		b = h.Sum(b)
		if b[0] == 0 && b[1] == 0 && b[2] < 16 {
			return i
		}
	}
}

func part2(key string) int {
	for i := 1; ; i++ {
		b := []byte{}
		h := md5.New()
		io.WriteString(h, key)
		io.WriteString(h, fmt.Sprintf("%d", i))
		b = h.Sum(b)
		if b[0] == 0 && b[1] == 0 && b[2] == 0 {
			return i
		}
	}
}
