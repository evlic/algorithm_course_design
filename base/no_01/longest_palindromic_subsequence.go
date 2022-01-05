// Package no1 题目描述： 给你一个字符串 s，找出其中最长的回文子序列，并返回该序列的长度。
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
package no1

// Solution >> 	思路
// 				区间 DP
//					小区间内的回文串状态推导更大一级区间回文串状态
//
//					区间 DP 实现：指定 len ，从小到大枚举，迭代 dp 数组
//						使用 dp[l][r] 表示区间 l >> r 回文串长度，
//					迭代过程中	  len = 1, 必然是回文串
//							  	len = 2, 若两字符相等，也是回文串，否则可能是左，也可能是右
//								len else 「s[l] == s[r]」 >> dp[l][r] <+ 2
//										 「s[l] != s[r]」 >> dp[l][r] <= dp[l + 1][r] 或 dp[l][r - 1] 取最大值
// 																		寓意 >> s[l] s[r] 中只保留一个加入回文串中
func fill(slice *[]int, val int) {
	for idx := range *slice {
		(*slice)[idx] = val
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Solution(s string) int {
	n := len(s)
	dp := make([][]int, n)

	for idx := range dp {
		dp[idx] = make([]int, n)
		fill(&(dp[idx]), 1)
	}
	// 对长度为 2 对初始化
	for idx := 1; idx < n; idx++ {
		if s[idx-1] == s[idx] {
			dp[idx-1][idx]++
		}
	}

	// 迭代 因为 1 必定是回文串，初始化时使用 fill 函数填充了初始值 1
	// 所以 len 从 1 开始 到 n 为止
	for len := 3; len <= n; len++ {
		for l, r := 0, len-1; r < n; {
			var tmp = 0
			if s[l] == s[r] {
				tmp += 2
			}
			tmp += dp[l+1][r-1]
			tmp = max(tmp, max(dp[l][r-1], dp[l+1][r]))
			dp[l][r] = tmp
			l++
			r++
		}
	}

	// 从 0 >> n - 1 也就是整个字符串最大值
	return dp[0][n-1]
}
