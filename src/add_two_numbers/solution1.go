package main

type ListNode struct {
    Val int
    Next *ListNode
}

func moveNextDigit(l *ListNode) (int, *ListNode) {
    if l == nil {
        return 0, nil
    }
    return l.Val, l.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode { 0, nil }
    current := dummyHead
    isCarryOver := false

    for l1 != nil || l2 != nil || isCarryOver {
        var value1, value2 int
        value1, l1 = moveNextDigit(l1)
        value2, l2 = moveNextDigit(l2)

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

