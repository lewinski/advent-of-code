package main

import (
	"fmt"
)

func main() {
	fmt.Println("part1:", part1())
	fmt.Println("part2:", part2())
}

func chunk(w, x, y, z, inp, divZ, addX, addY int) (int, int, int, int) {
	w = inp      // inp w
	x *= 0       // mul x 0
	x = x + z    // add x z
	x = x % 26   // mod x 26
	z = z / divZ // div z ??
	x = x + addX // add x ??
	if x == w {  // eql x w
		x = 1
	} else {
		x = 0
	}
	if x == 0 { // eql x 0
		x = 1
	} else {
		x = 0
	}
	y *= 0       // mul y 0
	y = y + 25   // add y 25
	y = y * x    // mul y x
	y = y + 1    // add y 1
	z = z * y    // mul z y
	y *= 0       // mul y 0
	y = y + w    // add y w
	y = y + addY // add y ??
	y = y * x    // mul y x
	z = z + y    // add z y
	return w, x, y, z
}

func monad(num [14]int) int {
	w, x, y, z := 0, 0, 0, 0

	w, x, y, z = chunk(w, x, y, z, num[0], 1, 12, 4)
	w, x, y, z = chunk(w, x, y, z, num[1], 1, 11, 11)
	w, x, y, z = chunk(w, x, y, z, num[2], 1, 13, 5)
	w, x, y, z = chunk(w, x, y, z, num[3], 1, 11, 11)
	w, x, y, z = chunk(w, x, y, z, num[4], 1, 14, 14)
	w, x, y, z = chunk(w, x, y, z, num[5], 26, -10, 7)
	w, x, y, z = chunk(w, x, y, z, num[6], 1, 11, 11)
	w, x, y, z = chunk(w, x, y, z, num[7], 26, -9, 4)
	w, x, y, z = chunk(w, x, y, z, num[8], 26, -3, 6)
	w, x, y, z = chunk(w, x, y, z, num[9], 1, 13, 5)
	w, x, y, z = chunk(w, x, y, z, num[10], 26, -5, 9)
	w, x, y, z = chunk(w, x, y, z, num[11], 26, -10, 12)
	w, x, y, z = chunk(w, x, y, z, num[12], 26, -4, 14)
	w, x, y, z = chunk(w, x, y, z, num[13], 26, -5, 14)

	return z
}

func join(n [14]int) int {
	v := 0
	for _, x := range n {
		v = (v * 10) + x
	}
	return v
}

func part1() int {
	// constraints
	// d13 = d0 - 1
	// d1 = d12 - 7
	// d11 = d2 - 5
	// d3 = d8 - 8
	// d4 = d5 - 4
	// d6 = d7 - 2
	// d10 = d9

	d0, d12, d2, d8, d5, d7, d9 := 9, 9, 9, 9, 9, 9, 9
	serial := [14]int{d0, d12 - 7, d2, d8 - 8, d5 - 4, d5, d7 - 2, d7, d8, d9, d9, d2 - 5, d12, d0 - 1}
	if monad(serial) == 0 {
		return join(serial)
	} else {
		panic("whoops")
	}
}

func part2() int {
	// constraints
	// d0 = d13 + 1
	// d12 = d1 + 7
	// d2 = d11 + 5
	// d8 = d3 + 8
	// d5 = d4 + 4
	// d7 = d6 + 2
	// d10 = d9

	d13, d1, d11, d3, d4, d6, d9 := 1, 1, 1, 1, 1, 1, 1
	serial := [14]int{d13 + 1, d1, d11 + 5, d3, d4, d4 + 4, d6, d6 + 2, d3 + 8, d9, d9, d11, d1 + 7, d13}
	if monad(serial) == 0 {
		return join(serial)
	} else {
		panic("whoops")
	}
}
