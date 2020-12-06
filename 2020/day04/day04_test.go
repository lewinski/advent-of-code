package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	tests := []bool{true, false, true, false}

	records := util.Records("input-example.txt")
	passports := parseBatchFile(records)
	if len(passports) != 4 {
		t.Errorf("expected 4 passports, got %d", len(passports))
	}

	for i, passport := range passports {
		if completePassport(passport) != tests[i] {
			t.Errorf("test %d: expected completePassport(%v) to be %t, got %t", i, passport, tests[i], !tests[i])
		}
	}
}

func TestCompletePassport(t *testing.T) {
	p := passport{"byr": "a", "iyr": "b", "eyr": "c", "hgt": "d", "hcl": "e", "ecl": "f", "pid": "g", "cid": "h"}
	if !completePassport(p) {
		t.Errorf("expected %v to be complete, wasn't", p)
	}

	delete(p, "cid")
	if !completePassport(p) {
		t.Errorf("expected %v to be complete, wasn't", p)
	}
}

func TestIncompletePassport(t *testing.T) {
	p := passport{"byr": "a", "iyr": "b"}
	if completePassport(p) {
		t.Errorf("expected %v to be incomplete, wasn't", p)
	}
}

func TestValidPassport(t *testing.T) {
	tests := []struct {
		want bool
		data string
	}{
		{false, "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"},
		{false, "iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946"},
		{false, "hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"},
		{false, "hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007"},
		{true, "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"},
		{true, "eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm"},
		{true, "hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022"},
		{true, "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"},
		{false, "eyr:1972"},
		{true, "byr:1972 iyr:2010 eyr:2020 hgt:190cm hcl:#aabbcc ecl:oth pid:012345678"},
		{false, "byr:x972 iyr:2010 eyr:2020 hgt:190cm hcl:#aabbcc ecl:oth pid:012345678"},
		{false, "byr:1972 iyr:20x0 eyr:2020 hgt:190cm hcl:#aabbcc ecl:oth pid:012345678"},
		{false, "byr:1972 iyr:2010 eyr:x020 hgt:190cm hcl:#aabbcc ecl:oth pid:012345678"},
		{false, "byr:1972 iyr:2010 eyr:2020 hgt:213 hcl:#aabbcc ecl:oth pid:012345678"},
		{false, "byr:1972 iyr:2010 eyr:2020 hgt:190cm hcl:#axxbcc ecl:oth pid:012345678"},
		{false, "byr:1972 iyr:2010 eyr:2020 hgt:190cm hcl:#aabbcc ecl:xxx pid:012345678"},
		{false, "byr:1972 iyr:2010 eyr:2020 hgt:190cm hcl:#aabbcc ecl:oth pid:xxxx"},
		{false, "byr:1972 iyr:2010 eyr:2020 hgt:200cm hcl:#aabbcc ecl:oth pid:012345678"},
		{false, "byr:1972 iyr:2010 eyr:2020 hgt:99in hcl:#aabbcc ecl:oth pid:012345678"},
	}

	for i, tt := range tests {
		input := []string{tt.data}
		p := parseBatchFile(input)
		got, errors := validPassport(p[0])
		if got != tt.want {
			t.Errorf("test %d: wanted validPassport(%v) to be %t, got %t (%v)", i, tt.data, tt.want, got, errors)
		}
	}
}
