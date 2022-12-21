package main

import (
	. "../common_utils"
	"fmt"
)

func main() {
	node2 := TreeNode{2, nil, nil}
	node_3 := TreeNode{-3, nil, nil}
	node5 := TreeNode{5, &node2, &node_3}

	fmt.Println(findFrequentTreeSum(&node5))
}

func findFrequentTreeSum(root *TreeNode) []int {
	result := []int{}
	set := map[int]int{}
	maxCount := 0
	var recursive func(root *TreeNode) int

	recursive = func(root *TreeNode) int {
		if root != nil {
			sum := 0

			sum += root.Val
			sum += recursive(root.Left)
			sum += recursive(root.Right)

			set[sum]++
			if set[sum] > maxCount {
				maxCount = set[sum]
			}

			return sum
		}
		return 0
	}

	recursive(root)

	for k, v := range set {
		if v == maxCount {
			result = append(result, k)
		}
	}

	return result
}
