package simple

func HasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	queNode := make([]*TreeNode, 0)
	queVal := make([]int, 0)
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		if now.Left == nil && now.Right == nil {
			if temp == target {
				return true
			}
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
}

// HasPathSumOpt 递归
func HasPathSumOpt(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return target == root.Val
	}
	return HasPathSumOpt(root.Left, target-root.Val) || HasPathSumOpt(root.Right, target-root.Val)
}
