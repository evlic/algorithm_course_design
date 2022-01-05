package no_06

import (
	"ACD/common"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) String() string {
	if root == nil {
		return "null,"
	}
	return strconv.Itoa(root.Val) + "," + root.Left.String() + root.Right.String()
}

// Solution >> max 左侧为左子树，右侧为右子树
func Solution(n []int) *TreeNode {
	if len(n) == 0 {
		return nil
	}
	maxIdx := common.MaxInSlice(n)
	root := &TreeNode{
		Val:   n[maxIdx],
		Left:  Solution(n[:maxIdx]),
		Right: Solution(n[maxIdx+1:]),
	}
	return root
}
