package no_05

import (
	"ACD/common"
	"flag"
	"fmt"
	"testing"
)

var (
	data = [][]string{
		{"horse", "ros"},
		{"intention", "execution"},
	}
	expectAns = []int{
		3,
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

// TestByStdin 使用命令行输入
func TestByStdin(t *testing.T) {
	flag.Parse()
	args := flag.Args()
	n := len(args)
	log.Infof("您的入参长： %v", n)

	for idx := 0; idx < len(args)-1; idx += 2 {
		solution := Solution(args[idx], args[idx+1])
		fmt.Printf(
			"\n| %v| \n|输入 >> %v, %v | \n| 输出 >> %v | \n", idx/2, args[idx], args[idx+1], solution,
		)
	}
}
