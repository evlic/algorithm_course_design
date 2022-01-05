package no_13

var n []int

func Solution(nums []int) (res [][]int) {
	n = nums
	// 000000 << int32  11111111
	backtrack([]int{}, 0, &res)

	return
}

func backtrack(track []int, v int, res *[][]int) {
	// 结束条件
	if len(track) == len(n) {
		// res <+ track
		*res = append(*res, track)
	}

	// 决策列表
	for idx := range n {
		t := 1 << idx
		if v&t == 0 {
			backtrack(append(track, n[idx]), v+t, res)
		}
	}
}
