/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil {
		return root.Right.Val > root.Val
	}

	if root.Right == nil {
		return root.Left.Val < root.Val
	}

	ok := (root.Left.Val < root.Val) && (root.Right.Val > root.Val)
	return ok && isValidBST(root.Left) && isValidBST(root.Right)
}
