package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

type Element struct {
	val   int
	times int
}

type PriorityQueue []Element

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].times < pq[j].times
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	(*pq) = append((*pq), x.(Element))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq
	n := len((*pq))
	ret := (*old)[n-1]
	(*pq) = (*pq)[:n-1]
	return ret
}

func topKFrequent(nums []int, k int) []int {
	var set = map[int]int{}

	for _, v := range nums {
		set[v]++
	}

	pq := new(PriorityQueue)
	heap.Push(pq, Element{val: -1, times: 0})
	for key, v := range set {
		var ele = Element{val: key, times: v}
		if len(*pq) < k {
			heap.Push(pq, ele)
		} else {
			var temp = heap.Pop(pq).(Element)
			if ele.times > temp.times {
				temp = ele
			}
			heap.Push(pq, temp)
		}
	}

	var result = []int{}
	for _, v := range *pq {
		result = append(result, v.val)
	}

	return result
}
