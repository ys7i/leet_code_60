package main

func detectCycle(head *ListNode) *ListNode {
    hash := make(map[*ListNode]struct{})
    node := head

    for {
        if node == nil {
            return nil
        }

        if _, ok := hash[node]; ok {
            return node
        }

        hash[node] = struct{}{}
        node = node.Next
    }
}

// Time Complexity: O(n)
// Space Complexity: O(n)
// I spent 5m.