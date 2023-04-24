package main

import (
	"algorithms/solutions/ShipLoading"
	"fmt"
)

func main() {
	// 货船装载
	err := ShipLoading.Run()
	if err != nil {
		fmt.Println(err)
	}
}
