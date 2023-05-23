package main

import (
	"algorithms/solutions/MultiMachineScheduling"
	"fmt"
)

func main() {
	// 多机调度问题
	err := MultiMachineScheduling.Run()
	if err != nil {
		fmt.Println(err)
	}
}
