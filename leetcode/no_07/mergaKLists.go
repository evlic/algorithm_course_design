package no_07

type ListNode struct {
	Val  int
	Next *ListNode
}

func (root *ListNode) Insert(a []int) {
	p := root
	for _, val := range a {
		p.Next = &ListNode{
			Val: val,
		}
		p = p.Next
	}
}

// 返回 较小 指针，拼接到上一个较小指针上，完成排序
// 返回 min(l0, l1), 并拼接
func mergeList(l0, l1 *ListNode) *ListNode {
	// 当某一个列表为空直接返回
	if l0 == nil {
		return l1
	}
	if l1 == nil {
		return l0
	}

	// l0 <+ l1
	// 维持升序
	var h *ListNode
	if l0.Val <= l1.Val {
		h = l0
		h.Next = mergeList(h.Next, l1)
	} else {
		// 如果 l1 小
		h = l1
		h.Next = mergeList(l0, h.Next)
	}

	return h
}

func Solution(ls []*ListNode) *ListNode {
	n := len(ls)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return ls[0]
	}
	if n == 2 {
		return mergeList(ls[0], ls[1])
	}
	// 分治
	mid := n >> 1
	// 归并
	return mergeList(Solution(ls[:mid]), Solution(ls[mid:]))
}
