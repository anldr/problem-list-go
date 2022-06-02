package main

import (
	. "../common_utils"
	"fmt"
)

func main() {
	var node1 = &TreeNode{Val: 1}
	fmt.Println(averageOfSubtree(node1))
}

func averageOfSubtree(root *TreeNode) int {
	var result = 0
	var dfs func(root *TreeNode) (int, int)

	dfs = func(root *TreeNode) (int, int) {
		if root != nil {
			left, leftnums := dfs(root.Left)
			right, rightnums := dfs(root.Right)

			var nums = leftnums + rightnums + 1
			var ret = root.Val + left + right
			if ret/nums == root.Val {
				result++
			}
			return root.Val + left + right, nums
		}
		return 0, 0
	}

	dfs(root)

	return result
}
