package day2

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type Policy struct {
	char byte
	min  int
	max  int
}

func isPasswordValidPart1(password string, policy Policy) bool {
	dict := utils.GetUniqueValueMapByte([]byte(password))

	if dict[policy.char] < policy.min {
		return false
	} else if dict[policy.char] > policy.max {
		return false
	}
	return true
}

func isPasswordValidPart2(password string, policy Policy) bool {
	b := []byte(password)

	if (b[policy.min-1] == policy.char) != (b[policy.max-1] == policy.char) {
		return true
	}
	return false
}

func parseLogEntry(entry string) (string, Policy) {
	var policy Policy
	fields := strings.Fields(entry)

	split := strings.Split(fields[0], "-")
	policy.min, _ = strconv.Atoi(split[0])
	policy.max, _ = strconv.Atoi(split[1])
	policy.char = fields[1][0]

	return fields[2], policy
}

func Run() {
	log := utils.GetFileContentStr("2/input.txt")

	var validp1Count, validp2Count int
	for _, entry := range log {
		password, policy := parseLogEntry(entry)
		if isPasswordValidPart1(password, policy) {
			validp1Count++
		}

		if isPasswordValidPart2(password, policy) {
			validp2Count++
		}
	}

	fmt.Printf("Day 2 part 1 answer: %d\n", validp1Count)
	fmt.Printf("Day 2 part 2 answer: %d\n", validp2Count)
}
