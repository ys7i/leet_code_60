package main

func firstUniqChar(s string) int {
    charToIndex := make(map[rune]int, 26)

    for index, char := range s {
        if _, ok := charToIndex[char]; ok {
            charToIndex[char] = -1
        } else {
            charToIndex[char] = index
        }
    }

    targetIndex := -1

    for _, index := range charToIndex {
        if index == -1 {
            continue
        }

        if targetIndex == -1 || targetIndex > index {
            targetIndex = index
        }
    }

    return targetIndex
}

// 時間計算量 O(n)
// 空間計算量 O(1)