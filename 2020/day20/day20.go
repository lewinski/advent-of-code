package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	records := util.Records("input.txt")

	tiles := parseTiles(records)
	image := solvePuzzle(tiles)
	size := len(image)

	part1 := image[0][0].id * image[0][size-1].id * image[size-1][0].id * image[size-1][size-1].id
	fmt.Println("part1:", part1)

	bitmap := assembleImage(image)
	bitmap = identifySeaMonsters(bitmap)
	roughness := roughness(bitmap)

	fmt.Println("part2:", roughness)
}

type bitmap struct {
	size int
	data [][]byte
}

func newBitmap(size int) bitmap {
	bitmap := bitmap{size: size, data: make([][]byte, size)}
	for i := 0; i < size; i++ {
		bitmap.data[i] = make([]byte, size)
	}
	return bitmap
}

func (b bitmap) rotate() bitmap {
	rotated := newBitmap(b.size)

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			rotated.data[i][j] = b.data[b.size-1-j][i]
		}
	}

	return rotated
}

func (b bitmap) flip() bitmap {
	flipped := newBitmap(b.size)

	for i := 0; i < b.size; i++ {
		flipped.data[i] = b.data[b.size-i-1]
	}

	return flipped
}

func (b bitmap) variations() []bitmap {
	return []bitmap{
		b,
		b.rotate(),
		b.rotate().rotate(),
		b.rotate().rotate().rotate(),
		b.flip(),
		b.flip().rotate(),
		b.flip().rotate().rotate(),
		b.flip().rotate().rotate().rotate(),
	}
}

func (b bitmap) String() string {
	var sb strings.Builder
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			sb.WriteByte(b.data[i][j])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func reverse10bits(x int) int {
	y := 0
	for i := 0; i < 10; i++ {
		if x&(1<<i) != 0 {
			y += 1 << (9 - i)
		}
	}
	return y
}

type tile struct {
	id     int
	top    int
	right  int
	bottom int
	left   int
	data   bitmap
}

func (t tile) rotate() tile {
	return tile{
		id:     t.id,
		top:    reverse10bits(t.left),
		right:  t.top,
		bottom: reverse10bits(t.right),
		left:   t.bottom,
		data:   t.data.rotate(),
	}
}

func (t tile) flip() tile {
	return tile{
		id:     t.id,
		top:    t.bottom,
		right:  reverse10bits(t.right),
		bottom: t.top,
		left:   reverse10bits(t.left),
		data:   t.data.flip(),
	}
}

func (t tile) variations() []tile {
	return []tile{
		t,
		t.rotate(),
		t.rotate().rotate(),
		t.rotate().rotate().rotate(),
		t.flip(),
		t.flip().rotate(),
		t.flip().rotate().rotate(),
		t.flip().rotate().rotate().rotate(),
	}
}

func parseTiles(records []string) map[int]tile {
	tiles := map[int]tile{}

	for _, record := range records {
		t := tile{}
		lines := strings.Split(record, "\n")
		fmt.Sscanf(lines[0], "Tile %d:", &t.id)

		for i := 0; i < 10; i++ {
			bit := 1 << i
			if lines[1][i] == '#' {
				t.top += bit
			}
			if lines[10][i] == '#' {
				t.bottom += bit
			}
			if lines[i+1][0] == '#' {
				t.left += bit
			}
			if lines[i+1][9] == '#' {
				t.right += bit
			}
		}

		bitmap := newBitmap(8)
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				bitmap.data[j][i] = lines[j+2][i+1]
			}
		}
		t.data = bitmap

		tiles[t.id] = t
	}

	return tiles
}

func solvePuzzle(tiles map[int]tile) [][]tile {
	size := int(math.Sqrt(float64(len(tiles))))

	image := make([][]tile, size)
	for i := 0; i < size; i++ {
		image[i] = make([]tile, size)
	}

	used := map[int]bool{}

outer:
	for _, tile := range tiles {
		for _, try := range tile.variations() {
			// try each orientation of each tile at 0,0 and recurse from there
			if tryPlacement(&image, tiles, &used, try, 0, 0) {
				break outer
			}
		}
	}

	return image
}

func tryPlacement(image *[][]tile, tiles map[int]tile, used *map[int]bool, t tile, x, y int) bool {
	// place tile
	(*image)[y][x] = t
	(*used)[t.id] = true

	// next tile to try
	x++
	if x == len(*image) {
		x = 0
		y++
		if y == len(*image) {
			return true
		}
	}

	// find constraints
	wantAbove, wantLeft := 0, 0
	if x > 0 {
		wantLeft = (*image)[y][x-1].right
	}
	if y > 0 {
		wantAbove = (*image)[y-1][x].bottom
	}

	// look at candidate tiles
	for _, next := range tiles {
		if (*used)[next.id] {
			continue
		}
		for _, try := range next.variations() {
			if (wantAbove == 0 || wantAbove == try.top) && (wantLeft == 0 || wantLeft == try.left) {
				// this tile might work, lets try it
				if tryPlacement(image, tiles, used, try, x, y) {
					return true
				}
			}
		}
	}

	// we failed, so lets back this one out
	(*image)[y][x] = tile{}
	(*used)[t.id] = false
	return false
}

func assembleImage(image [][]tile) bitmap {
	size := len(image)
	bitmap := newBitmap(size * 8)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for x := 0; x < 8; x++ {
				for y := 0; y < 8; y++ {
					bitmap.data[i*8+x][j*8+y] = image[i][j].data.data[x][y]
				}
			}
		}
	}
	return bitmap
}

func identifySeaMonsters(bitmap bitmap) bitmap {
	for _, b := range bitmap.variations() {
		found := 0
		for i := 0; i < b.size-2; i++ {
		next:
			for j := 0; j < b.size-20; j++ {
				sm := seaMonster(i, j)
				for _, p := range sm {
					if b.data[p[0]][p[1]] != '#' {
						continue next
					}
				}
				found++
				for _, p := range sm {
					b.data[p[0]][p[1]] = 'O'
				}
			}
		}
		if found > 0 {
			return b
		}
	}
	return bitmap
}

func seaMonster(i, j int) []util.Point2 {
	//  01234567890123456789
	// 0                  #
	// 1#    ##    ##    ###
	// 2 #  #  #  #  #  #
	return []util.Point2{
		{i + 0, j + 18},
		{i + 1, j + 0},
		{i + 1, j + 5},
		{i + 1, j + 6},
		{i + 1, j + 11},
		{i + 1, j + 12},
		{i + 1, j + 17},
		{i + 1, j + 18},
		{i + 1, j + 19},
		{i + 2, j + 1},
		{i + 2, j + 4},
		{i + 2, j + 7},
		{i + 2, j + 10},
		{i + 2, j + 13},
		{i + 2, j + 16},
	}
}

func roughness(bitmap bitmap) int {
	r := 0
	for _, ch := range bitmap.String() {
		if ch == '#' {
			r++
		}
	}
	return r
}
