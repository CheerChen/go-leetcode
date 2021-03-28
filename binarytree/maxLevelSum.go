package binarytree

func maxLevelSum(root *TreeNode) int {
	level := 1
	maxVal := make(map[int]int)
	maxLevel(root, level, maxVal)
	maxSum := 0
	minLevel := len(maxVal)
	for l := len(maxVal); l > 0; l-- {
		if maxVal[l] >= maxSum {
			maxSum = maxVal[l]
			minLevel = l
		}
	}
	return minLevel
}

func maxLevel(root *TreeNode, level int, maxVal map[int]int) {
	maxVal[level] += root.Val
	if root.Left != nil {
		maxLevel(root.Left, level+1, maxVal)
	}
	if root.Right != nil {
		maxLevel(root.Right, level+1, maxVal)
	}
	return
}
