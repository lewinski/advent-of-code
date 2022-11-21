package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func main() {
	salt := "cuanljph"

	fmt.Println("part1:", part1(salt))
	fmt.Println("part2:", part2(salt))
}

func part1(salt string) int {
	size := 50000
	hashes := make([]string, 0, size)
	for i := 0; i < size; i++ {
		hashes = append(hashes, md5salt(salt, i))
	}
	return findKeys(hashes, 64)
}

func part2(salt string) int {
	size := 30000
	hashes := make([]string, 0, size)
	for i := 0; i < size; i++ {
		hashes = append(hashes, stretch(md5salt(salt, i)))
	}
	return findKeys(hashes, 64)
}

func findKeys(hashes []string, goal int) int {
	keysFound := 0
outer:
	for i, h := range hashes {
		var c byte
		for j := 0; j < len(h)-2; j++ {
			if h[j] == h[j+1] && h[j] == h[j+2] {
				c = h[j]
				break
			}
		}

		if c == 0 {
			continue
		}

		five := strings.Repeat(string(c), 5)
		for j := i + 1; j < i+1001; j++ {
			if strings.Contains(hashes[j], five) {
				keysFound++
				if keysFound == goal {
					return i
				}
				continue outer
			}
		}
	}
	panic("oh no")
}

func stretch(s string) string {
	for i := 0; i < 2016; i++ {
		h := md5.Sum([]byte(s))
		s = hex.EncodeToString(h[:])
	}
	return s
}

func md5salt(s string, i int) string {
	h := md5.New()
	io.WriteString(h, s)
	io.WriteString(h, fmt.Sprintf("%d", i))

	b := []byte{}
	b = h.Sum(b)
	return hex.EncodeToString(b)
}
