package LuckyNum

import "fmt"

func Run() (err error) {
	num, err := input()
	if err != nil {
		return
	}
	count := start(num)
	fmt.Println(count)
	return
}

func input() (num int, err error) {
	fmt.Print("enter n: ")
	_, err = fmt.Scan(&num)
	return
}

func start(num int) (count int) {
	for i := 1; i <= num; i++ {
		if lucky(i) {
			count++
		}
	}
	return
}

func lucky(num int) (found bool) {
	fNum := num
	f := 0
	for fNum > 0 {
		f += fNum % 10
		fNum /= 10
	}
	gNum := num
	g := 0
	for gNum > 0 {
		g += gNum & 1
		gNum >>= 1
	}
	return f == g
}
