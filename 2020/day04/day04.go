package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type passport map[string]string

func main() {
	records := util.Records("input.txt")
	passports := parseBatchFile(records)

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

func parseBatchFile(records []string) (passports []passport) {
	passports = make([]passport, 0, len(records))

	for _, record := range records {
		p := passport{}
		for _, field := range strings.Fields(record) {
			pair := strings.SplitN(field, ":", 2)
			p[pair[0]] = pair[1]
		}
		passports = append(passports, p)
	}

	return
}

type validatorFunc func(string) error

func validators() map[string]validatorFunc {
	return map[string]validatorFunc{
		"byr": validBirthYear,
		"iyr": validIssueYear,
		"eyr": validExpirationYear,
		"hgt": validHeight,
		"hcl": validHairColor,
		"ecl": validEyeColor,
		"pid": validPassportID,
	}
}

func completePassport(p passport) bool {
	for k := range validators() {
		if _, ok := p[k]; !ok {
			return false
		}
	}
	return true
}

func validPassport(p passport) (valid bool, errs []error) {
	for k, f := range validators() {
		if _, ok := p[k]; ok {
			if err := f(p[k]); err != nil {
				errs = append(errs, err)
			}
		} else {
			errs = append(errs, fmt.Errorf("Missing %s value", k))
		}
	}

	if len(errs) == 0 {
		valid = true
	} else {
		valid = false
	}

	return
}

func validBirthYear(x string) error {
	if match, _ := regexp.MatchString(`^\d{4}$`, x); match {
		byr, _ := strconv.Atoi(x)
		if byr < 1920 || byr > 2002 {
			return errors.New("Invalid byr value out of range")
		}
	} else {
		return errors.New("Invalid byr format")
	}
	return nil
}

func validIssueYear(x string) error {
	if match, _ := regexp.MatchString(`^\d{4}$`, x); match {
		iyr, _ := strconv.Atoi(x)
		if iyr < 2010 || iyr > 2020 {
			return errors.New("Invalid iyr value out of range")
		}
	} else {
		return errors.New("Invalid iyr format")
	}
	return nil
}

func validExpirationYear(x string) error {
	if match, _ := regexp.MatchString(`^\d{4}$`, x); match {
		eyr, _ := strconv.Atoi(x)
		if eyr < 2020 || eyr > 2030 {
			return errors.New("Invalid eyr value out of range")
		}
	} else {
		return errors.New("Invalid eyr format")
	}
	return nil
}

func validHeight(x string) error {
	if match, _ := regexp.MatchString(`^\d{3}cm$`, x); match {
		hgt, _ := strconv.Atoi(x[0:3])
		if hgt < 150 || hgt > 193 {
			return errors.New("Invalid hgt centimeter value out of range")
		}
	} else if match, _ := regexp.MatchString(`^\d{2}in`, x); match {
		hgt, _ := strconv.Atoi(x[0:2])
		if hgt < 59 || hgt > 76 {
			return errors.New("Invalid hgt inch value out of range")
		}
	} else {
		return errors.New("Invalid hgt format")
	}
	return nil
}

func validHairColor(x string) error {
	if match, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, x); !match {
		return errors.New("Invalid hcl format")
	}
	return nil
}

func validEyeColor(x string) error {
	if match, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, x); !match {
		return errors.New("Invalid ecl format")
	}
	return nil
}

func validPassportID(x string) error {
	if match, _ := regexp.MatchString(`^\d{9}$`, x); !match {
		return errors.New("Invalid pid format")
	}
	return nil
}
