package SurroundingArea

import (
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"math/rand"
	"os"
	"strings"
)

var moveRule = [2][4]int{
	//右,下,左,上
	{0, 1, 0, -1}, // row
	{1, 0, -1, 0}} // column

func Run(filePath string) (err error) {
	var (
		size   int
		matrix []string
	)
	if filePath == "" {
		size, matrix, err = generateRandArea()
	} else {
		size, matrix, err = readFile(filePath)
	}
	if err != nil {
		return
	}
	printMatrix(size, matrix)
	surroundingArea(size, matrix)
	printMatrix(size, matrix)
	return
}

func readFile(filePath string) (size int, matrix []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}
	contentSlice := strings.Fields(string(content))
	size = gconv.Int(contentSlice[0])
	matrix = make([]string, size*size)
	for i := 0; i < size*size; i++ {
		matrix[i] = contentSlice[i+1]
	}
	return
}

func generateRandArea() (size int, matrix []string, err error) {
	// input size from user
	fmt.Print("Please input the size of the area: ")
	_, err = fmt.Scan(&size)
	if err != nil {
		return
	}
	matrix = make([]string, size*size)
	for i := 0; i < size*size; i++ {
		switch rand.Intn(3) {
		case 0, 2:
			matrix[i] = "X"
		case 1:
			matrix[i] = "O"
		}
	}
	return
}

func printMatrix(size int, matrix []string) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", matrix[i*size+j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func surroundingArea(size int, matrix []string) {
	queue := glist.New()
	stack := glist.New()
	// 避开第 0 行和最后一行
	areaSize := size*size - size - 1
	for i := size + 1; i < areaSize; i++ {
		if i%size == 0 || i%size == size-1 || matrix[i] == "X" {
			// 避开第 0 列和最后一列且不为 X
			continue
		}
		// 每新到一个 O，清空队列和栈
		queue.Clear()
		stack.Clear()
		queue.PushBack(i)
		invalid := false
		// bfs
		for queue.Len() > 0 {
			index := queue.PopFront().(int)
			matrix[index] = "X"
			// 记录修改过的 O
			stack.PushBack(index)
			for j := 0; j < 4; j++ {
				nextRow := index/size + moveRule[0][j]
				nextCol := index%size + moveRule[1][j]
				if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size {
					// 越界
					invalid = true
					break
				}
				nextIndex := nextRow*size + nextCol
				if matrix[nextIndex] == "O" {
					queue.PushBack(nextIndex)
				}
			}
			if invalid {
				// 越界就退出 bfs
				break
			}
		}
		if invalid {
			// 越界就恢复被修改成 X 的 O
			for stack.Len() > 0 {
				matrix[stack.PopBack().(int)] = "O"
			}
		}
	}
}
