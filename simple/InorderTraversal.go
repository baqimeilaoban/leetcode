package simple

func InorderTraversal(root *TreeNode) []int {
	var res []int
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		res = append(res, node.Val)
		order(node.Right)
	}
	order(root)
	return res
}
