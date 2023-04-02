package main

import "algorithms/solutions/LuckyNum"

func main() {
	// 幸运数
	err := LuckyNum.Run()
	if err != nil {
		panic(err)
	}
}
