package main

type ListNode struct {
    Val int
    Next *ListNode
}


func deleteDuplicates1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    preNode := head
    node := head.Next

    for node != nil {
        if preNode.Val == node.Val {
            preNode.Next = node.Next
        } else {
            preNode = preNode.Next
        }
        node = node.Next
    }

    return head
}

// 時間計算量: O(n)
// 空間計算量: O(1)
