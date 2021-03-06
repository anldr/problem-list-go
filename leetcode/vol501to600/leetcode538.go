package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var sum = 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum = sum + root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}

	dfs(root)

	return root
}
