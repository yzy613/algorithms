package CountNumCouple

import "fmt"

func Run() (err error) {
	arr, target, err := input()
	if err != nil {
		return
	}
	count := start(arr, target)
	fmt.Println(count)
	return
}

func input() (arr []int, target int, err error) {
	n := 0
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&target)
	if err != nil {
		return
	}
	arr = make([]int, n, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&arr[i])
		if err != nil {
			return
		}
	}
	return
}

func start(arr []int, target int) (count int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]-arr[j] == target {
				count++
			} else if arr[j]-arr[i] == target {
				count++
			}
		}
	}
	return
}
