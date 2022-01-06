package no_02

import "ACD/common"

func Solution(ns []int) (maxAns int) {
	n := len(ns)
	maxAns = 1
	if n == 0 {
		return
	}

	// max 用于记录 idx 前的最大值下标
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		// 枚举 j
		for j := 0; j < i; j++ {
			if ns[j] < ns[i] {
				// 转移，每次有可选的进行比较
				dp[i] = common.MaxInt(dp[i], dp[j]+1)
			}
		}
		maxAns = common.MaxInt(maxAns, dp[i])
	}

	return
}
