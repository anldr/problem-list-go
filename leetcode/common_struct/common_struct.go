package common_struct

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Min(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

func Max(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}
