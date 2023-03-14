package sudoku

import "github.com/gogf/gf/v2/container/garray"

const (
	// 矩阵大小
	matrixSize = 9
	// 块大小
	blockSize = 3
)

var (
	// 数独矩阵
	matrix []*node
	// 处理数组
	dealArray *garray.Array
)

type node struct {
	Val        int
	Addr       int
	Candidates *[]int
}
