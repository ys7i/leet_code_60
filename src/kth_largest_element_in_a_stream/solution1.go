package main

import (
	"slices"
	"sort"
)

type KthLargest struct {
    k int
    values []int
}


func Constructor(k int, nums []int) KthLargest {
  sort.Slice(nums, func(i int, j int) bool { return nums[i] > nums[j] })
  if len(nums) >= k {
    nums = nums[0:k]
  }
  return KthLargest{ k, nums }
}

func (this *KthLargest) FindIndexOfValues(val int) int {
    for i, elm := range this.values {
        if elm < val {
            return i
        }
    }
    return -1
}

func (this *KthLargest) Add(val int) int {
    insertIndex := this.FindIndexOfValues(val)

    if insertIndex == -1 {
        if this.k <= len(this.values) {
            return this.values[this.k - 1]
        }

        this.values = append(this.values, val)
        return val
    }

    newValues := slices.Concat(this.values[0:insertIndex],[]int{ val }, this.values[insertIndex:])
    this.values = newValues[0:this.k]
    return this.values[this.k - 1]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

 // Addの時間計算量はO(n)