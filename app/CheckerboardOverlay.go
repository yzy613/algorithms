package main

import (
	"algorithms/solutions/CheckerboardOverlay"
)

func main() {
	// 棋盘覆盖
	err := CheckerboardOverlay.Run()
	if err != nil {
		panic(err)
	}
}
