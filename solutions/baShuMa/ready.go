package baShuMa

import (
	"errors"
	"fmt"
	"os"
)

// 从文件读写 (默认从 data.txt 读取)
func inputAndPreprocess() (srcZeroPosition, dstZeroPosition int, err error) {
	srcFile, getErr := os.Open("data.txt")
	if getErr != nil {
		err = getErr
		return
	}
	defer srcFile.Close()

	_, err = fmt.Fscanln(srcFile, &matrixRow)
	if err != nil {
		return
	}

	// 初始矩阵查重用
	verificationArr := make([]bool, matrixRow*matrixRow, matrixRow*matrixRow)

	// 读取初始矩阵
	srcArr = make([]int, 0, matrixRow*matrixRow)
	for i := 0; i < matrixRow*matrixRow; i++ {
		temp := 0
		_, getErr = fmt.Fscan(srcFile, &temp)
		if getErr != nil {
			err = getErr
			return
		}
		if verificationArr[temp] {
			err = errors.New("初始矩阵有重复方块")
		} else {
			verificationArr[temp] = true
			if temp == 0 {
				srcZeroPosition = i
			}
		}
		srcArr = append(srcArr, temp)
	}

	// 目标矩阵查重用
	verificationArr = make([]bool, matrixRow*matrixRow, matrixRow*matrixRow)

	// 读取目标矩阵
	dstArr = make([]int, 0, matrixRow*matrixRow)
	for i := 0; i < matrixRow*matrixRow; i++ {
		temp := 0
		_, getErr = fmt.Fscan(srcFile, &temp)
		if getErr != nil {
			err = getErr
			return
		}
		if verificationArr[temp] {
			err = errors.New("目标矩阵有重复方块")
		} else {
			verificationArr[temp] = true
			if temp == 0 {
				dstZeroPosition = i
			}
		}
		dstArr = append(dstArr, temp)
	}

	// 预处理阶乘
	preprocessFactorialArr(matrixRow * matrixRow)
	return
}

func preprocessFactorialArr(upperLimit int) {
	factorialArr = make([]int, 0, upperLimit)
	figure := 1
	factorialArr = append(factorialArr, figure)
	for i := 1; i <= upperLimit; i++ {
		figure *= i
		factorialArr = append(factorialArr, figure)
	}
}

func verify(srcZeroPosition, dstZeroPosition int) (err error) {
	srcArrTemp, dstArrTemp := new([]int), new([]int)
	err = DeepCopy(srcArr, srcArrTemp)
	if err != nil {
		return
	}
	err = DeepCopy(dstArr, dstArrTemp)
	if err != nil {
		return
	}
	// 行列数为奇数时，初始矩阵的逆序对和目标矩阵的逆序对不是同奇或同偶，则不能被还原
	if matrixRow%2 == 1 {
		if countInverseNumber(*srcArrTemp)%2 != countInverseNumber(*dstArrTemp)%2 {
			err = errors.New("初始矩阵不能被还原")
			return
		}
	} else {
		// 逆序对是奇数，待还原和初始空白块的行数差值也是奇数，偶数反之
		inverseNumber := countInverseNumber(*srcArrTemp)
		switch {
		// +2 是因为要把从 0 开始计数变为从 1 开始（行和列都要+1）
		case inverseNumber%2 == 1 && (srcZeroPosition-dstZeroPosition+2)/2 == 0:
			fallthrough
		case inverseNumber%2 == 0 && (srcZeroPosition-dstZeroPosition+2)/2 == 1:
			err = errors.New("初始矩阵不能被还原")
			return
		}
	}
	return
}

func countInverseNumber(src []int) int {
	swapQ := make([]int, 0)
	inverseNumber := 0
	mergeSort(&src, &swapQ, 0, len(src)-1, &inverseNumber)
	return inverseNumber
}

func mergeSort(src, queue *[]int, left, right int, ans *int) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	mergeSort(src, queue, left, mid, ans)
	mergeSort(src, queue, mid+1, right, ans)
	i, j := left, mid+1
	for i <= mid && j <= right {
		if (*src)[i] <= (*src)[j] {
			*queue = append(*queue, (*src)[i])
			i++
		} else {
			*queue = append(*queue, (*src)[j])
			*ans += mid - i + 1
			j++
		}
	}
	for i <= mid {
		*queue = append(*queue, (*src)[i])
		i++
	}
	for j <= right {
		*queue = append(*queue, (*src)[j])
		j++
	}
	for addr := left; addr <= right; addr++ {
		(*src)[addr] = (*queue)[0]
		// 队列 pop
		*queue = (*queue)[1:]
	}
}
