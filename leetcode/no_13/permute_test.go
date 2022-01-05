package no_13

import (
	"ACD/common"
	"flag"
	"fmt"
	"sort"
	"strconv"
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

// TestByStdin 使用命令行输入
func TestByStdin(t *testing.T) {
	flag.Parse()
	args := flag.Args()
	log.Infof("由于题目使用位运算优化，仅支持最多 32 位输入 （题目限制为 6）, 您的入参长： %v", len(args))

	var input []int
	for _, v := range args {
		if i, err := strconv.Atoi(v); err == nil {
			input = append(input, i)
		} else {
			log.Errorln("错误输入：", v, "\t 转换为 int 失败 >> ", err)
			return
		}
	}

	solution := Solution(input)
	fmt.Printf(
		"\n| 输入 >> %v | \n| 全排列 >> %v | \n| 共计 >> %v | \n", input, solution, len(solution),
	)
}
