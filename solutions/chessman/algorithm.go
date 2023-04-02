package chessman

import (
	"fmt"
	"strconv"
)

var (
	step = 1
)

func chess(board []byte) {
	tail := len(board) - 2
	mid := 0
	for {
		mid = tail/2 - 1
		chessSwapCouple(board, mid, tail)
		tail -= 2
		if tail < 7 {
			break
		}
		chessSwapCouple(board, mid, tail)
	}
	tail++
	chessSwapCouple(board, mid, tail)
	mid -= 2
	chessSwapCouple(board, mid, tail)
	tail--
	chessSwapCouple(board, mid, tail)
	mid--
	chessSwapCouple(board, mid, tail)
}

func chessSwapCouple(board []byte, i, j int) {
	board[i], board[j] = board[j], board[i]
	board[i+1], board[j+1] = board[j+1], board[i+1]
	// print
	printBoard(board, step)
	step++
}

func printBoard(board []byte, step int) {
	fmt.Print("step " + strconv.Itoa(step) + ": ")
	for i := range board {
		fmt.Printf("%c", board[i])
	}
	fmt.Println()
}
