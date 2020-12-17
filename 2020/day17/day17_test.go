package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	steps := 6

	cube := parseInput1(input, steps)
	for i := 0; i < steps; i++ {
		cube = step1(cube)
	}

	cubeCount := 0
	cube.Each(func(p util.Point3, x int) {
		if x > 0 {
			cubeCount++
		}
	})

	if cubeCount != 112 {
		t.Errorf("expected 112 cells, got %d", cubeCount)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	steps := 6

	hypercube := parseInput2(input, steps)
	for i := 0; i < steps; i++ {
		hypercube = step2(hypercube)
	}

	hypercubeCount := 0
	hypercube.Each(func(p util.Point4, x int) {
		if x > 0 {
			hypercubeCount++
		}
	})

	if hypercubeCount != 848 {
		t.Errorf("expected 848 cells, got %d", hypercubeCount)
	}
}
