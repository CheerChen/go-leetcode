package binarytree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	cmpLeft := isSameTree(p.Left, q.Left)
	cmpRight := isSameTree(p.Right, q.Right)
	return cmpRight && cmpLeft
}
