package no_12

import (
	"ACD/common"
	"fmt"
	"sort"
	"testing"
)

var (
	data      = [][]int{{1, 2, 3}, {0, 1}, {1}}
	expectAns = [][][]int{{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, {{0, 1}, {1, 0}}, {{1}}}
	log       = *common.GetLogger()
)

func isAccess(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	
	var sortByVal = func(v [][]int) {
		sort.Slice(
			v, func(i, j int) bool {
				if v[i][0] == v[j][0] {
					return v[i][1] > v[j][1]
				}
				return v[i][0] > v[j][0]
			},
		)
	}
	sortByVal(a)
	sortByVal(b)
	// log.Info("sort.after >> ", "a >> ", a, " b >> ", b)
	for idx := range a {
		aa, bb := a[idx], b[idx]
		
		if len(aa) != len(bb) {
			return false
		}
		for idy := range aa {
			if aa[idy] != bb[idy] {
				return false
			}
		}
	}
	return true
}

func TestByBuiltinData(t *testing.T) {
	var logInfo string
	
	logInfo = "\n| 测试用例编号 | 程序输入以及步骤          | 期待结果（输出） | 实际结果（输出） | 是否通过 |" + "\n| ------------ | ------------------------- | ---------------- | ---------------- | -------- |"
	log.Info(logInfo)
	for idx, val := range data {
		solution := Solution(val)
		expect := expectAns[idx]
		status := isAccess(solution, expect)
		fmt.Printf(
			"| %v | %v | %v | %v | %v | \n", idx, val,
			expect, solution, status,
		)
	}
}

func subAppend(o *[][]int) {
	*o = append(*o, []int{11, 22, 44})
}

func TestAppendInts(t *testing.T) {
	var origin [][]int
	subAppend(&origin)
	fmt.Println(origin)
}
