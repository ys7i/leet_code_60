package main

type ListNode struct {
    Val int
    Next *ListNode
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode { 0, nil }
    current := dummyHead
    isCarryOver := false

    for l1 != nil || l2 != nil || isCarryOver {
        value1 := 0
        if l1 != nil {
            value1 = l1.Val
            l1 = l1.Next
        }

        value2 := 0
        if l2 != nil {
            value2 = l2.Val
            l2 = l2.Next
        }

        currentValue := value1 + value2
        if isCarryOver {
            currentValue += 1
            isCarryOver = false
        }

        if currentValue >= 10 {
            currentValue -= 10
            isCarryOver = true
        }

        current.Next = &ListNode { currentValue, nil }
        current = current.Next
    }
    return dummyHead.Next
}