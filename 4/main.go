package day4

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func GetPassport(entry string) Passport {
	var passport Passport

	fields := strings.Fields(entry)
	for _, field := range fields {
		keyValue := strings.Split(field, ":")

		switch keyValue[0] {
		case "byr":
			passport.Byr = keyValue[1]
		case "iyr":
			passport.Iyr = keyValue[1]
		case "eyr":
			passport.Eyr = keyValue[1]
		case "hgt":
			passport.Hgt = keyValue[1]
		case "hcl":
			passport.Hcl = keyValue[1]
		case "ecl":
			passport.Ecl = keyValue[1]
		case "pid":
			passport.Pid = keyValue[1]
		case "cid":
			passport.Cid = keyValue[1]
		}
	}
	return passport
}

func GetPassports() []Passport {
	batchFile := utils.GetFileContentStr("4/input.txt")
	var entry string
	var passports []Passport

	for _, line := range batchFile {
		if line == "" {
			passports = append(passports, GetPassport(entry))
			entry = ""
		}
		entry += " " + line
	}
	return passports
}

func isValidPassportPart1(passport Passport) bool {
	return !(passport.Byr == "" || passport.Ecl == "" || passport.Hcl == "" || passport.Eyr == "" ||
		passport.Hgt == "" || passport.Iyr == "" || passport.Pid == "")
}

func isValidPassportPart2(passport Passport) bool {
	byr, err := strconv.Atoi(passport.Byr)
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	iyr, err := strconv.Atoi(passport.Iyr)
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, err := strconv.Atoi(passport.Eyr)
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	if strings.HasSuffix(passport.Hgt, "in") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(passport.Hgt, "in"))
		if err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	} else if strings.HasSuffix(passport.Hgt, "cm") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(passport.Hgt, "cm"))
		if err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	} else {
		return false
	}

	matched, _ := regexp.Match(`^#[0-9a-f]{6}$`, []byte(passport.Hcl))
	if !matched {
		return false
	}

	_, found := utils.FindStringInSlice([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, passport.Ecl)
	if !found {
		return false
	}

	matched, _ = regexp.Match(`^[0-9]{9}$`, []byte(passport.Pid))
	if !matched {
		return false
	}

	return true
}

func Run() {
	var validPart1, validPart2 int
	for _, passport := range GetPassports() {
		if isValidPassportPart1(passport) {
			validPart1++
		}
		if isValidPassportPart2(passport) {
			validPart2++
		}
	}
	fmt.Printf("Day 4 part 1 answer: %d\n", validPart1)
	fmt.Printf("Day 4 part 2 answer: %d\n", validPart2)
}
