package main

import (
	"algorithms/common"
	"algorithms/solutions/OptimalDelivery"
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	// 最优配送
	err := OptimalDelivery.Run(*common.File)
	if err != nil {
		fmt.Println(err)
	}
}
