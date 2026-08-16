/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
        return true
    }
    if root == nil {
        return false
    }
    
    return checkSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func checkSameTree(root1, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }

    if root1 == nil || root2 == nil || root1.Val != root2.Val {
        return false
    }

    return checkSameTree(root1.Left, root2.Left) && checkSameTree(root1.Right, root2.Right)
}