package no_04

import (
	"ACD/common"
	"flag"
	"fmt"
	"strconv"
	"testing"
)

var (
	data = [][][]int{
		{{3,4,-3,-2,-4}, {-5}},
		{{1,2,3,4}, {1}},
		{{1,3,5,7}, {1}},
		{{1,5,7,8,5,3,4,2,1}, {-2}},
	}
	expectAns = []int{
		2,
		4,
		1,
		4,
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
		solution := Solution(val[0], val[1][0])
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
	
	var input []int
	for _, v := range args {
		if i, err := strconv.Atoi(v); err == nil {
			input = append(input, i)
		} else {
			log.Errorln("错误输入：", v, "\t 转换为 int 失败 >> ", err)
			return
		}
	}
	
	solution := Solution(input[:n - 1], input[n - 1])
	fmt.Printf(
		"\n| 输入 >> %v | \n| 输出 >> %v | \n", input, solution,
	)
}
