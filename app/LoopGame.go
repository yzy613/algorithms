package main

import "algorithms/solutions/LoopGame"

func main() {
	// 循环游戏
	err := LoopGame.Run()
	if err != nil {
		panic(err)
	}
}
