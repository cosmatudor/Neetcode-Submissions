/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxVal(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}

	return maxVal(root.Right)
}

func minVal(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}

	return minVal(root.Left)
}

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

	ok1 := (root.Left.Val < root.Val) && (root.Right.Val > root.Val)
	ok2 := minVal(root.Right) > root.Val
	ok3 := maxVal(root.Left) < root.Val
	return ok1 && ok2 && ok3 && isValidBST(root.Left) && isValidBST(root.Right)
}
