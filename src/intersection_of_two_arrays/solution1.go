package main

func intersection(nums1 []int, nums2 []int) []int {
    // when value is 0, key doesn't exisst in nums1.
    // when value is 1, key exists in nums1.
    // when value is 2, key exists in nums2.
    nums1ValuesSet := map[int]int{}

    for _, value := range nums1 {
        nums1ValuesSet[value] = 1
    }

    for _, value := range nums2 {
        if _, ok := nums1ValuesSet[value]; ok {
            nums1ValuesSet[value] = 2
        }
    }

    intersects := make([]int, 0)

    for key, value := range nums1ValuesSet {
        if value == 2 {
            intersects = append(intersects, key)
        }
    }

    return intersects
}

// 時間計算量 O(n + m)
// 空間計算量 O(n)