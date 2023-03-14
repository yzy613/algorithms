package main

import (
	"algorithms/solutions/BaShuMa"
	flag "github.com/spf13/pflag"
)

var (
	file = flag.StringP("file", "f", "", "file path")
)

func main() {
	flag.Parse()
	if *file == "" {
		flag.Usage()
		return
	}
	// 运行指定算法
	err := BaShuMa.Run(*file)
	if err != nil {
		panic(err)
	}
}
