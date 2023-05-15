package main

import (
	"algorithms/solutions/SingleSourceShortestPath"
	"fmt"
)

func main() {
	// 单源最短路径
	err := SingleSourceShortestPath.Run()
	if err != nil {
		fmt.Println(err)
	}
}
