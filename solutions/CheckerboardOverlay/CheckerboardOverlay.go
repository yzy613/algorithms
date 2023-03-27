package CheckerboardOverlay

import (
	"fmt"
	"strconv"
)

func Run() (err error) {
	checkerboard, size, row, col, err := enterArgs()
	if err != nil {
		return
	}
	overlay(checkerboard, size, 0, 0, row, col)
	printCheckerboard(checkerboard)
	return
}

func enterArgs() (checkerboard [][]int, size, row, col int, err error) {
	k := 0
	fmt.Print("enter k: ")
	_, err = fmt.Scan(&k)
	if err != nil {
		return
	}
	size = 1
	for i := 0; i < k; i++ {
		size *= 2
	}
	checkerboard = make([][]int, size, size)
	for i := 0; i < size; i++ {
		checkerboard[i] = make([]int, size, size)
	}
	fmt.Print("enter special block position row and column: ")
	_, err = fmt.Scan(&row, &col)
	if err != nil {
		return
	}
	checkerboard[row][col] = 1
	return
}

func printCheckerboard(checkerboard [][]int) {
	maxInterval := 0
	currLayer := layer
	for currLayer > 0 {
		currLayer /= 10
		maxInterval++
	}
	for _, row := range checkerboard {
		for _, col := range row {
			colLength := len(strconv.Itoa(col))
			for i := 0; i < maxInterval-colLength; i++ {
				fmt.Print(" ")
			}
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
