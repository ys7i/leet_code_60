package main

func isValid1(s string) bool {
    stack := make([]rune, 0, len(s)/2)

    pairs := map[rune]rune {
        '}': '{',
        ')': '(',
        ']': '[',
    }

    for _, char := range s { 
        switch (char) {
            case '{', '(', '[':
                stack = append(stack, char)
            default:
                lastInsertedIndex := len(stack) - 1
                if lastInsertedIndex < 0 {
                    return false
                }

                pair := pairs[char]
                if stack[lastInsertedIndex] != pair {
                    return false
                }
                stack = stack[:lastInsertedIndex]
        }
    }
    return len(stack) == 0
}

// 時間計算量 O(n)
// (最悪)空間計算量 O(n)