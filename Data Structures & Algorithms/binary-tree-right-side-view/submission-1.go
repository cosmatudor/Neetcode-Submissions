/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	mrn := make(map[int]int)
	queue := []*TreeNode{}

	//bfs
	queue = append(queue, root)
	level := 0
	for len(queue) != 0 {
		n := len(queue)

		for i := 0; i < n; i ++ {
			node := queue[0]
			queue = queue[1:]
			mrn[level] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	fmt.Println(len(mrn))

	// compute result
	res := []int{}
	for i := 0; i < len(mrn); i++ {
		res = append(res, mrn[i])
	}

	return res
}
