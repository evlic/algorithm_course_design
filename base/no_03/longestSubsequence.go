package no_03

import (
	"ACD/common"
)

func Solution(arr []int, d int) int {
	n := len(arr)
	dp := make([][]int, n)

	// 数 >> 对应逆序下标
	mapp := make(map[int]int)

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 初始化
	dp[0][1] = 1
	mapp[arr[0]] = 0

	for i := 1; i < n; i++ {
		// 选择默认为 1
		dp[i][1] = 1

		dp[i][0] = common.MaxInt(dp[i-1][0], dp[i-1][1])

		keyCur, keyPrv := arr[i], arr[i]-d

		if v, has := mapp[keyPrv]; has {
			// 如果存在，则计算状态转移值
			dp[i][1] = dp[v][1] + 1
		}

		mapp[keyCur] = i

	}

	return common.MaxInt(dp[n-1][0], dp[n-1][1])
}
