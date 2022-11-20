package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
	"go.arsenm.dev/pcre"
)

func main() {
	tests := util.Lines("input.txt")

	tlsCount := 0
	sslCount := 0
	for _, test := range tests {
		if supportsTLS(test) {
			tlsCount++
		}
		if supportsSSL(test) {
			sslCount++
		}
	}
	fmt.Println("part1:", tlsCount)
	fmt.Println("part2:", sslCount)

}

func supportsTLS(ip string) bool {
	hypernet := pcre.MustCompile(`\[.*?\]`)
	hypernetMatches := hypernet.FindAllString(ip, -1)
	for _, match := range hypernetMatches {
		if hasAbba(match) {
			return false
		}
	}
	return hasAbba(ip)
}

func hasAbba(s string) bool {
	abba := pcre.MustCompile(`(.)(.)\2\1`)
	for _, match := range abba.FindAllString(s, -1) {
		if match[0] != match[1] {
			return true
		}
	}
	return false
}

func supportsSSL(ip string) bool {
	supernet := pcre.MustCompile(`\[(.*?)\]`).ReplaceAllString(ip, " ")
	hypernet := pcre.MustCompile(`\[.*?\]`)
	hypernetMatches := hypernet.FindAllString(ip, -1)
	for _, match := range hypernetMatches {
		for i := 0; i < len(match)-2; i++ {
			if match[i] == match[i+2] && match[i] != match[i+1] {
				if strings.Contains(supernet, string([]byte{match[i+1], match[i], match[i+1]})) {
					return true
				}
			}
		}
	}
	return false
}
