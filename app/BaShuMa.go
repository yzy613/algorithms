package main

import (
	"algorithms/common"
	"algorithms/solutions/BaShuMa"
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	if *common.File == "" {
		flag.Usage()
		return
	}
	// 八数码
	err := BaShuMa.Run(*common.File)
	if err != nil {
		fmt.Println(err)
	}
}
