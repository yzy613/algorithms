package Chessman

import "fmt"

func Run() (err error) {
	board, err := generateBoard()
	if err != nil {
		return
	}
	printBoard(board, 0)
	chess(board)
	return
}

func generateBoard() (board []byte, err error) {
	fmt.Print("enter board size n: ")
	n := 0
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}
	length := 2*n + 2
	board = make([]byte, length, length)
	i := 0
	for i < n {
		board[i] = 'o'
		board[n+i] = '*'
		i++
	}
	for i < n+2 {
		board[n+i] = '-'
		i++
	}
	return
}
