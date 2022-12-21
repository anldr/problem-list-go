package main

import (
	. "../common_utils"
	"fmt"
)

func main() {
	node7 := TreeNode{Val: 7}
	node6 := TreeNode{Val: 6}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5, Left: &node7}
	node2 := TreeNode{Val: 2, Left: &node4}
	node3 := TreeNode{Val: 3, Left: &node5, Right: &node6}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	fmt.Println(findBottomLeftValue(&node1))
}

func findBottomLeftValue(root *TreeNode) int {
	_, result := recursive(root, 0)
	return result
}

func recursive(root *TreeNode, dep int) (int, int) {
	if root != nil {
		leftDep, leftVal := recursive(root.Left, dep+1)
		rightDep, rightVal := recursive(root.Right, dep+1)

		if leftDep == rightDep {
			if root.Left == nil {
				return leftDep, root.Val
			} else {
				return leftDep, leftVal
			}
		} else if leftDep > rightDep {
			return leftDep, leftVal
		} else {
			return rightDep, rightVal
		}
	}
	return dep - 1, 0
}
