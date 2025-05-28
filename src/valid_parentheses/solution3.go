package main

type RuneStack struct {
    values []rune
}

func (stack *RuneStack) size() int {
    return len(stack.values)
} 

func (stack *RuneStack) pop() rune {
    lastIndex := stack.size() - 1
    popped := stack.values[lastIndex]
    stack.values = stack.values[0:lastIndex]
    return popped
}

func (stack *RuneStack) push(value rune) {
    stack.values = append(stack.values, value)
}

func isValid3(s string) bool {
    stack := RuneStack{}

    closeToOpen := map[rune]rune {
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        switch char {
            case '(', '{', '[':
                stack.push(char)
            default:
                if stack.size() == 0 {
                    return false
                }
                popped := stack.pop()
                if closeToOpen[char] != popped {
                    return false
                }
        }
    }

    return stack.size() == 0
}
