package main

import "fmt"

func main() {
	input := 3014387
	fmt.Println("part1:", game1(input))
	fmt.Println("part2:", game2(input)) // slow
}

func game1(size int) int {
	elfs := make([]int, size)
	for i := 0; i < size; i++ {
		elfs[i] = 1
	}

	pos := 0
	for {
		if elfs[pos] != 0 {
			for j := 1; j < size; j++ {
				if elfs[(pos+j)%size] != 0 {
					elfs[pos] += elfs[(pos+j)%size]
					elfs[(pos+j)%size] = 0
					break
				}
			}
		}
		if elfs[pos] == size {
			return pos + 1
		}
		pos = (pos + 1) % size
	}
}

func game2(size int) int {
	elfs := make([]int, size)
	for i := 0; i < size; i++ {
		elfs[i] = i + 1
	}

	pos := 0
	for len(elfs) > 1 {
		// find elf across
		across := (pos + len(elfs)/2) % len(elfs)

		// remove them
		elfs = append(elfs[:across], elfs[across+1:]...)
		// if elf was before us, account for removal
		if across < pos {
			pos--
		}

		// go to next elf
		pos = (pos + 1) % len(elfs)
	}
	return elfs[0]
}
