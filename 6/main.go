package day5

import (
	"adventofcode/utils"
	"fmt"
)

func Run() {
	var part1Answer, part2Answer, members int
	dict := make(map[byte]int)

	for _, line := range utils.GetFileContentByteArray("6/input.txt") {
		if len(line) != 0 {
			members++
			for _, q := range line {
				dict[q]++
			}
			continue
		}
		part1Answer += len(dict)

		for _, v := range dict {
			if v == members {
				part2Answer++
			}
		}
		dict = make(map[byte]int)
		members = 0
	}
	// Last group
	part1Answer += len(dict)
	for _, v := range dict {
		if v == members {
			part2Answer++
		}
	}

	fmt.Printf("Day 6 part 1 answer: %d\n", part1Answer)
	fmt.Printf("Day 6 part 2 answer: %d\n", part2Answer)
}
