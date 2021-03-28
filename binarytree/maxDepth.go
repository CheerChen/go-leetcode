package binarytree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}
