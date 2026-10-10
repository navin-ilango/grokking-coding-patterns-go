package dfs

// Definition of a binary tree node
type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode[int]) int {

	diameter := 0

	var dfs func(*TreeNode[int]) int

	dfs = func(tn *TreeNode[int]) int {

		if tn == nil {
			return 0
		}

		leftHeight := dfs(tn.Left)
		rightHeight := dfs(tn.Right)

		if leftHeight+rightHeight > diameter {
			diameter = leftHeight + rightHeight
		}

		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)

	return diameter

}
