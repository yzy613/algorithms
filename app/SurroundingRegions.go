package main

import (
	"algorithms/common"
	"algorithms/solutions/SurroundingRegions"
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	// 环绕区域
	err := SurroundingRegions.Run(*common.File)
	if err != nil {
		fmt.Println(err)
	}
}
