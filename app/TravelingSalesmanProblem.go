package main

import (
	"algorithms/solutions/TravelingSalesmanProblem"
	"fmt"
)

func main() {
	err := TravelingSalesmanProblem.Run()
	if err != nil {
		fmt.Println(err)
	}
}
