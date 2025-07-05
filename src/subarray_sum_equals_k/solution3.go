package main

func subarraySum3(nums []int, k int) int {
    sumToOccurrence := make(map[int]int, len(nums))
    sumToOccurrence[0] = 1
    pairCount := 0
    sum := 0

    for _, num := range nums {
        sum += num

        occurrence := sumToOccurrence[sum - k]

        pairCount += occurrence
        sumToOccurrence[sum] += 1
    }

    return pairCount
}