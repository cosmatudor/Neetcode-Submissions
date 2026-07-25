/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxPathSum := math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		maxPathSum = max(maxPathSum, left + right + node.Val)

		return max(0, max(left + node.Val, right + node.Val))
	}

	dfs(root)
	return maxPathSum
}
