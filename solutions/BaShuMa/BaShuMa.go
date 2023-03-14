package BaShuMa

import (
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
	"os"
)

func Run(filePath string) (err error) {
	err = readFile(filePath)
	if err != nil {
		return
	}
	dstMatrixKey = matrix2Key(dstMatrix)
	beginNode := newBeginNode()
	err = start(beginNode)
	return
}

func readFile(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// 读取矩阵大小
	_, err = fmt.Fscanln(file, &matrixSize)
	// 查重矩阵
	once := make([]bool, matrixSize*matrixSize, matrixSize*matrixSize)
	// 读取源矩阵
	srcMatrix = make([]int, matrixSize*matrixSize)
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			temp := 0
			_, err = fmt.Fscan(file, &temp)
			if err != nil {
				return
			}
			if once[temp] {
				err = fmt.Errorf("srcMatrix has duplicate number")
				return
			}
			once[temp] = true
			srcMatrix[i*matrixSize+j] = temp
		}
	}
	// 验证是否写全
	for i := range once {
		if !once[i] {
			err = fmt.Errorf("srcMatrix has missing number")
			return
		}
	}
	// 重置查重矩阵
	once = make([]bool, matrixSize*matrixSize, matrixSize*matrixSize)
	// 读取目标矩阵
	dstMatrix = make([]int, matrixSize*matrixSize)
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			temp := 0
			_, err = fmt.Fscan(file, &temp)
			if err != nil {
				return
			}
			if once[temp] {
				err = fmt.Errorf("dstMatrix has duplicate number")
				return
			}
			once[temp] = true
			dstMatrix[i*matrixSize+j] = temp
		}
	}
	// 验证是否写全
	for i := range once {
		if !once[i] {
			err = fmt.Errorf("dstMatrix has missing number")
			return
		}
	}
	return
}

func newBeginNode() (beginNode *node) {
	addr := 0
	for i := range srcMatrix {
		if srcMatrix[i] == 0 {
			addr = i
			break
		}
	}
	beginNode = &node{
		Step:         0,
		ZeroPosition: addr,
		Matrix:       &srcMatrix,
		ParentKey:    "",
	}
	return
}

func printResult(curr *node) {
	fmt.Println()
	list := glist.New()
	for curr.ParentKey != "" {
		list.PushBack(curr)
		curr = hashTable.Get(curr.ParentKey).(*node)
	}
	list.PushBack(curr)
	listLen := list.Len()
	for i := 0; i < listLen; i++ {
		curr = list.PopBack().(*node)
		printMatrix(*curr.Matrix)
		fmt.Println()
	}
	fmt.Printf("需 %v 步完成\n", listLen-1)
}

func printMatrix(matrix []int) {
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			fmt.Print(matrix[i*matrixSize+j], " ")
		}
		fmt.Println()
	}
}
