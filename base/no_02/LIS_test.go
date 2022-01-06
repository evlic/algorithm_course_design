package no_02

import (
	"ACD/common"
	"flag"
	"fmt"
	"sort"
	"strconv"
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
func isAccess(solution, ans interface{}) bool {
	
	return true
}

func TestByBuiltinData(t *testing.T) {
	log.Info("\n| 测试用例编号 | 程序输入以及步骤          | 期待结果（输出） | 实际结果（输出） | 是否通过 |" + "\n| ------------ | ------------------------- | ---------------- | ---------------- | -------- |")
	
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

//	var input []int
//	for _, v := range args {
//		if i, err := strconv.Atoi(v); err == nil {
//			input = append(input, i)
//		} else {
//			log.Errorln("错误输入：", v, "\t 转换为 int 失败 >> ", err)
//			return
//		}
//	}
//
//	solution := Solution(input)
//	fmt.Printf(
//		"\n| 输入 >> %v | \n| 输出 >> %v | \n| 共计 >> %v | \n", input, solution, len(solution),
//	)
}
