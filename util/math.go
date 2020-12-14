package util

import "strconv"

// MustAtoi performs strconv.Atoi without error checking.
func MustAtoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}
