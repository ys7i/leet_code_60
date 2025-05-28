package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    var prevNode *ListNode = nil
    node := head

    for {
        if node.Next == nil {
            node.Next = prevNode
            break
        }

        next := node.Next
        node.Next = prevNode

        prevNode = node
        node = next
    }
    return node
}

// 初回実装時間 9m16s
// 時間計算量O(n), 空間計算量O(1)