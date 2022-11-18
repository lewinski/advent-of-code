package main

import "fmt"

func main() {

	code := 20151125

	row := 1
	col := 1

	for {
		// Enter the code at row 2978, column 3083.
		if row == 2978 && col == 3083 {
			fmt.Println("part1:", code)
			break
		}

		if row == 1 {
			row = col + 1
			col = 1
		} else {
			row--
			col++
		}

		code = (code * 252533) % 33554393
	}
}
