func divideAndConquer(hashMap map[int]int, preorder *[]int, inorder []int) *TreeNode {
    // 1. Cazul de bază trebuie să fie pentru slice gol, nu doar len == 1
    if len(inorder) == 0 {
        return nil
    }

    // 2. Extragem valoarea rădăcinii din "vârful" slice-ului global
    rootVal := (*preorder)[0]
    headIndex := hashMap[rootVal]
    
    // 3. Tăiem primul element din slice-ul global (modificăm referința)
    *preorder = (*preorder)[1:]

    // 4. Construim copiii (Inorder se împarte în funcție de headIndex)
    // Trimitem intervalele din inorder relative la poziția rădăcinii
    leftChild := divideAndConquer(hashMap, preorder, inorder[:headIndex])
    rightChild := divideAndConquer(hashMap, preorder, inorder[headIndex+1:])

    return &TreeNode{
        Val:   rootVal,
        Left:  leftChild,
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

    // Transmiți adresa lui preorder (&preorder)
    return divideAndConquer(hashMap, &preorder, inorder)
}