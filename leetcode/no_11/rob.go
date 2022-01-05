package no_11

import "ACD/common"

func Solution(ns []int) int {
	n := len(ns)
	if n == 1 {
		return ns[0]
	}
	if n == 2 {
		return common.MaxInt(ns[0], ns[1])
	}
	dp := make([]int, n+1)

	// 初始化
	for i := 1; i <= 2 && i < n; i++ {
		dp[i] = ns[i-1]
	}

	for i := 2; i <= n; i++ {
		dp[i] = common.MaxInt(dp[i-1], dp[i-2]+ns[i-1])
	}

	return dp[n]
}
