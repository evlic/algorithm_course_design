// Package no1 题目描述： 给你一个字符串 s，找出其中最长的回文子序列，并返回该序列的长度。
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
package no1

import "ACD/common"

var max = common.MaxInt

func Solution(s string) int {
	n := len(s)
	dp := make([][]int, n)

	// 初始化数组 n * n
	for idx := range dp {
		dp[idx] = make([]int, n)
		dp[idx][idx] = 1
	}

	for i := n - 1; i >= 0; i-- {

		for j := i + 1; j < n; j++ {
			if s[j] == s[i] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], max(dp[i][j], dp[i][j-1]))
			}
			// fmt.Println(dp)
		}
	}

	// 从 0 >> n - 1 也就是整个字符串最大值
	return dp[0][n-1]
}
