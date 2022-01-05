package no_09

import "container/heap"

func Solution(apples []int, days []int) int {
	// 遍历一遍，手上苹果能吃就直接吃，当天吃过苹果以后，去捡新的苹果，维护优先队列，尽量明天吃的时候尽量吃保质期短的
	h := hp{}
	n, time, result := len(apples), 0, 0
	for ; time < n || len(h) > 0; time++ {
		// 清理堆内过期食物
		for len(h) > 0 && h[0].life <= time {
			heap.Pop(&h)
		}

		// 上新
		if time < n && days[time] > 0 && apples[time] > 0 {
			heap.Push(&h, apple{days[time] + time, apples[time]})
		}

		// 取出堆顶苹果，开恰！
		if len(h) > 0 {
			result++

			if h[0].cnt == 1 {
				heap.Pop(&h)
			} else {
				h[0].cnt--
			}
		}
	}

	return result
}

// 苹果 >> 保质期 + 个数
type apple struct{ life, cnt int }

type hp []apple

func (h hp) Len() int { return len(h) }

// 比较器，队列排序依据
// a < b ｜ a 在前
func (h hp) Less(a, b int) bool { return h[a].life < h[b].life }

// 交换
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(apple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
