package baShuMa

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func viewMatrix(src []int) {
	for i := 0; i < matrixRow; i++ {
		for j := 0; j < matrixRow; j++ {
			fmt.Print(src[i*matrixRow+j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// DeepCopy 深拷贝，传入的变量 dst 必须是指针
func DeepCopy(src, dst interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(src); err != nil {
		return err
	}
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

// 康托展开
func CantorExpansion(src []int) (cantorNum int) {
	srcLen := len(src)
	for i := 0; i < srcLen; i++ {
		leftNum := 0
		for j := i + 1; j < srcLen; j++ {
			if src[i] > src[j] {
				leftNum++
			}
		}
		cantorNum += leftNum * factorialArr[srcLen-1-i]
	}
	return
}

// 逆康托展开
func UnCantorExpansion(cantorNum, upperLimit int) (currentArr []int) {
	currentArr = make([]int, upperLimit, upperLimit)
	usedArr := make([]bool, upperLimit, upperLimit)
	for i := upperLimit - 1; i >= 0; i-- {
		leftFigure := cantorNum/factorialArr[i] + 1
		cantorNum %= factorialArr[i]
		countNum := 0
		for j := 0; j < upperLimit; j++ {
			if !usedArr[j] {
				countNum++
			}
			if countNum == leftFigure {
				usedArr[j] = true
				// 当前数字的范围是 0 ~ upperLimit-1
				// 如果范围是 1 ~ upperLimit，请把 j 改成 j+1
				currentArr[upperLimit-1-i] = j
				break
			}
		}
	}
	return
}
