package SingleSourceShortestPath

import (
	"fmt"
	"math/rand"
)

func Run() (err error) {
	size, err := inputFromConsole()
	if err != nil {
		return
	}
	start := 0
	graph := generateGraph(size, start)
	printGraph(graph)
	result := dijkstra(graph, start)
	printResult(result, start)
	return
}

func inputFromConsole() (size int, err error) {
	_, err = fmt.Scan(&size)
	return
}

func generateGraph(size int, start int) (graph [][]int) {
	graph = make([][]int, size)
	for i := 0; i < size; i++ {
		graph[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if j == start || i == j {
				graph[i][j] = unreachable
				continue
			}
			temp := rand.Intn(101)
			if temp%3 == 0 {
				temp = unreachable
			}
			graph[i][j] = temp
		}
	}
	return
}

func printGraph(graph [][]int) {
	for i := range graph {
		for j := range graph[i] {
			fmt.Printf("%3d ", graph[i][j])
		}
		fmt.Println()
	}
}

func printResult(result []singleSourceShortestPath, start int) {
	for i := range result {
		if i == start {
			continue
		}
		fmt.Printf("start: %d, end: %d, length: %d\n", start+1, i+1, result[i].Length)
		pathLen := len(result[i].Path)
		for j := range result[i].Path {
			fmt.Print(result[i].Path[j] + 1)
			if j != pathLen-1 {
				fmt.Print("->")
			}
		}
		fmt.Println()
	}
}
