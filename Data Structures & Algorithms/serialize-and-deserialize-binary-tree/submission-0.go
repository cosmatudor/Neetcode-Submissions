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
	fmt.Println("SER:", ser)
	return ser
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    treeArr := this.strToArr(data)

	i := 1
	parent := TreeNode{}
	var dfs func(parent *TreeNode, isLeft bool)
	dfs = func(parent *TreeNode, isLeft bool) {
		if i - 1 >= len(treeArr) {
			return 
		}

		if treeArr[i-1] == -1001 {
			return
		}

		node := &TreeNode{
			Val: treeArr[i-1],
			Left: nil,
			Right: nil,
		}

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

func (this *Codec) strToArr(data string) []int {
	res := []int{}
	parts := strings.Split(strings.TrimRight(data, "|"), "|")
	for _, part := range parts {
		if part == "n" {
			res = append(res, -1001)
		} else {
			n, _ := strconv.Atoi(part)
			res = append(res, n)
		}
	}

	fmt.Println(res)
	return res
}