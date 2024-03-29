package util

import (
	"strconv"
	"strings"
)

// MustAtoi performs strconv.Atoi without error checking.
func MustAtoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

// IntFields returns strings.Fields converted to integers
func IntFields(s string) []int {
	rv := make([]int, 0)
	for _, f := range strings.Fields(s) {
		rv = append(rv, MustAtoi(f))
	}
	return rv
}

// ISign returns the sign of the integer argument
func ISign(x int) int {
	if x == 0 {
		return 0
	} else if x < 0 {
		return -1
	}
	return 1
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

// IMin returns the smaller of the two integer arguments
func IMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
