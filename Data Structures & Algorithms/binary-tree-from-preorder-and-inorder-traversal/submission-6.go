func divideAndConquer(hashMap map[int]int, preorder []int, inorder []int, start int, end int) *TreeNode {
    if start > end {
        return nil
    }

    rootVal := preorder[0]
    preorder = preorder[1:]

    rootIndexInInorder := hashMap[rootVal]

    node := &TreeNode{Val: rootVal}

    node.Left = divideAndConquer(hashMap, preorder, inorder, start, rootIndexInInorder-1)
    node.Right = divideAndConquer(hashMap, preorder, inorder, rootIndexInInorder+1, end)

    return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    hashMap := make(map[int]int)
    for i, x := range inorder {
        hashMap[x] = i
    }

    return divideAndConquer(hashMap, preorder, inorder, 0, len(inorder)-1)
}