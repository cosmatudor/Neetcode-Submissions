type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, s := range strs {
        sb.WriteString(strconv.Itoa(len(s)))
        sb.WriteString("#")
        sb.WriteString(s)
    }

    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    var decodedStr []string
    i := 0
    for i < len(encoded) {
        j := i
        nr := 0
        for encoded[j] != '#' {
            digit := encoded[j] - '0'
            nr = nr * 10 + int(digit)
            j += 1
        }
        decodedStr = append(decodedStr, encoded[j+1:j+nr+1])

        i = j+nr+1
    }

    return decodedStr
}

