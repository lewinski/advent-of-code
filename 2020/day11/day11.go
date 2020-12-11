package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	ferry := parseInput(util.Lines("input.txt"))
	for {
		gerry := ferry.iter1()
		if ferry.String() == gerry.String() {
			break
		}
		ferry = gerry
	}
	fmt.Println("part1:", ferry.occupiedSeats())

	ferry = parseInput(util.Lines("input.txt"))
	for {
		gerry := ferry.iter2()
		if ferry.String() == gerry.String() {
			break
		}
		ferry = gerry
	}
	fmt.Println("part2:", ferry.occupiedSeats())
}

const (
	floor    int = -1
	empty    int = 0
	occupied int = 1
)

type ferry struct {
	layout [][]int
}

func parseInput(lines []string) ferry {
	var ferry ferry
	ferry.layout = make([][]int, len(lines))
	for i, line := range lines {
		ferry.layout[i] = make([]int, len(line))
		for j, x := range line {
			switch x {
			case '.':
				ferry.layout[i][j] = -1
			case 'L':
				ferry.layout[i][j] = 0
			case '#':
				ferry.layout[i][j] = 1
			default:
				ferry.layout[i][j] = 3
			}
		}
	}
	return ferry
}

func (f ferry) String() string {
	var sb strings.Builder
	for _, line := range f.layout {
		for _, x := range line {
			switch x {
			case floor:
				sb.WriteString(".")
			case empty:
				sb.WriteString("L")
			case occupied:
				sb.WriteString("#")
			default:
				sb.WriteString("?")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (f ferry) iter1() ferry {
	var n ferry
	n.layout = make([][]int, len(f.layout))
	for i := range f.layout {
		n.layout[i] = make([]int, len(f.layout[i]))
		for j := range f.layout[i] {
			o := f.occupiedAround(i, j)
			if f.layout[i][j] == empty && o == 0 {
				n.layout[i][j] = occupied
			} else if f.layout[i][j] == occupied && o >= 4 {
				n.layout[i][j] = empty
			} else {
				n.layout[i][j] = f.layout[i][j]
			}
		}
	}
	return n
}

func (f ferry) iter2() ferry {
	var n ferry
	n.layout = make([][]int, len(f.layout))
	for i := range f.layout {
		n.layout[i] = make([]int, len(f.layout[i]))
		for j := range f.layout[i] {
			o := f.occupiedDirectional(i, j)
			if f.layout[i][j] == empty && o == 0 {
				n.layout[i][j] = occupied
			} else if f.layout[i][j] == occupied && o >= 5 {
				n.layout[i][j] = empty
			} else {
				n.layout[i][j] = f.layout[i][j]
			}
		}
	}
	return n
}

func (f ferry) occupiedSeats() (count int) {
	for i := range f.layout {
		for j := range f.layout[i] {
			if f.layout[i][j] == occupied {
				count++
			}
		}
	}
	return
}

func (f ferry) occupiedAround(i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if i+x >= 0 && i+x < len(f.layout) {
				if j+y >= 0 && j+y < len(f.layout[i+x]) {
					if f.layout[i+x][j+y] == occupied {
						count++
					}
				}
			}
		}
	}
	return count
}

func (f ferry) occupiedDirectional(i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for mult := 1; ; mult++ {
				ix := i + (mult * x)
				jy := j + (mult * y)
				if ix < 0 || ix >= len(f.layout) {
					break
				}
				if jy < 0 || jy >= len(f.layout[ix]) {
					break
				}
				if f.layout[ix][jy] == floor {
					continue
				}
				if f.layout[ix][jy] == occupied {
					count++
				}
				break
			}
		}
	}
	return count
}
