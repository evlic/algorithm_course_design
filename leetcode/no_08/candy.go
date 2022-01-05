package no_08

import "ACD/common"

func Solution(r []int) int {
	n := len(r)
	a := make([]int, n)

	for idx, _ := range a {
		a[idx] = 1
	}

	// 从左向右比较，保证分比左侧高的获得多 1 个
	for i := 1; i < n; i++ {
		if r[i] > r[i-1] {
			a[i] = a[i-1] + 1
		}
	}

	// 补充
	// 从右向左比较，保证分比右高的获得更多 1 个
	res := a[n-1]
	for i := n - 2; i > -1; i-- {
		if r[i] > r[i+1] {
			a[i] = common.MaxInt(a[i], a[i+1]+1)
		}
		res += a[i]
	}

	return res
}
