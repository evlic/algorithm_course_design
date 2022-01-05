package no_12

var n []int

func Solution(nums []int) (res [][]int) {
	n = nums
	backtrack([]int{}, 0, &res)
	
	return
}

func backtrack(track []int, v int, res *[][]int) {
	if len(track) == len(n) {
		*res = append(*res, track)
	}
	
	for idx := range n {
		t := 1 << idx
		if v&t == 0 {
			backtrack(append(track, n[idx]), v+t, res)
		}
	}
}
