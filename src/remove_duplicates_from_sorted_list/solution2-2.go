package main

func deleteDuplicateRecursively22(current *ListNode) {
    if current.Next == nil {
        return
    }

    if current.Val == current.Next.Val {
        current.Next = current.Next.Next
    } else {
        current = current.Next
    }
}

func deleteDuplicates22(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    deleteDuplicateRecursively22(head)
    return head
}

// 時間計算量: O(n)
// 空間計算量: O(1)