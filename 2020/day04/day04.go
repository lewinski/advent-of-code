package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type passport map[string]string

func main() {
	input := util.Lines("input.txt")
	passports := parseBatchFile(input)

	complete := 0
	for _, passport := range passports {
		if completePassport(passport) {
			complete++
		}
	}
	fmt.Println("part1:", complete)

	valid := 0
	for _, passport := range passports {
		if ok, _ := validPassport(passport); ok {
			valid++
		}
	}
	fmt.Println("part2:", valid)

}

func parseBatchFile(lines []string) (passports []passport) {
	lines = append(lines, "")
	passports = make([]passport, 0)

	cur := passport{}
	for _, line := range lines {
		if len(line) == 0 {
			if len(cur) > 0 {
				passports = append(passports, cur)
			}
			cur = passport{}
			continue
		}

		f := strings.Fields(line)
		for _, field := range f {
			kv := strings.SplitN(field, ":", 2)
			cur[kv[0]] = kv[1]
		}
	}

	return
}

func completePassport(p passport) bool {
	need := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range need {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
}

func validPassport(p passport) (valid bool, errors []string) {
	if !completePassport(p) {
		valid = false
		errors = append(errors, "incomplete")
		return
	}

	if match, _ := regexp.MatchString(`^\d{4}$`, p["byr"]); match {
		byr, _ := strconv.Atoi(p["byr"])
		if byr < 1920 || byr > 2002 {
			errors = append(errors, "byr range")
		}
	} else {
		errors = append(errors, "byr format")
	}

	if match, _ := regexp.MatchString(`^\d{4}$`, p["iyr"]); match {
		iyr, _ := strconv.Atoi(p["iyr"])
		if iyr < 2010 || iyr > 2020 {
			errors = append(errors, "iyr range")
		}
	} else {
		errors = append(errors, "iyr format")
	}

	if match, _ := regexp.MatchString(`^\d{4}$`, p["eyr"]); match {
		eyr, _ := strconv.Atoi(p["eyr"])
		if eyr < 2020 || eyr > 2030 {
			errors = append(errors, "eyr range")
		}
	} else {
		errors = append(errors, "eyr range")
	}

	if match, _ := regexp.MatchString(`^\d{3}cm$`, p["hgt"]); match {
		hgt, _ := strconv.Atoi(p["hgt"][0:3])
		if hgt < 150 || hgt > 193 {
			errors = append(errors, "hgt cm range")
		}
	} else if match, _ := regexp.MatchString(`^\d{2}in`, p["hgt"]); match {
		hgt, _ := strconv.Atoi(p["hgt"][0:2])
		if hgt < 59 || hgt > 76 {
			errors = append(errors, "hgt in range")
		}
	} else {
		errors = append(errors, "hgt format")
	}

	if match, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, p["hcl"]); !match {
		errors = append(errors, "hcl format")
	}

	if match, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, p["ecl"]); !match {
		errors = append(errors, "hcl format")
	}

	if match, _ := regexp.MatchString(`^\d{9}$`, p["pid"]); !match {
		errors = append(errors, "pid format")
	}

	if len(errors) == 0 {
		valid = true
	} else {
		valid = false
	}

	return
}
