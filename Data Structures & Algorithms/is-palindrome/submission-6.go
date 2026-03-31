func isPalindrome(s string) bool {
    lowerS := strings.ToLower(s)

    i := 0
    j := len(lowerS) - 1
    for i <= j {
        for (lowerS[i] < 'a' || lowerS[i] > 'z') && (lowerS[i] < '0' || lowerS[i] > '9') && i < len(lowerS) - 1 {
            i += 1
        }
        for (lowerS[j] < 'a' || lowerS[j] > 'z') && (lowerS[j] < '0' || lowerS[j] > '9') && j > 0 {
            j -= 1
        }

        fmt.Println(i)
        fmt.Println(j)

        if i > j {
            return true
        }

        if lowerS[i] != lowerS[j] {
            fmt.Println(string(lowerS[i]))
            fmt.Println(string(lowerS[j]))
            return false
        }

        i += 1
        j -= 1
    }

    return true
}
