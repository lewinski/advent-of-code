package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type instruction struct {
	target  string
	command string
	amount  int
	test    string
	op      string
	value   int
}

func main() {
	program := []instruction{}

	lines := util.Lines("input.txt")
	for _, line := range lines {
		fields := strings.Fields(line)
		program = append(program, instruction{
			target:  fields[0],
			command: fields[1],
			amount:  util.MustAtoi(fields[2]),
			test:    fields[4],
			op:      fields[5],
			value:   util.MustAtoi(fields[6]),
		})
	}

	fmt.Println("part1:", part1(program))
	fmt.Println("part2:", part2(program))
}

func part1(program []instruction) int {
	registers, _ := runProgram(program)

	max := math.MinInt
	for _, v := range registers {
		if v > max {
			max = v
		}
	}

	return max
}

func part2(program []instruction) int {
	_, max := runProgram(program)
	return max
}

func runProgram(program []instruction) (map[string]int, int) {
	registers := map[string]int{}
	maxValue := math.MinInt

	for _, instruction := range program {
		testValue := registers[instruction.test]
		switch instruction.op {
		case "==":
			if testValue != instruction.value {
				continue
			}
		case "!=":
			if testValue == instruction.value {
				continue
			}
		case "<":
			if testValue >= instruction.value {
				continue
			}
		case "<=":
			if testValue > instruction.value {
				continue
			}
		case ">":
			if testValue <= instruction.value {
				continue
			}
		case ">=":
			if testValue < instruction.value {
				continue
			}
		default:
			panic(instruction.op)
		}

		newValue := registers[instruction.target]
		if instruction.command == "inc" {
			newValue += instruction.amount
		} else if instruction.command == "dec" {
			newValue -= instruction.amount
		} else {
			panic(instruction.command)
		}

		if newValue > maxValue {
			maxValue = newValue
		}
		registers[instruction.target] = newValue
	}

	return registers, maxValue
}
