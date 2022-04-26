package main

import (
	"fmt"
)

func main() {
	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 2, Left: node2, Right: node3}
	fmt.Println(isValidBST(node1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	nums := []int{}
	inorder(root, &nums)

	return isAscending(&nums)
}

func inorder(root *TreeNode, nums *[]int) {
	if root != nil {
		inorder(root.Left, nums)
		*nums = append(*nums, root.Val)
		inorder(root.Right, nums)
	}
}

func isAscending(nums *[]int) bool {
	arrLen := len(*nums)
	for i := arrLen - 1; i >= 1; i-- {
		if (*nums)[i] <= (*nums)[i-1] {
			return false
		}
	}

	return true
}
