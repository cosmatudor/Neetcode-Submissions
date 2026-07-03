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
	levels := make(map[*TreeNode]int)
	queue := []*TreeNode{}

	//bfs
	queue = append(queue, root)
	mrn[1] = root.Val
	levels[root] = 1
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		prevLevel := levels[node]

		if node.Left != nil {
			queue = append(queue, node.Left)
			levels[node.Left] = prevLevel + 1
			mrn[prevLevel + 1] = node.Left.Val
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			levels[node.Right] = prevLevel + 1
			mrn[prevLevel + 1] = node.Right.Val
		}
	}

	fmt.Println(len(mrn))

	// compute result
	res := []int{}
	for i := 1; i <= len(mrn); i++ {
		res = append(res, mrn[i])
	}

	return res
}
