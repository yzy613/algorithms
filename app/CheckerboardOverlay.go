package main

import (
	"algorithms/solutions/CheckerboardOverlay"
	"fmt"
)

func main() {
	// 棋盘覆盖
	err := CheckerboardOverlay.Run()
	if err != nil {
		fmt.Println(err)
	}
}
