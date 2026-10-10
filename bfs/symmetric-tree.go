package bfs

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func isSymmetric(root *TreeNode[int]) bool {

	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)

}

func isMirror(p *TreeNode[int], q *TreeNode[int]) bool {
	// If both nodes are nil, they are structurally symmetric
	if p == nil && q == nil {
		return true
	}
	// If only one node is nil, or values don't match, they are not symmetric
	if p == nil || q == nil || p.Data != q.Data {
		return false
	}

	// Check inner and outer subtrees:
	// 1. Left child of p matches Right child of q
	// 2. Right child of p matches Left child of q
	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}
