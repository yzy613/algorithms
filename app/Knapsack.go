package main

import (
	"algorithms/solutions/Knapsack"
	"fmt"
)

func main() {
	// Knapsack problem
	err := Knapsack.Run()
	if err != nil {
		fmt.Println(err)
	}
}
