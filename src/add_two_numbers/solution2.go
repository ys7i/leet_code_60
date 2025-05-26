package main

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, carryOverValue int8) *ListNode {
    if l1 == nil && l2 == nil && carryOverValue == 0 {
        return nil
    }

    var value1, value2 int8
    if l1 != nil {
        value1 = int8(l1.Val)
        l1 = l1.Next
    }

    if l2 != nil {
        value2 = int8(l2.Val)
        l2 = l2.Next
    }

    total := int(value1) + int(value2) + int(carryOverValue)
    nextCarryOverValue := int8(total / 10)
    value := total % 10
    
    return &ListNode { value, addTwoNumbersHelper(l1, l2, nextCarryOverValue) }
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
    return addTwoNumbersHelper(l1, l2, 0)
}

// 時間計算量 O(n)
// 空間計算量 O(n)