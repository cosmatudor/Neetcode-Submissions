/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func divideAndConquer (hashMap map[int]int, preorder *[]int, inorder []int) *TreeNode {
		if len(inorder) == 1 {
			return &TreeNode{Val: inorder[0]}
		}

		headIndex := hashMap[(*preorder)[0]]
		(*preorder) = (*preorder)[1:]

		var leftChild *TreeNode
		if headIndex > 0 {
			leftChild = divideAndConquer(hashMap, preorder, inorder[:headIndex])
		}

		var rightChild *TreeNode
		if headIndex+1 < len(inorder) {
			rightChild = divideAndConquer(hashMap, preorder, inorder[headIndex+1:])
		}

		return &TreeNode{
			Val: inorder[headIndex],
			Left: leftChild,
			Right: rightChild,
		}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
		return nil
	}

	hashMap := make(map[int]int)
	for i, x := range inorder {
		hashMap[x] = i
	}

	return divideAndConquer(hashMap, preorder, inorder)
}
