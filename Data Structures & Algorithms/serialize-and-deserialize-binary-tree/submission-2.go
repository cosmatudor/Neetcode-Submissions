/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    data string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ser := ""
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			ser += strconv.Itoa(root.Val) + "|"
			dfs(root.Left)
			dfs(root.Right)
		} else {
			ser += "n" + "|"
		}
	}
	
	dfs(root)
	return ser
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    parts := strings.Split(strings.TrimRight(data, "|"), "|")

	i := 0
	parent := TreeNode{}
	var dfs func(parent *TreeNode, isLeft bool)
	dfs = func(parent *TreeNode, isLeft bool) {
		if i >= len(parts) {
			return 
		}

		if parts[i] == "n" {
			return
		}

		x, _ := strconv.Atoi(parts[i])
		node := &TreeNode{Val: x}

		if isLeft {
			parent.Left = node
		} else {
			parent.Right = node
		}

		i++
		dfs(node, true)
		i++
		dfs(node, false)
	}

	dfs(&parent, true)

	return parent.Left
}
