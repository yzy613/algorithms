package main

import (
	"algorithms/common"
	"algorithms/solutions/SurroundingArea"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	// 环绕区域
	err := SurroundingArea.Run(*common.File)
	if err != nil {
		panic(err)
	}
}
