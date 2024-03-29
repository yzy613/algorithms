package main

import (
	"algorithms/common"
	"algorithms/solutions/sudoku"
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	if *common.File == "" {
		flag.Usage()
		return
	}
	// 数独
	err := sudoku.Run(*common.File)
	if err != nil {
		fmt.Println(err)
	}
}
