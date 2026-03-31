/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        aux := root.Left
        root.Left = root.Right
        root.Right = aux

        invertTree(root.Left)
        invertTree(root.Right)
    }

    return root
}
