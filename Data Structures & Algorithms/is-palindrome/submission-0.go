func isPalindrome(s string) bool {
    lowerS := strings.ToLower(s)

    i := 0
    j := len(lowerS) - 1
    for i <= j {
        if (lowerS[i] < 'a' || lowerS[i] > 'z') && (lowerS[i] < '0' || lowerS[i] > '9') {
            i += 1
        }
        if (lowerS[j] < 'a' || lowerS[j] > 'z') && (lowerS[j] < '0' || lowerS[j] > '9') {
            j -= 1
        }

        if lowerS[i] != lowerS[j] {
            return false
        }

        i += 1
        j -= 1
    }

    return true
}
