package no_06

import (
	"ACD/common"
	"flag"
	"fmt"
	"strconv"
	"testing"
)

var (
	data = [][]int{{3, 2, 1, 6, 0, 5}, {3, 2, 1}}
	// 先序
	expectAns = []string{
		"6,3,null,2,null,1,null,null,5,0,null,null,null,", "3,null,2,null,1,null,null,",
	}

	log = *common.GetLogger()
)

func isAccess(root *TreeNode, ans string) bool {
	if root.String() == ans {
		return true
	}
	return false
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
	log.Infof("您的入参长： %v", len(args))

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
		"\n| 输入 >> %v | \n| 输出 >> %v | \n", input, solution.String(),
	)
}
