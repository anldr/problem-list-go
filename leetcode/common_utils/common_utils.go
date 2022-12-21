package common_utils

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

func MinInt32(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

func MaxInt32(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}

func MinInt64(i, j int64) int64 {
	return int64(math.Min(float64(i), float64(j)))
}

func MaxInt64(i, j int64) int64 {
	return int64(math.Max(float64(i), float64(j)))
}

func AbsInt32(i int) int {
	return int(math.Abs(float64(i)))
}

func AbsInt64(i int64) int64 {
	return int64(int(math.Abs(float64(i))))
}
