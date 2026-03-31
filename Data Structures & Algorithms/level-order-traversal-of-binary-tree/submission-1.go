/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

    queue := []*TreeNode{}
	queue = append(queue, root)
	levels := make(map[*TreeNode]int)
	levels[root] = 1
	res := make(map[int][]int)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		level := levels[node]
		res[level] = append(res[level], node.Val)
		
		if node.Left != nil {
			queue = append(queue, node.Left)
			levels[node.Left] = level + 1
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			levels[node.Right] = level + 1
		}
	}

	levelNo := len(res)
	ans := [][]int{}
	for i := 1; i <= levelNo; i++ {
		ans = append(ans, res[i])
	}

	return ans
}
