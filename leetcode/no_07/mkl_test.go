package no_7

import (
	"ACD/common"
	"fmt"
	"testing"
)

var (
	data = [][][]int{
		{{1, 4, 5}, {1, 3, 4}, {2, 6}},
		{{}},
	}
	expectAns = [][]int{
		{1, 1, 2, 3, 4, 4, 5, 6},
		{},
	}
	log = *common.GetLogger()
)

// isAccess 判断答案是否符合预期
func isAccess(solution *ListNode, ans []int) bool {
	for _, v := range ans {
		if solution == nil || solution.Val != v {
			return false
		}
		solution = solution.Next
	}
	return true
}

func TestByBuiltinData(t *testing.T) {
	log.Info("\n| 测试用例编号 | 程序输入以及步骤          | 期待结果（输出） | 实际结果（输出） | 是否通过 |" + "\n| ------------ | ------------------------- | ---------------- | ---------------- | -------- |")
	
	for idx, val := range data {
		var ps []*ListNode
		for _, v := range val {
			p := &ListNode{}
			p.Insert(v)
			ps = append(ps, p.Next)
		}
		solution := Solution(ps)
		expect := expectAns[idx]
		status := isAccess(solution, expect)
		fmt.Printf(
			"| %v | %v | %v | %v | %v | \n", idx, val,
			expect, solution, status,
		)
	}
}

// // TestByStdin 使用命令行输入
// func TestByStdin(t *testing.T) {
// 	flag.Parse()
// 	args := flag.Args()
// 	log.Infof("您的入参长： %v", len(args))
// 
// }
