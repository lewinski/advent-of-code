package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

var UP = util.Point2{0, -1}
var DOWN = util.Point2{0, 1}
var LEFT = util.Point2{-1, 0}
var RIGHT = util.Point2{1, 0}

func main() {
	records := util.Records("input.txt")
	grid := util.IntGrid2{}

	max := util.Point2{0, 0}
	for y, line := range strings.Split(records[0], "\n") {
		for x, c := range line {
			if c == ' ' {
				continue
			}
			grid.SetCoords(x, y, int(c))
			max[0] = util.IMax(max[0], x)
			max[1] = util.IMax(max[1], y)
		}
	}

	pos := util.Point2{0, 0}
	for {
		if grid.Contains(pos) {
			break
		}
		pos[0]++
	}

	instructions := regexp.MustCompile(`(\d+|L|R)`).FindAllString(records[1], -1)

	fmt.Println("part1:", solve(instructions, pos, max, grid, wrap1))
	fmt.Println("part2:", solve(instructions, pos, max, grid, wrap2))
}

func password(pos, facing util.Point2) int {
	password := ((pos[1] + 1) * 1000) + ((pos[0] + 1) * 4)
	switch facing {
	case RIGHT:
		password += 0
	case DOWN:
		password += 1
	case LEFT:
		password += 2
	case UP:
		password += 3
	}
	return password
}

type wrapFunc func(pos, dir, max util.Point2, grid util.IntGrid2) (util.Point2, util.Point2)

func solve(instructions []string, pos, max util.Point2, grid util.IntGrid2, wrap wrapFunc) int {
	facing := util.Point2{1, 0}

	for _, i := range instructions {
		switch i {
		case "R":
			facing = util.Point2{-facing[1], facing[0]}
		case "L":
			facing = util.Point2{facing[1], -facing[0]}
		default:
			facing, pos = move(util.MustAtoi(i), pos, facing, max, grid, wrap)
		}
	}

	return password(pos, facing)
}

func move(steps int, pos, dir, max util.Point2, grid util.IntGrid2, wrap wrapFunc) (util.Point2, util.Point2) {
	for i := 0; i < steps; i++ {
		next := pos.Offset(dir)
		if grid.Get(next) == '#' {
			// hit a wall, can't walk further
			return dir, pos
		} else if !grid.Contains(next) {
			// walked outside grid, apply wrapping rule
			wdir, wpos := wrap(next, dir, max, grid)
			// wrapped into a wall, can't walk further
			if grid.Get(wpos) == '#' {
				return dir, pos
			}
			// no wall, continue walking
			dir = wdir
			pos = wpos
		} else {
			pos = next
		}
	}
	return dir, pos
}

func wrap1(pos, dir, max util.Point2, grid util.IntGrid2) (util.Point2, util.Point2) {
	if dir[0] > 0 {
		pos[0] = 0
	} else if dir[0] < 0 {
		pos[0] = max[0]
	} else if dir[1] > 0 {
		pos[1] = 0
	} else if dir[1] < 0 {
		pos[1] = max[1]
	} else {
		panic("oh no")
	}

	for {
		if grid.Contains(pos) {
			return dir, pos
		}
		pos = pos.Offset(dir)
	}
}

func wrap2(pos, dir, max util.Point2, grid util.IntGrid2) (util.Point2, util.Point2) {
	if max[0] != 149 {
		panic("not implemented for this grid")
	}

	// where did we come from
	from := pos.Offset(dir.Scale(-1))

	// helpful names
	col := from[0]
	row := from[1]

	// position in the block
	bcol := col % 50
	brow := row % 50

	var nextPos, nextDir util.Point2

	//     +-J-+-K-+
	//     |   |   |
	//     A   |   E
	//     |   |   |
	//     +---+-N-+
	//     |   |
	//     B   F
	//     |   |
	// +-I-+---+
	// |   |   |
	// C   |   G
	// |   |   |
	// +---+-M-+
	// |   |
	// D   H
	// |   |
	// +-L-+

	switch {
	case dir == LEFT && 0 <= row && row < 50 && col == 50:
		// A -> C
		nextDir = RIGHT
		nextPos = util.Point2{0, 149 - brow}
	case dir == LEFT && 50 <= row && row < 100 && col == 50:
		// B -> I
		nextDir = DOWN
		nextPos = util.Point2{brow, 100}
	case dir == LEFT && 100 <= row && row < 150 && col == 0:
		// C -> A
		nextDir = RIGHT
		nextPos = util.Point2{50, 49 - brow}
	case dir == LEFT && 150 <= row && row < 200 && col == 0:
		// D -> J
		nextDir = DOWN
		nextPos = util.Point2{50 + brow, 0}

	case dir == RIGHT && 0 <= row && row < 50 && col == 149:
		// E -> G
		nextDir = LEFT
		nextPos = util.Point2{99, 149 - brow}
	case dir == RIGHT && 50 <= row && row < 100 && col == 99:
		// F -> N
		nextDir = UP
		nextPos = util.Point2{100 + brow, 49}
	case dir == RIGHT && 100 <= row && row < 150 && col == 99:
		// G -> E
		nextDir = LEFT
		nextPos = util.Point2{149, 49 - brow}
	case dir == RIGHT && 150 <= row && row < 200 && col == 49:
		// H -> M
		nextDir = UP
		nextPos = util.Point2{50 + brow, 149}

	case dir == UP && 0 <= col && col < 50 && row == 100:
		// I -> B
		nextDir = RIGHT
		nextPos = util.Point2{50, 50 + bcol}
	case dir == UP && 50 <= col && col < 100 && row == 0:
		// J -> D
		nextDir = RIGHT
		nextPos = util.Point2{0, 150 + bcol}
	case dir == UP && 100 <= col && col < 150 && row == 0:
		// K -> L
		nextDir = UP
		nextPos = util.Point2{bcol, 199}

	case dir == DOWN && 0 <= col && col < 50 && row == 199:
		// L -> K
		nextDir = DOWN
		nextPos = util.Point2{100 + bcol, 0}
	case dir == DOWN && 50 <= col && col < 100 && row == 149:
		// M -> H
		nextDir = LEFT
		nextPos = util.Point2{49, 150 + bcol}
	case dir == DOWN && 100 <= col && col < 150 && row == 49:
		// N -> F
		nextDir = LEFT
		nextPos = util.Point2{99, 50 + bcol}

	default:
		panic("something unexpected happened")
	}

	return nextDir, nextPos
}
