package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	rules := parseRules(lines)

	g := util.IntGrid2{}
	g.SetCoords(1, 0, 1)
	g.SetCoords(2, 1, 1)
	g.SetCoords(0, 2, 1)
	g.SetCoords(1, 2, 1)
	g.SetCoords(2, 2, 1)
	size := 3

	for i := 0; i < 19; i++ {
		if i == 5 {
			fmt.Println("part1:", count(g))
		}
		if i == 18 {
			fmt.Println("part2:", count(g))
			break
		}

		next := util.IntGrid2{}
		if size%2 == 0 {
			for x := 0; x < size; x += 2 {
				nx := 3 * (x / 2)
				for y := 0; y < size; y += 2 {
					ny := 3 * (y / 2)
					pat := strings.SplitAfter(rules[pat2(g, x, y)], "")
					if pat[0] == "#" {
						next.SetCoords(nx, ny, 1)
					}
					if pat[1] == "#" {
						next.SetCoords(nx+1, ny, 1)
					}
					if pat[2] == "#" {
						next.SetCoords(nx+2, ny, 1)
					}
					if pat[3] == "#" {
						next.SetCoords(nx, ny+1, 1)
					}
					if pat[4] == "#" {
						next.SetCoords(nx+1, ny+1, 1)
					}
					if pat[5] == "#" {
						next.SetCoords(nx+2, ny+1, 1)
					}
					if pat[6] == "#" {
						next.SetCoords(nx, ny+2, 1)
					}
					if pat[7] == "#" {
						next.SetCoords(nx+1, ny+2, 1)
					}
					if pat[8] == "#" {
						next.SetCoords(nx+2, ny+2, 1)
					}
				}
			}
			size = 3 * (size / 2)
		} else {
			for x := 0; x < size; x += 3 {
				nx := 4 * (x / 3)
				for y := 0; y < size; y += 3 {
					ny := 4 * (y / 3)
					pat := strings.SplitAfter(rules[pat3(g, x, y)], "")
					if pat[0] == "#" {
						next.SetCoords(nx, ny, 1)
					}
					if pat[1] == "#" {
						next.SetCoords(nx+1, ny, 1)
					}
					if pat[2] == "#" {
						next.SetCoords(nx+2, ny, 1)
					}
					if pat[3] == "#" {
						next.SetCoords(nx+3, ny, 1)
					}
					if pat[4] == "#" {
						next.SetCoords(nx, ny+1, 1)
					}
					if pat[5] == "#" {
						next.SetCoords(nx+1, ny+1, 1)
					}
					if pat[6] == "#" {
						next.SetCoords(nx+2, ny+1, 1)
					}
					if pat[7] == "#" {
						next.SetCoords(nx+3, ny+1, 1)
					}
					if pat[8] == "#" {
						next.SetCoords(nx, ny+2, 1)
					}
					if pat[9] == "#" {
						next.SetCoords(nx+1, ny+2, 1)
					}
					if pat[10] == "#" {
						next.SetCoords(nx+2, ny+2, 1)
					}
					if pat[11] == "#" {
						next.SetCoords(nx+3, ny+2, 1)
					}
					if pat[12] == "#" {
						next.SetCoords(nx, ny+3, 1)
					}
					if pat[13] == "#" {
						next.SetCoords(nx+1, ny+3, 1)
					}
					if pat[14] == "#" {
						next.SetCoords(nx+2, ny+3, 1)
					}
					if pat[15] == "#" {
						next.SetCoords(nx+3, ny+3, 1)
					}
				}
			}
			size = 4 * (size / 3)
		}
		g = next

	}
}

func parseRules(lines []string) map[string]string {
	rules := map[string]string{}

	for _, line := range lines {
		f := strings.Split(line, " => ")
		a := strings.Replace(f[0], "/", "", -1)
		b := strings.Replace(f[1], "/", "", -1)

		if len(a) == 4 {
			for _, k := range rotateFlip2(a[0:1], a[1:2], a[2:3], a[3:4]) {
				rules[k] = b
			}
		}
		if len(a) == 9 {
			for _, k := range rotateFlip3(a[0:1], a[1:2], a[2:3], a[3:4], a[4:5], a[5:6], a[6:7], a[7:8], a[8:9]) {
				rules[k] = b
			}
		}
	}

	return rules
}

func count(g util.IntGrid2) int {
	cnt := 0
	g.Each(func(p util.Point2, x int) {
		cnt += x
	})
	return cnt
}

func pat2(grid util.IntGrid2, x, y int) string {
	a, b, c, d := ".", ".", ".", "."
	if grid.GetCoords(x, y) == 1 {
		a = "#"
	}
	if grid.GetCoords(x+1, y) == 1 {
		b = "#"
	}
	if grid.GetCoords(x, y+1) == 1 {
		c = "#"
	}
	if grid.GetCoords(x+1, y+1) == 1 {
		d = "#"
	}
	return s2(a, b, c, d)
}

func pat3(grid util.IntGrid2, x, y int) string {
	a, b, c, d, e, f, g, h, i := ".", ".", ".", ".", ".", ".", ".", ".", "."
	if grid.GetCoords(x, y) == 1 {
		a = "#"
	}
	if grid.GetCoords(x+1, y) == 1 {
		b = "#"
	}
	if grid.GetCoords(x+2, y) == 1 {
		c = "#"
	}
	if grid.GetCoords(x, y+1) == 1 {
		d = "#"
	}
	if grid.GetCoords(x+1, y+1) == 1 {
		e = "#"
	}
	if grid.GetCoords(x+2, y+1) == 1 {
		f = "#"
	}
	if grid.GetCoords(x, y+2) == 1 {
		g = "#"
	}
	if grid.GetCoords(x+1, y+2) == 1 {
		h = "#"
	}
	if grid.GetCoords(x+2, y+2) == 1 {
		i = "#"
	}
	return s3(a, b, c, d, e, f, g, h, i)
}

// 2x2 rotates and flips
//
// ab  ca  dc  bd
// cd  db  ba  ac
//
// ba  db  cd  ac
// dc  ca  ab  bd
func rotateFlip2(a, b, c, d string) []string {
	return []string{
		s2(a, b, c, d),
		s2(r2(a, b, c, d)),
		s2(r2(r2(a, b, c, d))),
		s2(r2(r2(r2(a, b, c, d)))),

		s2(t2(a, b, c, d)),
		s2(t2(r2(a, b, c, d))),
		s2(t2(r2(r2(a, b, c, d)))),
		s2(t2(r2(r2(r2(a, b, c, d))))),
	}
}

func r2(a, b, c, d string) (string, string, string, string) {
	// ab  ca
	// cd  db
	return c, a, d, b
}

func t2(a, b, c, d string) (string, string, string, string) {
	// ab  ac
	// cd  bd
	return a, c, b, d
}

func s2(a, b, c, d string) string {
	return a + b + c + d
}

// 3x3 rotates and flips
//
// abc  gda  ihg  cfi
// def  heb  fed  beh
// ghi  ifc  cba  adg
//
// cba  adg  ghi  ifc
// fed  beh  def  heb
// ihg  cfi  abc  gda
func rotateFlip3(a, b, c, d, e, f, g, h, i string) []string {
	return []string{
		s3(a, b, c, d, e, f, g, h, i),
		s3(r3(a, b, c, d, e, f, g, h, i)),
		s3(r3(r3(a, b, c, d, e, f, g, h, i))),
		s3(r3(r3(r3(a, b, c, d, e, f, g, h, i)))),

		s3(t3(a, b, c, d, e, f, g, h, i)),
		s3(t3(r3(a, b, c, d, e, f, g, h, i))),
		s3(t3(r3(r3(a, b, c, d, e, f, g, h, i)))),
		s3(t3(r3(r3(r3(a, b, c, d, e, f, g, h, i))))),
	}
}

func t3(a, b, c, d, e, f, g, h, i string) (string, string, string, string, string, string, string, string, string) {
	// abc  adg
	// def  beh
	// ghi  cfi
	return a, d, g, b, e, h, c, f, i
}

func r3(a, b, c, d, e, f, g, h, i string) (string, string, string, string, string, string, string, string, string) {
	// abc  gda
	// def  heb
	// ghi  ifc
	return g, d, a, h, e, b, i, f, c
}

func s3(a, b, c, d, e, f, g, h, i string) string {
	return a + b + c + d + e + f + g + h + i
}
