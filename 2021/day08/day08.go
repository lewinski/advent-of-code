package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(lines []string) int {
	count := 0
	for _, line := range lines {
		f := strings.Fields(line)
		for i := 11; i < 15; i++ {
			switch len(f[i]) {
			case 2:
				count++
			case 3:
				count++
			case 4:
				count++
			case 7:
				count++
			}
		}
	}
	return count
}

/*

figuring out

0 abc efg 6
1   c  f  2
2 a cde g 5
3 a cd fg 5
4  bcd f  4
5 ab d fg 5
6 ab defg 6
7 a c  f  3
8 abcdefg 7
9 abcd fg 6

2 -> 1
3 -> 7
4 -> 4
7 -> 8

5 -> 2 or 3 or 5
  whichever has both 1 segments => 3
  if we know 6, then if all segments are part of 6, it is 5
    else 2

6 -> 0 or 6 or 9
  whichever has all 4 segments => 9
  if not, then whichever has both 1 segments => 0
    else 6

*/

func decode(fields []string) int {
	digits := map[string]int{}
	reverseDigits := map[int]string{}

	for len(digits) < 10 {
		for i := 0; i < 10; i++ {
			if _, found := digits[fields[i]]; found {
				continue
			}

			switch len(fields[i]) {
			case 2:
				digits[fields[i]] = 1
				reverseDigits[1] = fields[i]
			case 3:
				digits[fields[i]] = 7
				reverseDigits[7] = fields[i]
			case 4:
				digits[fields[i]] = 4
				reverseDigits[4] = fields[i]
			case 7:
				digits[fields[i]] = 8
				reverseDigits[8] = fields[i]

			case 5:
				// 2 or 3 or 5
				if one, found := reverseDigits[1]; found {
					// whichever has both 1 segments is 3
					if strings.ContainsRune(fields[i], rune(one[0])) &&
						strings.ContainsRune(fields[i], rune(one[1])) {
						digits[fields[i]] = 3
						reverseDigits[3] = fields[i]
					} else if six, found := reverseDigits[6]; found {
						// otherwise, if we know 6, then 6 contains 4
						if strings.ContainsRune(six, rune(fields[i][0])) &&
							strings.ContainsRune(six, rune(fields[i][1])) &&
							strings.ContainsRune(six, rune(fields[i][2])) &&
							strings.ContainsRune(six, rune(fields[i][3])) &&
							strings.ContainsRune(six, rune(fields[i][4])) {
							digits[fields[i]] = 5
							reverseDigits[5] = fields[i]
						} else {
							// otherwise it must be 2
							digits[fields[i]] = 2
							reverseDigits[2] = fields[i]
						}
					}
				}
			case 6:
				// 0 or 6 or 9
				if four, found := reverseDigits[4]; found {
					// if we know 4, then 9 contains 4
					if strings.ContainsRune(fields[i], rune(four[0])) &&
						strings.ContainsRune(fields[i], rune(four[1])) &&
						strings.ContainsRune(fields[i], rune(four[2])) &&
						strings.ContainsRune(fields[i], rune(four[3])) {
						digits[fields[i]] = 9
						reverseDigits[9] = fields[i]
					} else if one, found := reverseDigits[1]; found {
						// otherwise, if we know 1, then 0 contains 1
						if strings.ContainsRune(fields[i], rune(one[0])) &&
							strings.ContainsRune(fields[i], rune(one[1])) {
							digits[fields[i]] = 0
							reverseDigits[0] = fields[i]
						} else {
							// otherwise it is a 6
							digits[fields[i]] = 6
							reverseDigits[6] = fields[i]
						}
					}
				}

			}
		}
	}

	return (digits[fields[11]] * 1000) +
		(digits[fields[12]] * 100) +
		(digits[fields[13]] * 10) +
		digits[fields[14]]
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		fields := []string{}
		for _, f := range strings.Fields(line) {
			fields = append(fields, SortString(f))
		}
		total += decode(fields)
	}
	return total
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
