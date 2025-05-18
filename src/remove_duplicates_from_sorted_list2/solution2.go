package main

func skipTargetNode(node *ListNode, targetVal int) {
    for node.Next != nil && node.Next.Val == targetVal {
        node.Next = node.Next.Next
    }
}

func deleteDuplicates2(head *ListNode) *ListNode {
    dummyNode := ListNode { -1, head }
    current := &dummyNode

    for current.Next != nil && current.Next.Next != nil {
        if current.Next.Val == current.Next.Next.Val {
            skipTargetNode(current, current.Next.Val)
        } else {
            current = current.Next
        }
    }
    return dummyNode.Next
}