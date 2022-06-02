package main

import . "../../common_utils"

func main() {

}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}
	var ret = inorderSuccessor(root.Left, p)

	if ret == nil {
		return root
	}
	return ret
}
