package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	vals := make(map[int]bool)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		if _, ok := vals[val]; ok {
			panic("duplicate value")
		}
		vals[val] = true
	}

	goal := 2020

part1:
	for k := range vals {
		if _, ok := vals[goal-k]; ok {
			fmt.Println("part1:", k*(goal-k))
			break part1
		}
	}

part2:
	for k := range vals {
		for j := range vals {
			if k == j {
				continue
			}
			if _, ok := vals[goal-k-j]; ok {
				fmt.Println("part2:", k*j*(goal-k-j))
				break part2
			}
		}
	}
}
