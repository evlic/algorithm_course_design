package no_05

import "ACD/common"

func Solution(w1 string, w2 string) int {
	// dp
	n, m := len(w1), len(w2)
	dp := make([][]int, n+1)

	// 初始化 m+n
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}

	for j := range dp[0] {
		dp[0][j] = j
	}

	// 枚举 i、j
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.MinIntSliceVal([]int{dp[i-1][j], dp[i-1][j-1], dp[i][j-1]}) + 1
			}
		}
	}
	return dp[n][m]
}
