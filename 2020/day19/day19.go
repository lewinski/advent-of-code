package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Records("input.txt")

	rules := parseRules(input[0])
	lines := strings.Split(input[1], "\n")

	matches := countPart1Matches(lines, rules)
	fmt.Println("part1:", matches)

	matches = countPart2Matches(lines, rules)
	fmt.Println("part2:", matches)
}

func countPart1Matches(lines []string, rules map[string]string) int {
	matches := 0
	regex := regexp.MustCompile("^" + resolveRule("0", rules) + "$")
	for _, line := range lines {
		if regex.MatchString(line) {
			matches++
		}
	}
	return matches
}

func countPart2Matches(lines []string, rules map[string]string) int {
	// rules["0"] = "8 | 11"
	// rules["8"] = "42 | 42 8"
	// rules["11"] = "42 31 | 42 11 31"
	//   =>
	// rules["0"] = "42+ 42{n} 31{n}"

	head := resolveRule("42", rules)
	headRegex := regexp.MustCompile(head)

	tail := resolveRule("31", rules)
	tailRegex := regexp.MustCompile(tail)

	matches := 0
	regex := regexp.MustCompile("^(?P<head>" + head + "+)(?P<tail>" + tail + "+)$")
	for _, line := range lines {
		if m := regex.FindStringSubmatch(line); m != nil {
			heads := len(headRegex.ReplaceAllString(m[regex.SubexpIndex("head")], "x"))
			tails := len(tailRegex.ReplaceAllString(m[regex.SubexpIndex("tail")], "x"))
			if heads > tails {
				matches++
			}
		}
	}

	return matches
}

func parseRules(rules string) map[string]string {
	r := map[string]string{}

	for _, line := range strings.Split(rules, "\n") {
		fields := strings.SplitN(line, ": ", 2)
		r[fields[0]] = fields[1]
	}

	return r
}

func resolveRule(id string, rules map[string]string) string {
	if strings.HasPrefix(rules[id], "\"") {
		return strings.Trim(rules[id], "\"")
	}

	fields := strings.Fields(rules[id])

	wrap := false
	var sb strings.Builder
	for _, x := range fields {
		if x == "|" {
			sb.WriteString("|")
			wrap = true
		} else {
			sb.WriteString(resolveRule(x, rules))
		}
	}

	rv := sb.String()
	if wrap {
		rv = "(" + rv + ")"
	}
	return rv
}
