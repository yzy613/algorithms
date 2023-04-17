package OptimalDelivery

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	shopPoint     = 1
	customerPoint = 2
	obstaclePoint = 3
)

type point struct {
	Role        int
	NeedDeliver int
	Cost        int
}

func Run(filePath string) (err error) {
	var (
		size        int
		matrix      []point
		shopIndex   []int
		customerNum int
	)
	if filePath == "" {
		matrix, shopIndex, size, customerNum, err = input(os.Stdin)

	} else {
		var fileReader *os.File
		fileReader, err = os.Open(filePath)
		if err != nil {
			return
		}
		matrix, shopIndex, size, customerNum, err = input(fileReader)
	}
	if err != nil {
		return
	}
	// debug
	//printMatrix(size, matrix)
	fmt.Println(bfs(matrix, size, customerNum, shopIndex))
	// debug
	//printMatrix(size, matrix)
	return
}

func input(src io.Reader) (matrix []point, shopIndex []int, size, customerNum int, err error) {
	obstaclesNum, shopNum := 0, 0
	sc := bufio.NewScanner(src)
	sc.Scan()
	_, err = fmt.Sscanf(sc.Text(), "%d %d %d %d", &size, &shopNum, &customerNum, &obstaclesNum)
	if err != nil {
		return
	}
	matrix = make([]point, size*size)
	// scan for shop points
	for i := 0; i < shopNum; i++ {
		var row, col int
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "%d %d", &row, &col)
		if err != nil {
			return
		}
		// 坐标的开始从 (1,1) 修正为 (0,0)
		row--
		col--
		index := row*size + col
		matrix[index].Role = shopPoint
		shopIndex = append(shopIndex, index)
	}
	// scan for customer points
	numOfCustomer := customerNum
	for i := 0; i < numOfCustomer; i++ {
		var row, col, need int
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "%d %d %d", &row, &col, &need)
		if err != nil {
			return
		}
		// 坐标的开始从 (1,1) 修正为 (0,0)
		row--
		col--
		index := row*size + col
		matrix[index].Role = customerPoint
		// 可能有多个客户在方格图中的同一个位置
		if matrix[index].NeedDeliver != 0 {
			customerNum--
		}
		matrix[index].NeedDeliver += need
	}
	// scan for obstacle points
	for i := 0; i < obstaclesNum; i++ {
		var row, col int
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "%d %d", &row, &col)
		if err != nil {
			return
		}
		// 坐标的开始从 (1,1) 修正为 (0,0)
		row--
		col--
		matrix[row*size+col].Role = obstaclePoint
	}
	return
}

// for debug
func printMatrix(size int, matrix []point) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", matrix[i*size+j].Role)
		}
		fmt.Println()
	}
	fmt.Println()
}
