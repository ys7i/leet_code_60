package main

func isPair(left rune, right rune) bool {
    if left == '(' && right == ')' {
        return true
    }
    if left == '{' && right == '}' {
        return true
    }
    if left == '[' && right == ']' {
        return true
    }
    return false
}

func isValidHelper(s string) (bool, string) {
    if len(s) == 0 {
        return true, ""
    }
    if len(s) == 1 {
        return false, ""
    }

    if s[0] == '}' || s[0] == ')' || s[0] == ']' {
        return false, ""
    }

    if isPair(rune(s[0]), rune(s[1])) {
        return true, s[2:]
    }

    firstChar := s[0]
    s = s[1:]
    for {
        valid, left := isValidHelper(s)
        if !valid || len(left) == 0 {
            return false, ""
        }
        if isPair(rune(firstChar), rune(left[0])) {
            return true, left[1:]
        }
        s = left
    }
}

func isValid2(s string) bool {
    validPair, left_s := isValidHelper("(" + s + ")")
    return validPair && len(left_s) == 0
}
