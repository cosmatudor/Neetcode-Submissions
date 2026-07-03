/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    res := 0

	var dfs func(prev, curr *TreeNode)
	dfs = func(prev, curr *TreeNode) {
		if curr == nil {
			return
		}

		if prev == nil {
			res++
		} else {
			if prev.Val > curr.Val {
				curr.Val = prev.Val
			} else {
				res++
			}
		}

		dfs(curr, curr.Left)
		dfs(curr, curr.Right)
	}

	prev := TreeNode{
		Val: -101,
	}
	dfs(&prev, root)

	return res
}
