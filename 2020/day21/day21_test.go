package main

import (
	"sort"
	"strings"
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	lines := util.Lines("input-example.txt")

	recipes := parseRecipes(lines)
	possible := possibleAllergens(recipes)

	unsafe := set{}
	for _, s := range possible {
		unsafe = unsafe.union(newSet(s.values()))
	}

	got := 0
	for _, r := range recipes {
		got += len(newSet(r.ingredients).difference(unsafe))
	}

	if got != 5 {
		t.Errorf("wanted 5 safe ingredients, got %d", got)
	}
}

func TestExamplePart2(t *testing.T) {
	lines := util.Lines("input-example.txt")

	recipes := parseRecipes(lines)
	possible := possibleAllergens(recipes)
	definite := resolveAllergens(possible)

	allergens := []string{}
	for a := range definite {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)

	var sb strings.Builder
	for i, allergen := range allergens {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(definite[allergen])
	}
	got := sb.String()

	want := "mxmxvkd,sqjhc,fvjkl"
	if got != want {
		t.Errorf("wanted safe ingredients = '%s', got '%s'", want, got)
	}
}
