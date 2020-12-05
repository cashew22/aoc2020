package day5

import (
	"adventofcode/utils"
	"fmt"
	"sort"
)

var rows, columns []int

func getRow(boardingPass []byte, _rows []int) int {
	if len(boardingPass) == 3 {
		return _rows[0]
	}

	if boardingPass[0] == 'B' {
		return getRow(boardingPass[1:], _rows[len(_rows)/2:])
	} else {
		return getRow(boardingPass[1:], _rows[:len(_rows)/2])
	}
}

func getColumns(boardingPass []byte, _columns []int) int {
	if len(boardingPass) == 0 {
		return _columns[0]
	}

	if boardingPass[0] == 'R' {
		return getColumns(boardingPass[1:], _columns[len(_columns)/2:])
	} else if boardingPass[0] == 'L' {
		return getColumns(boardingPass[1:], _columns[:len(_columns)/2])
	} else {
		return getColumns(boardingPass[1:], _columns)
	}
}

func getSeatId(boardingPass []byte) int {
	return getRow(boardingPass, rows)*8 + getColumns(boardingPass, columns)
}

func Run() {
	var seatIds []int
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}
	for i := 0; i < 8; i++ {
		columns = append(columns, i)
	}

	for _, boardingPass := range utils.GetFileContentByteArray("5/input.txt") {
		seatIds = append(seatIds, getSeatId(boardingPass))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(seatIds)))
	fmt.Printf("Day 5 part 1 answer: %d\n", seatIds[0])

	for i := 0; i <= len(seatIds); i++ {
		if seatIds[i]-seatIds[i+1] != 1 {
			fmt.Printf("Day 5 part 2 answer: %d\n", seatIds[i]-1)
			break
		}
	}
}
