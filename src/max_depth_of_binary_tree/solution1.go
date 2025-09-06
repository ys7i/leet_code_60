package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return maxDepthRecursively(root, 0)
}

func maxDepthRecursively(node *TreeNode, maxDepth int) int {
	if node == nil {
		return maxDepth
	}

	maxDepth += 1

	leftDepth := maxDepthRecursively(node.Left, maxDepth)
	rightDepth := maxDepthRecursively(node.Right, maxDepth)

	return max(maxDepth, leftDepth, rightDepth)
}

// Let's say the number of nodes is n.
// Time Complexity: O(n)
// Space Complexity: O(n)
