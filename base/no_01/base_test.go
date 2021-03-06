package no1

import (
	"ACD/common"
	"flag"
	"testing"
)

var (
	data      = []string{"bbbab", "cbbd", "xxxxx22xxxs4aossd2xxxxx"}
	expectAns = []int{4, 2, 15}
	log       = *common.GetLogger()
)

// TestByBuiltinData 使用内置数据完成测试
func TestByBuiltinData(t *testing.T) {

	for idx, val := range data {
		solution := Solution(val)
		expect := expectAns[idx]
		status := solution == expect
		if status {
			log.Infof("\n\t\t输入: \"%v\" \n\t\t输出: \t%v \n\t\t预期: \t%v\n", val, solution, expect)
		} else {
			log.Errorf("\n\t\t输入: \"%v\" \n\t\t输出: \t%v \n\t\t预期: \t%v\n", val, solution, expect)
		}
	}
}

// TestByStdin 使用命令行输入
func TestByStdin(t *testing.T) {
	flag.Parse()
	args := flag.Args()
	for _, input := range args {
		if input != "" {
			solution := Solution(input)
			log.Infof("\n\t\t输入: \"%v\" \n\t\t输出: \t%v ", input, solution)
		}
	}
}
