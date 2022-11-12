package main

import "fmt"

func main() {
	fmt.Println("part1:", part1(591, 393, 40000000))
	fmt.Println("part2:", part2(591, 393, 5000000))

}

func part1(a, b, n int) int {
	cnt := 0

	for i := 0; i < n; i++ {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647
		if a&0xFFFF == b&0xFFFF {
			cnt++
		}
	}

	return cnt
}

func part2(a, b, n int) int {
	cnt := 0

	for i := 0; i < n; i++ {
		for {
			a = (a * 16807) % 2147483647
			if a&0x3 == 0 {
				break
			}
		}
		for {
			b = (b * 48271) % 2147483647
			if b&0x7 == 0 {
				break
			}
		}
		if a&0xFFFF == b&0xFFFF {
			cnt++
		}
	}

	return cnt
}
