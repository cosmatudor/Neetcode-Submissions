/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	isBalanced := true
    var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if -1 > left - right || left - right > 1 {
			isBalanced =  false
		}

		return 1 + max(left, right)
	}

	dfs(root)
	return isBalanced
}
