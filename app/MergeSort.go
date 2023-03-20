package main

import (
	"algorithms/common"
	"algorithms/solutions/MergeSort"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	// 归并排序
	err := MergeSort.Run(*common.File)
	if err != nil {
		panic(err)
	}
}
