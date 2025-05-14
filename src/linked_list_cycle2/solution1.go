package main

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return nil
    }
    fast := head.Next.Next
    slow := head.Next

    for fast != slow {
        if fast.Next == nil || fast.Next.Next == nil {
            return nil
        }

        fast = fast.Next.Next
        slow = slow.Next
    }

    hitNode := fast
    startNode := head

    for hitNode != startNode {
        hitNode = hitNode.Next
        startNode = startNode.Next
    }
    return hitNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)