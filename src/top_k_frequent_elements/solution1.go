package main

import (
	"cmp"
	"slices"
)

func topKFrequent1(nums []int, k int) []int {
    countMap := make(map[int]int, len(nums))

    for _, v := range nums {
        countMap[v] += 1 
    }

    numFrequencies := make([][2]int, len(countMap))

    index := 0
    for key, v := range countMap {
        numFrequencies[index] = [2]int{ key, v }
        index += 1
    }

    slices.SortFunc(numFrequencies, func(a [2]int, b [2]int) int { return -1 * cmp.Compare(a[1], b[1]) })

    numFrequencies = numFrequencies[:k]
    frequentNums := make([]int, k)

    for i, pair := range numFrequencies {
        frequentNums[i] = pair[0]
    }
    return frequentNums
}