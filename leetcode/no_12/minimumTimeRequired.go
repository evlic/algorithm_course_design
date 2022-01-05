package no_12

import "ACD/common"

var (
	ans, n, k int
	j         []int
)

// backtrack 回溯 >> DFS
//          idx 当前尝试分配工作的下标
//          max 当前最大工时
//          track 当前方案所有人的工时情况
func backtrack(idx, lowLimit int, track []int) {
	// 剪枝， 当此时 lowLimit 大于 之前获得的 工时结果时， 直接剪掉
	if lowLimit >= ans {
		return
	}

	// 当一次 dfs 搜索结束，赋值并结束，此时 lowLimit 一定比 ans 小
	if idx == n {
		ans = lowLimit
		return
	}

	// 把工作分配给每一个人
	for i := 0; i < k; i++ {
		track[i] += j[idx]
		backtrack(idx+1, common.MaxInt(lowLimit, track[i]), track)
		track[i] -= j[idx]
	}
}

func Solution(jobs []int, _k int) int {
	j, k, n, ans = jobs, _k, len(jobs), 0xfffffff

	backtrack(0, 0, make([]int, k))
	return ans
}
