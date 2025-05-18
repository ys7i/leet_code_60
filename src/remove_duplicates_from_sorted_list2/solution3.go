package main

func getVisited3(head *ListNode) map[int]int {
    visited := make(map[int]int)
    current := head
    
    for current != nil {
        visited[current.Val] += 1
        current = current.Next
    }

    return visited
}

func skipNode(node *ListNode, skipCount int) {
    for i:=0; i<skipCount; i++ {
        node.Next = node.Next.Next
    }
}

func deleteDuplicates3(head *ListNode) *ListNode {
    visited := getVisited3(head)
    dummyHead := ListNode{ -1, nil }
    dummyHead.Next = head
    current := &dummyHead

    for current.Next != nil {
        if visited[current.Next.Val] > 1 {
            skipNode(current, visited[current.Next.Val])
        } else {
            current = current.Next
        }
    }

    return dummyHead.Next
}

// 時間計算量: O(n)
// 空間計算量: O(n)