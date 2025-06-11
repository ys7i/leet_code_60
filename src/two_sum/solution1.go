package main

func twoSum1(nums []int, target int) []int {
    num_to_index := make(map[int]int, len(nums))

    for index, value := range nums {
        if other_index, ok := num_to_index[target - value]; ok {
            return []int{ other_index, index }
        }
        num_to_index[value] = index
    }
    // This is for type checking
    return []int{ 0, 0 }
}

// Time Complexity: O(n)
// Space Complexity: O(n)