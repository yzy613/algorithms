package main

import (
	"algorithms/solutions/TravelingSalesmanProblem"
	"fmt"
)

func main() {
	// 旅行商问题
	err := TravelingSalesmanProblem.Run()
	if err != nil {
		fmt.Println(err)
	}
}
