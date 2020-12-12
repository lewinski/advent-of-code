package main

import (
	"fmt"
	"strconv"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	p1pos := p1nav(input)
	fmt.Println("part1:", abs(p1pos.x)+abs(p1pos.y))
	p2pos := p2nav(input)
	fmt.Println("part1:", abs(p2pos.x)+abs(p2pos.y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type point struct {
	x, y int
}

func p1nav(input []string) point {
	position := point{}
	facing := point{x: 1, y: 0}

	for _, line := range input {
		command := line[0:1]
		argument, _ := strconv.Atoi(line[1:])

		if command == "R" {
			command = "L"
			argument = 360 - argument
		}

		switch command {
		case "N":
			position.y += argument
		case "S":
			position.y -= argument
		case "E":
			position.x += argument
		case "W":
			position.x -= argument
		case "L":
			switch argument % 360 {
			case 90:
				facing = point{x: -facing.y, y: facing.x}
			case 180:
				facing = point{x: -facing.x, y: -facing.y}
			case 270:
				facing = point{x: facing.y, y: -facing.x}
			}
		case "F":
			position.x += argument * facing.x
			position.y += argument * facing.y
		}
	}
	return position
}

func p2nav(input []string) point {
	position := point{}
	waypoint := point{x: 10, y: 1}

	for _, line := range input {
		command := line[0:1]
		argument, _ := strconv.Atoi(line[1:])

		if command == "R" {
			command = "L"
			argument = 360 - argument
		}

		switch command {
		case "N":
			waypoint.y += argument
		case "S":
			waypoint.y -= argument
		case "E":
			waypoint.x += argument
		case "W":
			waypoint.x -= argument
		case "L":
			switch argument % 360 {
			case 90:
				waypoint = point{x: -waypoint.y, y: waypoint.x}
			case 180:
				waypoint = point{x: -waypoint.x, y: -waypoint.y}
			case 270:
				waypoint = point{x: waypoint.y, y: -waypoint.x}
			}
		case "F":
			position.x += argument * waypoint.x
			position.y += argument * waypoint.y
		}
	}
	return position
}
