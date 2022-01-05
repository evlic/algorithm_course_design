package no_09

import (
	"ACD/common"
	"fmt"
	"testing"
)

var (
	data = [][][]int{
		{{1, 2, 3, 5, 2}, {3, 2, 1, 4, 2}},
		{{3, 0, 0, 0, 0, 2}, {3, 0, 0, 0, 0, 2}},
	}
	expectAns = []int{
		7,
		5,
	}
	log = *common.GetLogger()
)

// isAccess 判断答案是否符合预期
func isAccess(solution, ans int) bool {
	return solution == ans
}

func TestByBuiltinData(t *testing.T) {
	log.Info("\n| 测试用例编号 | 程序输入以及步骤          | 期待结果（输出） | 实际结果（输出） | 是否通过 |" + "\n| ------------ | ------------------------- | ---------------- | ---------------- | -------- |")

	for idx, val := range data {
		solution := Solution(val[0], val[1])
		expect := expectAns[idx]
		status := isAccess(solution, expect)
		fmt.Printf(
			"| %v | %v | %v | %v | %v | \n", idx, val,
			expect, solution, status,
		)
	}
}
