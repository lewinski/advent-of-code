package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	recipes := parseRecipes(lines)
	possible := possibleAllergens(recipes)

	unsafe := set{}
	for _, s := range possible {
		unsafe = unsafe.union(newSet(s.values()))
	}

	count := 0
	for _, r := range recipes {
		count += len(newSet(r.ingredients).difference(unsafe))
	}

	fmt.Println("part1:", count)

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

	fmt.Println("part2:", sb.String())
}

type recipe struct {
	ingredients []string
	allergens   []string
}

func parseRecipes(lines []string) []recipe {
	recipes := []recipe{}
	for _, line := range lines {
		r := recipe{}
		beforeContains := true
		for _, f := range strings.Fields(line) {
			f = strings.Trim(f, "(),")
			if f == "contains" {
				beforeContains = false
			} else if beforeContains {
				r.ingredients = append(r.ingredients, f)
			} else {
				r.allergens = append(r.allergens, f)
			}
		}
		recipes = append(recipes, r)
	}
	return recipes
}

func possibleAllergens(recipes []recipe) map[string]set {
	possible := map[string]set{}

	for _, r := range recipes {
		for _, a := range r.allergens {
			// exactly 1 ingredient is the allergen
			// to be the allergen, the ingredient must show up in all of
			// the recipes that the allergen shows up in
			if _, found := possible[a]; found {
				possible[a] = possible[a].intersect(newSet(r.ingredients))
			} else {
				possible[a] = newSet(r.ingredients)
			}
		}
	}

	return possible
}

func resolveAllergens(possible map[string]set) map[string]string {
	definite := map[string]string{}
	for len(possible) > 0 {
		for a, i := range possible {
			if len(i) == 1 {
				definite[a] = i.first()
				delete(possible, a)
				for ao := range possible {
					delete(possible[ao], definite[a])
				}
				break
			}
		}
	}
	return definite
}

type set map[string]bool

func newSet(xs []string) set {
	m := set{}
	for _, x := range xs {
		m[x] = true
	}
	return m
}

func (x set) union(y set) set {
	m := set{}
	for i := range x {
		m[i] = true
	}
	for i := range y {
		m[i] = true
	}
	return m
}

func (x set) intersect(y set) set {
	m := set{}
	for i := range x {
		if y[i] {
			m[i] = true
		}
	}
	return m
}

func (x set) difference(y set) set {
	m := set{}
	for i := range x {
		if !y[i] {
			m[i] = true
		}
	}
	return m
}

func (x set) first() string {
	for i := range x {
		return i
	}
	return ""
}

func (x set) values() []string {
	a := make([]string, 0, len(x))
	for i := range x {
		a = append(a, i)
	}
	return a
}
