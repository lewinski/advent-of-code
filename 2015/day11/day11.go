package main

import "fmt"

func main() {
	input := "vzbxkghb"

	next := nextPassword(input)
	fmt.Println("part1:", next)

	next = nextPassword(next)
	fmt.Println("part2:", next)
}

func nextPassword(input string) string {
	p := newPassword(input)
	for {
		p.increment(7)
		if p.valid() {
			break
		}
	}
	return string(p[:])
}

type password [8]byte

func (p password) valid() bool {
	c1 := false
	for i := 0; i < 5; i++ {
		if p[i] == p[i+1]-1 && p[i] == p[i+2]-2 {
			c1 = true
			break
		}
	}

	c2 := true
	for i := 0; i < 8; i++ {
		if p[i] == 'i' || p[i] == 'o' || p[i] == 'l' {
			c2 = false
			break
		}
	}

	c3 := false
	for i := 0; i < 7; i++ {
		if p[i] == p[i+1] {
			i += 2
			for i < 7 {
				if p[i] == p[i+1] {
					c3 = true
					break
				}
				i++
			}
		}
	}

	return c1 && c2 && c3
}

func (p *password) increment(i int) {
	if i < 0 {
		panic("oh no")
	}
	if p[i] == 'z' {
		p[i] = 'a'
		p.increment(i - 1)
	} else {
		p[i]++
	}
}

func newPassword(input string) password {
	p := password{}
	for i := 0; i < 8; i++ {
		p[i] = input[i]
	}
	return p
}
