package main

func deleteDuplicates12(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
	current := head
    
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
 	}
	return head
}

// 時間計算量: O(n)
// 空間計算量: O(1)