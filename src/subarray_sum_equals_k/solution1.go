package main

func subarraySum1(nums []int, k int) int {
    pairCount := 0

    for startIndex, _ := range nums {
        sum := 0

        for _, value := range nums[startIndex:] {
            sum += value

            if sum == k {
                pairCount += 1
            }
        }
    }

    return pairCount
}