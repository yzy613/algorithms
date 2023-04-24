package main

import (
	"algorithms/solutions/Chessman"
	"fmt"
)

func main() {
	// 黑白棋子移动
	err := Chessman.Run()
	if err != nil {
		fmt.Println(err)
	}
}
