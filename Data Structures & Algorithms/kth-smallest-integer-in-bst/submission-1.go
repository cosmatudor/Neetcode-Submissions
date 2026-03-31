/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	createArrayFromNodes(root, &arr)
	return arr[k-1]
}

func createArrayFromNodes(root *TreeNode, arr *[]int) {
	if root != nil {
		createArrayFromNodes(root.Left, arr)
		*arr = append(*arr, root.Val)
		createArrayFromNodes(root.Right, arr)
	}
}
