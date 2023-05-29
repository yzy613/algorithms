package main

import (
	"algorithms/solutions/LongestCommonSubsequence"
	"fmt"
)

func main() {
	// 最长公共子序列
	// https://leetcode-cn.com/problems/longest-common-subsequence/
	err := LongestCommonSubsequence.Run()
	if err != nil {
		fmt.Println(err)
	}
}
