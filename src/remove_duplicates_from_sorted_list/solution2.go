package main

func deleteDuplicateRecursively(preNode *ListNode, node *ListNode) {
    if node == nil {
        return
    }

    if preNode.Val == node.Val {
        preNode.Next = node.Next
    } else {
        preNode = preNode.Next
    }
    deleteDuplicateRecursively(preNode, node.Next)
}

func deleteDuplicates2(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    deleteDuplicateRecursively(head, head.Next)
    return head
}

// 時間計算量: O(n)
// 空間計算量: O(1)