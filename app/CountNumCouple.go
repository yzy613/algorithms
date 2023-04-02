package main

import "algorithms/solutions/CountNumCouple"

func main() {
	// 满足条件的元素对个数
	err := CountNumCouple.Run()
	if err != nil {
		panic(err)
	}
}
