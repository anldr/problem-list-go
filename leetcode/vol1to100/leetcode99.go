package main

import (
	. "../common_utils"
	"fmt"
)

func main() {
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1}
	node4 := TreeNode{Val: 4, Left: &node2}
	node3 := TreeNode{Val: 3, Left: &node1, Right: &node4}

	recoverTree(&node3)
	fmt.Println(node1.Val)
}

func recoverTree(root *TreeNode) {
	var x *TreeNode
	var y *TreeNode
	var preNode *TreeNode
	// binary tree morris traversal
	current := root
	for current != nil {
		mostRight := current.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != current {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = current
				current = current.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}

		if preNode != nil && preNode.Val >= current.Val {
			y = current
			if x == nil {
				x = preNode
			}
		}

		preNode = current
		current = current.Right
	}

	x.Val, y.Val = y.Val, x.Val
}
