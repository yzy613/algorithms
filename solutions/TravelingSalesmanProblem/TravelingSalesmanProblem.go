package TravelingSalesmanProblem

import (
	"fmt"
	"math/rand"
)

func Run() (err error) {
	cityNum, err := inputCityNumFromConsole()
	if err != nil {
		return
	}
	cityMap := generateCityMap(cityNum)
	printCityMap(cityMap)
	startCity := rand.Intn(cityNum)
	// 剪枝的 bfs
	path, length := bfs(cityMap, startCity)
	// for progressBar()
	fmt.Println()
	// 剪枝结果
	fmt.Println("剪枝结果: ")
	printPathAndLength(path, length)
	//// 不剪枝的 bfs
	//path, length = bfsSlow(cityMap, startCity)
	//// for progressBar()
	//fmt.Println()
	//// 验证结果
	//fmt.Println("验证结果: ")
	//printPathAndLength(path, length)
	return
}

func inputCityNumFromConsole() (cityNum int, err error) {
	_, err = fmt.Scan(&cityNum)
	return
}

func generateCityMap(cityNum int) (cityMap [][]int) {
	cityMap = make([][]int, cityNum)
	for i := 0; i < cityNum; i++ {
		cityMap[i] = make([]int, cityNum)
		for j := 0; j < cityNum; j++ {
			if i == j {
				continue
			}
			cityMap[i][j] = rand.Intn(20) + 1
		}
	}
	return
}

func printCityMap(cityMap [][]int) {
	cityNum := len(cityMap)
	for i := 0; i < cityNum; i++ {
		for j := 0; j < cityNum; j++ {
			fmt.Printf("%2d ", cityMap[i][j])
		}
		fmt.Println()
	}
}

func printPathAndLength(path []int, length int) {
	pathLen := len(path)
	for i := 0; i < pathLen; i++ {
		fmt.Print(path[i] + 1)
		if i != pathLen-1 {
			fmt.Print("->")
		}
	}
	fmt.Println()
	fmt.Println(length)
}
