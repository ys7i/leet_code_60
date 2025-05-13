package main

type ListNode struct {
  Val int
  Next *ListNode
}

func hasCycle1(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head.Next

    for slow != fast {
        slow = slow.Next
        if fast.Next == nil || fast.Next.Next == nil {
            return false
        }
        fast = fast.Next.Next
    }
    return true
}

// When the lenth of list is n,
// Time Complexity: O(n)
// Space Complexity: O(1)
// I spent 06m51s.
