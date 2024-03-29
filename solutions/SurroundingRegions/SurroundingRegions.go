package SurroundingRegions

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gqueue"
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
		size, matrix, err = generateRandRegions()
	} else {
		size, matrix, err = readFile(filePath)
	}
	if err != nil {
		return
	}
	printMatrix(size, matrix)
	surroundingRegions(size, matrix)
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

func generateRandRegions() (size int, matrix []string, err error) {
	// input size from user
	fmt.Print("Please input the size of the regions: ")
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

func surroundingRegions(size int, matrix []string) {
	queue := gqueue.New()
	// 遍历最外层的一圈，遇到 O 就 bfs
	for i := 0; i < size*size; i++ {
		if i/size > 0 && i/size < size-1 && i%size == 1 {
			// 跳到这一行的最后一个位置
			i = (i/size+1)*size - 1
		}
		if matrix[i] != "O" {
			continue
		}
		queue.Push(i)
		// bfs
		for queue.Len() > 0 {
			index := queue.Pop().(int)
			matrix[index] = "Y"
			row := index / size
			col := index % size
			for j := 0; j < 4; j++ {
				nextRow := row + moveRule[0][j]
				nextCol := col + moveRule[1][j]
				if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size {
					continue
				}
				nextIndex := nextRow*size + nextCol
				if matrix[nextIndex] == "O" {
					queue.Push(nextIndex)
				}
			}
		}
	}
	for i := 0; i < size*size; i++ {
		if matrix[i] == "Y" {
			matrix[i] = "O"
		} else if matrix[i] == "O" {
			matrix[i] = "X"
		}
	}
}
