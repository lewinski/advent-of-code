package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	records := util.Records("input-example.txt")

	tiles := parseTiles(records)
	image := solvePuzzle(tiles)
	size := len(image)

	want := 20899048083289
	got := image[0][0].id * image[0][size-1].id * image[size-1][0].id * image[size-1][size-1].id

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestExamplePart2(t *testing.T) {
	records := util.Records("input-example.txt")

	tiles := parseTiles(records)
	image := solvePuzzle(tiles)
	bitmap := assembleImage(image)
	bitmap = identifySeaMonsters(bitmap)

	roughness := roughness(bitmap)

	if roughness != 273 {
		t.Errorf("wanted roughness = 273, got %d", roughness)
	}
}
