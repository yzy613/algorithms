package Knapsack

import (
	"fmt"
	"math/rand"
)

func Run() (err error) {
	n, w, err := inputFromConsole()
	if err != nil {
		return
	}
	weights, values := generateData(n)
	result, choice := knapsack(weights, values, w)
	printResult(weights, w, result, choice)
	return
}

func inputFromConsole() (n, w int, err error) {
	_, err = fmt.Scan(&n, &w)
	return
}

func generateData(n int) (weights, values []int) {
	weights = make([]int, n)
	values = make([]int, n)
	for i := 0; i < n; i++ {
		weights[i] = rand.Intn(10) + 1
		values[i] = rand.Intn(10) + 1
	}
	return
}

func printResult(weights []int, w int, result int, choice []int) {
	totalWeight := 0
	fmt.Print("choice: [")
	for i := range choice {
		totalWeight += weights[choice[i]]
		fmt.Print(choice[i] + 1)
		if i != len(choice)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
	fmt.Println("weight limit:", w)
	fmt.Println("total weight:", totalWeight)
	fmt.Println("total value:", result)
}

func knapsack(weights, values []int, w int) (int, []int) {
	n := len(weights)
	dp := make([][]int, n+1)
	choice := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
		choice[i] = make([][]int, w+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if j < weights[i] {
				dp[i+1][j] = dp[i][j]
				choice[i+1][j] = choice[i][j]
			} else {
				next := dp[i][j-weights[i]] + values[i]
				if dp[i][j] < next {
					dp[i+1][j] = next
					c := make([]int, len(choice[i][j-weights[i]]), len(choice[i][j-weights[i]])+1)
					copy(c, choice[i][j-weights[i]])
					choice[i+1][j] = append(c, i)
				} else {
					dp[i+1][j] = dp[i][j]
					choice[i+1][j] = choice[i][j]
				}
			}
		}
	}
	//debug(weights, values, dp, choice)
	return dp[n][w], choice[n][w]
}

func debug(weights, values []int, dp [][]int, choice [][][]int) {
	fmt.Println("weights: ", weights)
	fmt.Println("values: ", values)
	fmt.Println("dp: ")
	for i := range dp {
		fmt.Println(dp[i])
	}
	fmt.Println("choice: ")
	for i := range choice {
		fmt.Println(choice[i])
	}
}
