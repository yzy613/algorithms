package MergeSort

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func Run(filePath string) (err error) {
	var arr []int
	if filePath == "" {
		arr = generateRandArr()
	} else {
		arr, err = readFile(filePath)
		if err != nil {
			return
		}
	}
	fmt.Println("Unsorted Array:")
	fmt.Println(arr)
	mergeSort(arr)
	fmt.Println("Sorted Array:")
	fmt.Println(arr)
	return
}

func readFile(filePath string) (arr []int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// 读取数组
	arr = make([]int, 0)
	for {
		temp := 0
		_, e := fmt.Fscan(file, &temp)
		if e != nil {
			if e == io.EOF {
				break
			} else {
				err = e
				return
			}
		}
		arr = append(arr, temp)
	}
	return
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	arrLen := len(arr)
	mid := arrLen / 2
	mergeSort(arr[:mid])
	mergeSort(arr[mid:])
	a := make([]int, arrLen, arrLen)
	copy(a, arr)
	i, j, k := mid-1, arrLen-1, arrLen-1
	for {
		if i < 0 && j < mid {
			break
		}
		if i < 0 {
			arr[k] = a[j]
			j--
			k--
			continue
		}
		if j < mid {
			arr[k] = a[i]
			i--
			k--
			continue
		}
		if a[i] > a[j] {
			arr[k] = a[i]
			i--
		} else {
			arr[k] = a[j]
			j--
		}
		k--
	}
}

func generateRandArr() []int {
	fmt.Print("input array length: ")
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
