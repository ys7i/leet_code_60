package main

func hasCycle2(head *ListNode) bool {
    hashMap := make(map[*ListNode]struct{})

    for node := head; node != nil; node = node.Next {
        if _, ok := hashMap[node]; ok {
            return true
        }
        hashMap[node] = struct{}{}
    }
    return false
}

// When the lenth of list is n,
// Time Complexity: O(n)
// Space Complexity: O(n) (Node.Val has large range, so we can say the length of nodeMap value is constant)
// I spent about 10m.
