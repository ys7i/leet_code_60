package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return mergeTreesRecursively(root1, root2)
}

func mergeTreesRecursively(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTreesRecursively(root1.Left, root2.Left),
		Right: mergeTreesRecursively(root1.Right, root2.Right),
	}
}

// Time Complexity: O(n)
// Space Complexity: O(n)
