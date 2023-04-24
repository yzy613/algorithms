package main

import (
	"algorithms/solutions/CountNumCouple"
	"fmt"
)

func main() {
	// 满足条件的元素对个数
	err := CountNumCouple.Run()
	if err != nil {
		fmt.Println(err)
	}
}
