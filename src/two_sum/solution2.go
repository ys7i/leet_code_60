package main

func twoSumHelper(nums []int, currentIndex int, target int, numToIndex map[int]int) []int {
	if currentIndex >= len(nums) {
		return []int{ 0, 0 }
	}
	value := nums[currentIndex]
	if otherIndex, ok := numToIndex[target - value]; ok {
		return []int{ otherIndex, currentIndex }
	}
	numToIndex[value] = currentIndex
	return twoSumHelper(nums, currentIndex + 1, target, numToIndex)
}

func twoSum2(nums []int, target int) []int {
numToIndex := make(map[int]int, len(nums))

return twoSumHelper(nums, 0,target, numToIndex)
}