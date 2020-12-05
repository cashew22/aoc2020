package day3

import (
	"adventofcode/utils"
	"fmt"
)

var (
	forest  [][]byte
	tree    byte = '#'
	hole    byte = '.'
	row_len int
)

func slope(moves_x, moves_y int) int {
	var treeCount int
	pos_x := 0
	for row := 0; row < len(forest); row += moves_y {
		line := forest[row]
		if line[pos_x] == tree {
			treeCount++
		}

		pos_x += moves_x
		if pos_x >= row_len {
			pos_x = pos_x - row_len
		}
	}
	return treeCount
}

func Run() {
	forest = utils.GetFileContentByteArray("3/input.txt")
	row_len = len(forest[0])

	slope_1 := slope(1, 1)
	slope_2 := slope(3, 1)
	slope_3 := slope(5, 1)
	slope_4 := slope(7, 1)
	slope_5 := slope(1, 2)

	fmt.Printf("Day 3 part 1 answer: %d\n", slope_2)
	fmt.Printf("Day 3 part 2 answer: %d\n", slope_1*slope_2*slope_3*slope_4*slope_5)
}
