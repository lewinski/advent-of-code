package main

import (
	"fmt"
	"math"
	"path"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	// generate lists of fully-qualified files and directories
	files := map[string]int{}
	dirs := map[string]int{}

	var cwd string
	for _, line := range lines {
		f := strings.Fields(line)
		if f[0] == "$" && f[1] == "cd" {
			if strings.HasPrefix(f[2], "/") {
				cwd = f[2]
			} else {
				cwd = path.Join(cwd, f[2])
			}
			dirs[cwd] = 0
		} else if f[0] == "$" && f[1] == "ls" {
			continue
		} else if f[0] == "dir" {
			dirs[path.Join(cwd, f[1])] = 0
		} else {
			files[path.Join(cwd, f[1])] = util.MustAtoi(f[0])
		}
	}

	// calculate total size of each directory
	for d := range dirs {
		for f, s := range files {
			if d == "/" || strings.HasPrefix(f, d+"/") {
				dirs[d] += s
			}
		}
	}

	// solve part 1
	part1 := 0
	for _, s := range dirs {
		if s < 100000 {
			part1 += s
		}
	}
	fmt.Println("part1:", part1)

	// solve part 2
	free := 70000000 - dirs["/"]
	required := 30000000 - free

	part2 := math.MaxInt
	for _, s := range dirs {
		if s > required && s < part2 {
			part2 = s
		}
	}

	fmt.Println("part2:", part2)
}
