func getFreqArrOfStr(s string) [26]int {
    var freq [26]int
    for _, x := range s {
        freq[x - 'a'] += 1
    }

    return freq
}

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }

    myMap := make(map[[26]int][]string)

    for _, x := range strs {
        freqOfX := getFreqArrOfStr(x)
        myMap[freqOfX] = append(myMap[freqOfX], x)
    }

    var res [][]string
    for _, value := range myMap {
        res = append(res, value)
    }

    return res
}
