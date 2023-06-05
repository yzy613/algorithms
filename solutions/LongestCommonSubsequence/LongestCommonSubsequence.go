package LongestCommonSubsequence

import (
	"fmt"
	"math/rand"
	"strings"
)

func Run() (err error) {
	m, n, err := inputFromConsole()
	if err != nil {
		return
	}
	text1, text2 := generateText(m, n)
	fmt.Println("text1:", text1)
	fmt.Println("text2:", text2)
	result, subsequence := longestCommonSubsequence(text1, text2)
	fmt.Println("L.C.S.:", result)
	fmt.Println("subsequence:", subsequence)
	return
}

func inputFromConsole() (m, n int, err error) {
	_, err = fmt.Scan(&m, &n)
	return
}

func generateText(m, n int) (text1, text2 string) {
	characterSet := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	t1Builder := strings.Builder{}
	t1Builder.Grow(m)
	for i := 0; i < m; i++ {
		temp := rand.Intn(len(characterSet))
		t1Builder.WriteRune(characterSet[temp])
	}
	text1 = t1Builder.String()
	t2Builder := strings.Builder{}
	t2Builder.Grow(n)
	for i := 0; i < n; i++ {
		temp := rand.Intn(len(characterSet))
		t2Builder.WriteRune(characterSet[temp])
	}
	text2 = t2Builder.String()
	return
}

func longestCommonSubsequence(text1, text2 string) (int, string) {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	subsequence := make([][][]rune, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		subsequence[i] = make([][]rune, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
				rs := make([]rune, len(subsequence[i][j]), len(subsequence[i][j])+1)
				copy(rs, subsequence[i][j])
				subsequence[i+1][j+1] = append(rs, c1)
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
					subsequence[i+1][j+1] = subsequence[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
					subsequence[i+1][j+1] = subsequence[i+1][j]
				}
			}
		}
	}
	//debug(text1, text2, dp)
	return dp[m][n], string(subsequence[m][n])
}

func debug(text1, text2 string, dp [][]int) {
	for i, row := range dp {
		if i == 0 {
			fmt.Print("    ")
			for _, c := range text2 {
				fmt.Printf("%c ", c)
			}
			fmt.Println()
		}
		for j, col := range row {
			if j == 0 {
				if i == 0 {
					fmt.Print("  ")
				} else {
					fmt.Printf("%c ", text1[i-1])
				}
			}
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
