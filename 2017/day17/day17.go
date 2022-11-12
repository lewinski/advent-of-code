package main

import "fmt"

func main() {
	step := 355
	fmt.Println("part1:", part1(step))
	fmt.Println("part2:", part2(step))

}

func part1(step int) int {
	buf := []int{0}
	pos := 0
	i := 0

	for {
		pos = (pos + step) % len(buf)

		b2 := make([]int, len(buf)+1)
		for j := 0; j < len(buf); j++ {
			if j <= pos {
				b2[j] = buf[j]
			} else {
				b2[j+1] = buf[j]
			}
		}

		pos++
		i++

		b2[pos] = i

		buf = b2

		if i == 2017 {
			return buf[pos+1]
		}
	}
}

func part2(step int) int {
	length := 1
	pos := 0
	last := 0
	for i := 1; i <= 50000000; i++ {
		pos = ((pos + step) % length) + 1
		length++
		if pos == 1 {
			last = i
		}
	}
	return last
}
