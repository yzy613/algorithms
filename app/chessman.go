package main

import "algorithms/solutions/chessman"

func main() {
	// 黑白棋子移动
	err := chessman.Run()
	if err != nil {
		panic(err)
	}
}
