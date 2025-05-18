package main

 type ListNode struct {
     Val int
     Next *ListNode
 }

 func getVisited(head *ListNode) map[int]int {
    visited := make(map[int]int)
    current := head
    
    for current != nil {
        visited[current.Val] += 1
        current = current.Next
    }

    return visited
}

func deleteDuplicates1(head *ListNode) *ListNode {
    visited := getVisited(head)
    dummyHead := ListNode{ -1, nil }
    dummyHead.Next = head
    current := &dummyHead

    for current.Next != nil {
        if visited[current.Next.Val] > 1 {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }

    return dummyHead.Next
}

// 時間計算量: O(n)
// 空間計算量: O(n)