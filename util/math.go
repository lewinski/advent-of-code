package util

import "strconv"

// MustAtoi performs strconv.Atoi without error checking.
func MustAtoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

// IAbs returns the absolute value of the integer argument
func IAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IMax returns the larger of the two integer arguments
func IMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
