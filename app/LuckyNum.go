package main

import (
	"algorithms/solutions/LuckyNum"
	"fmt"
)

func main() {
	// 幸运数
	err := LuckyNum.Run()
	if err != nil {
		fmt.Println(err)
	}
}
