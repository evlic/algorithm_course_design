package no_12

import (
	"ACD/common"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var (
	dataJ = [][]int{
		{254, 256, 256, 254, 251, 256, 254, 253, 255, 251, 251, 255},
		{3, 2, 3},
		{1, 2, 4, 7, 8},
	}
	dataK     = []int{10, 3, 2}
	expectAns = []int{504, 3, 11}
	log       = *common.GetLogger()
)

func isAccess(a, b int) bool {
	return a == b
}

func TestByBuiltinData(t *testing.T) {
	var logInfo string

	logInfo = "\n| 测试用例编号 | 程序输入以及步骤          | 期待结果（输出） | 实际结果（输出） | 是否通过 |" + "\n| ------------ | ------------------------- | ---------------- | ---------------- | -------- |"
	log.Info(logInfo)
	for idx, _ := range dataK {
		solution := Solution(dataJ[idx], dataK[idx])
		expect := expectAns[idx]
		status := isAccess(solution, expect)
		fmt.Printf(
			"| %v | J >> %v, K >> %v | %v | %v | %v | \n", idx, dataJ[idx], dataK[idx],
			expect, solution, status,
		)
	}
}

// TestByStdin 使用命令行输入
func TestByStdin(t *testing.T) {
	flag.Parse()
	args := flag.Args()
	log.Infof("args.len: %v", len(args))

	var (
		inputJ []int
		inputK int
	)

	for idx := 0; idx < len(args)-1; idx += 2 {
		inputJ = []int{}
		inputK = 0
		args[idx] = strings.Trim(args[idx], "[")
		args[idx] = strings.Trim(args[idx], "]")
		t := strings.Split(args[idx], ",")
		for _, str := range t {
			if v, err := strconv.Atoi(str); err == nil {
				inputJ = append(inputJ, v)
			}
		}

		if v, err := strconv.Atoi(args[idx+1]); err == nil {
			inputK = v
		} else {
			log.Errorln("错误输入：", v, "\t 转换为 int 失败 >> ", err)
			return
		}
		solution := Solution(inputJ, inputK)
		fmt.Printf(
			"\n| 输入 >> %v, %v | \n| %v | \n", inputJ, inputK, solution,
		)
	}
}
