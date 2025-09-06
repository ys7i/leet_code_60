package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueElement struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	queue := []QueueElement{QueueElement{root, 0}}

	maxDepth := 0
	for len(queue) > 0 {
		elm := queue[0]
		queue = queue[1:]

		node := elm.node
		if node == nil {
			maxDepth = max(maxDepth, elm.depth)
			continue
		}

		queue = append(queue, QueueElement{node.Left, elm.depth + 1}, QueueElement{node.Right, elm.depth + 1})
	}

	return maxDepth
}
