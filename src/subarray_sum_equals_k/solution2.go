package main

func subarraySum2(nums []int, k int) int {
    sumArray := make([]int, 1, len(nums) + 1)
    sumArray[0] = 0

    for i, num := range nums {
        value :=  sumArray[i] + num
        sumArray = append(sumArray, value)
    }

    pairCount := 0

    sumArrayLength := len(sumArray)
    for startIndex := 0; startIndex < (sumArrayLength - 1); startIndex++ {
        for endIndex := startIndex + 1; endIndex < sumArrayLength; endIndex++ {
            if sumArray[endIndex] -  sumArray[startIndex] == k {
                pairCount += 1
            }
        }
    }
    return pairCount
}