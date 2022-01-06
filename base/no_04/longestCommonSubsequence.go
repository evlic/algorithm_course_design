package no_04

import "ACD/common"

func Solution(t1 string, t2 string) int {
	n, m := len(t1), len(t2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if t1[i-1] == t2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = common.MaxInt(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n][m]
}
