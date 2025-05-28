package main

func reverseHeadTail(node *ListNode, processed *ListNode) *ListNode {
    if node == nil {
        return processed
    }
    next := node.Next
    node.Next = processed
    return reverseHeadTail(next, node)
}

func reverseList(head *ListNode) *ListNode {
    return reverseHeadTail(head, nil)
}