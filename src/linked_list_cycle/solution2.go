package main

func hasCycle2(head *ListNode) bool {
    nodeMap := map[int][]*ListNode{}
    node := head

    for node != nil {
        list := nodeMap[node.Val]
        for _, s:= range list {
            if s == node {
                return true
            }
        }
        nodeMap[node.Val] = append(list, node)

        node = node.Next
    }

    return false
}

// When the lenth of list is n,
// Time Complexity: O(n)
// Space Complexity: O(n) (Node.Val has large range, so we can say the length of nodeMap value is constant)
// I spent about 10m.
