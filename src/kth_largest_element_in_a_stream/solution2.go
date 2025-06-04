package main

import (
	"cmp"
	"slices"
)

func Constructor2(k int, nums []int) KthLargest {
    slices.SortFunc(nums, func(i int, j int) int { return cmp.Compare(i, j) })
    return KthLargest{k, nums}
}


func (this *KthLargest) Add2(val int) int {
    values := append(this.values, val)
    slices.SortFunc(values, func(i int, j int) int { return cmp.Compare(i, j) })
    values = values[0:this.k]
    this.values = values
    return values[this.k - 1]
}
