package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type sensor struct {
	position util.Point2
	beacon   util.Point2
	dist     int
}

func main() {
	lines := util.Lines("input.txt")
	sensors := parseSensors(lines)

	fmt.Println("part1:", blockedLocationsInRow(sensors, 2000000))

	distress := findEmptyLocation(sensors, 4000000)
	fmt.Println("part2:", 4000000*distress[0]+distress[1])
}

func blockedLocationsInRow(sensors []sensor, y int) int {
	minX, maxX := rowExtent(sensors, y)
	count := 0
point:
	for x := minX; x < maxX; x++ {
		for _, s := range sensors {
			dist := s.position.Distance(util.Point2{x, y})
			if dist <= s.dist {
				count++
				continue point
			}
		}
	}
	return count
}

func findEmptyLocation(sensors []sensor, maxCoord int) util.Point2 {
	for y := 0; y <= maxCoord; y++ {
		minX, maxX := rowExtent(sensors, y)
		minX = util.IMax(0, minX)
		maxX = util.IMin(maxCoord, maxX)
	point:
		for x := minX; x <= maxX; x++ {
			for _, s := range sensors {
				dist := s.position.Distance(util.Point2{x, y})
				if dist <= s.dist {
					x = util.IMax(x, s.position[0]+(s.dist-util.IAbs(s.position[1]-y)))
					continue point
				}
			}
			return util.Point2{x, y}
		}
	}
	panic("no solution")
}

func parseSensors(lines []string) []sensor {
	sensors := []sensor{}

	for _, line := range lines {
		line = strings.TrimSpace(regexp.MustCompile("[^-0-9]+").ReplaceAllString(line, " "))
		f := strings.Fields(line)

		s := sensor{
			position: util.Point2{util.MustAtoi(f[0]), util.MustAtoi(f[1])},
			beacon:   util.Point2{util.MustAtoi(f[2]), util.MustAtoi(f[3])},
		}

		s.dist = s.position.Distance(s.beacon)

		sensors = append(sensors, s)
	}

	return sensors
}

func rowExtent(sensors []sensor, y int) (int, int) {
	minX := math.MaxInt
	maxX := math.MinInt
	for _, s := range sensors {
		dx := s.dist - util.IAbs(s.position[1]-y)
		if dx <= 0 {
			continue
		}
		minX = util.IMin(s.position[0]-dx, minX)
		maxX = util.IMax(s.position[0]+dx, maxX)
	}
	return minX, maxX
}
