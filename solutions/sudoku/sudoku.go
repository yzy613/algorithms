package sudoku

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"os"
)

func Run(filePath string) (err error) {
	err = readFile(filePath)
	if err != nil {
		return
	}
	dealArray = garray.NewArray()
	// 扫描矩阵
	if !scanMatrix() {
		fmt.Println("无解")
		return
	}
	// 按候选数个数排序
	dealArray.SortFunc(func(a, b interface{}) bool {
		return len(*a.(*node).Candidates)-len(*b.(*node).Candidates) < 0
	})
	// 递归搜索
	if dfs(0) {
		printMatrix()
	} else {
		fmt.Println("无解")
	}
	return
}

func readFile(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// 读取矩阵
	matrix = make([]*node, matrixSize*matrixSize, matrixSize*matrixSize)
	for i := 0; i < matrixSize*matrixSize; i++ {
		temp := 0
		_, err = fmt.Fscan(file, &temp)
		if err != nil {
			return
		}
		matrix[i] = &node{
			Val:  temp,
			Addr: i,
		}
	}
	return
}

func scanMatrix() (success bool) {
	for i := range matrix {
		if matrix[i].Val == 0 {
			temp := searchCandidates(i)
			candidatesLen := len(*temp)
			switch candidatesLen {
			case 0:
				return
			case 1:
				matrix[i].Val = (*temp)[0]
			default:
				matrix[i].Candidates = temp
				dealArray.Append(matrix[i])
			}
		}
	}
	return true
}

func searchCandidates(addr int) (candidates *[]int) {
	c := make([]int, 0, matrixSize)
	for i := 1; i <= 9; i++ {
		if isValid(addr, i) {
			c = append(c, i)
		}
	}
	return &c
}

func isValid(addr int, val int) bool {
	return !isInRow(addr, val) && !isInColumn(addr, val) && !isInBlock(addr, val)
}

func isInRow(addr int, val int) bool {
	rowNum := addr / matrixSize * matrixSize
	for i := 0; i < matrixSize; i++ {
		if matrix[rowNum+i].Val == val {
			return true
		}
	}
	return false
}

func isInColumn(addr int, val int) bool {
	columnNum := addr % matrixSize
	for i := 0; i < matrixSize; i++ {
		if matrix[i*matrixSize+columnNum].Val == val {
			return true
		}
	}
	return false
}

func isInBlock(addr int, val int) bool {
	blockRowStartAt := addr / matrixSize / blockSize * blockSize
	blockColumnStartAt := addr % matrixSize / blockSize * blockSize
	for i := 0; i < blockSize; i++ {
		rowNum := (blockRowStartAt + i) * matrixSize
		for j := 0; j < blockSize; j++ {
			if matrix[rowNum+blockColumnStartAt+j].Val == val {
				return true
			}
		}
	}
	return false
}

func printMatrix() {
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			if j != 0 && j%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Print(matrix[i*matrixSize+j].Val, " ")
		}
		if i != 0 && i%3 == 2 && i != matrixSize-1 {
			fmt.Println()
			for j := 0; j < matrixSize; j++ {
				if j != 0 && j%3 == 0 {
					fmt.Print("+ ")
				}
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}
