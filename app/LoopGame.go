package main

import (
	"algorithms/solutions/LoopGame"
	"fmt"
)

func main() {
	// 循环游戏
	err := LoopGame.Run()
	if err != nil {
		fmt.Println(err)
	}
}
