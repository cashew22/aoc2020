package day1

import (
	"adventofcode/utils"
	"fmt"
)

func Run() {
	values := utils.GetFileContentInt("1/input.txt")

	for _, value := range values {
		valToFind := 2020 - value

		for _, val := range values {
			if val == valToFind {
				fmt.Printf("Day 1 part 1 answer: %d\n", value*valToFind)
				goto part2
			}
		}
	}

part2:
	for _, first := range values {
		for _, second := range values {
			if first+second >= 2020 {
				continue
			}
			for _, third := range values {
				if first+second+third == 2020 {
					fmt.Printf("Day 1 part 2 answer: %d\n", first*second*third)
					return
				}
			}
		}
	}
}
