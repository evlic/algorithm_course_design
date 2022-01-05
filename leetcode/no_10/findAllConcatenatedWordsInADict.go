package no_10

import "ACD/common"

var (
	set map[int64]struct{}

	E               = struct{}{}
	P, OFFSET int64 = 13131, 128
)

func Solution(words []string) []string {
	set = make(map[int64]struct{})

	for _, str := range words {
		var hash int64
		for idx, _ := range str {
			hash = hash*P + int64(str[idx]) + OFFSET
		}
		set[hash] = E
	}
	// fmt.Println(set)
	var res []string
	for _, str := range words {
		if check(str) {
			res = append(res, str)
		}
	}
	return res
}

func check(s string) bool {
	n := len(s)
	dp := make([]int, n+1)

	for idx := 1; idx <= n; idx++ {
		dp[idx] = -1
	}

	for i := 0; i <= n; i++ {
		// 剪枝
		if dp[i] == -1 {
			continue
		}
		var hash int64
		for j := i + 1; j <= n; j++ {
			hash = hash*P + int64(s[j-1]) + OFFSET
			// 如果当前 hash 存在，则状态转移
			if _, has := set[hash]; has {
				// fmt.Println(hash)
				dp[j] = common.MaxInt(dp[j], dp[i]+1)
			}
		}
		// 达成条件 提前退出
		if dp[n] > 1 {
			return true
		}
	}

	return false
}
