package main

import (
	. "../common_utils"
	"fmt"
)

func main() {
	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 2, Left: node2, Right: node3}
	fmt.Println(isValidBST(node1))
}

func isValidBST(root *TreeNode) bool {
	//nums := []int{}
	//inorder(root, &nums)
	//
	//return isAscending(&nums)
	return inorder_2(root)
}

func inorder_2(root *TreeNode) bool {
	if root != nil {
		if !inorder_2(root.Left) {
			return false
		}

		if root.Left != nil && root.Right != nil && !(root.Left.Val < root.Val && root.Val < root.Right.Val) {
			return false
		}
		if root.Left == nil && root.Right != nil && !(root.Val < root.Right.Val) {
			return false
		}
		if root.Left != nil && root.Right == nil && !(root.Left.Val < root.Val) {
			return false
		}

		if !inorder_2(root.Right) {
			return false
		}
	}

	return true
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
