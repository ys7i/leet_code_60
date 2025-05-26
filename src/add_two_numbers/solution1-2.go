package main

func addTwoNumbers12(l1 *ListNode, l2 *ListNode) *ListNode {
    sentinel := &ListNode { 0, nil }
    currentNode := sentinel
    carryOverValue := 0

    for l1 != nil || l2 != nil || carryOverValue > 0 {
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
        currentValue += carryOverValue

        carryOverValue = currentValue / 10
        currentValue -= carryOverValue * 10

        currentNode.Next = &ListNode { currentValue, nil }
        currentNode = currentNode.Next
    }
    return sentinel.Next
}

// 時間計算量 O(n)
// 空間計算量 O(n)