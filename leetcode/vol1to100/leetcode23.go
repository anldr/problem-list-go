package main

import "container/heap"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	var dummy = &ListNode{}
	preNode := dummy

	pq := new(PriorityQueue)
	pushKNode(pq, lists)
	for len(*pq) != 0 {
		temp := heap.Pop(pq).(*ListNode)
		preNode.Next = temp
		preNode = temp
		if temp.Next != nil {
			heap.Push(pq, temp.Next)
		}
	}

	return dummy.Next
}

func pushKNode(pq *PriorityQueue, lists []*ListNode) {
	var idx = 0
	for _, v := range lists {
		if v != nil {
			heap.Push(pq, v)
			idx++
		}
	}
}
