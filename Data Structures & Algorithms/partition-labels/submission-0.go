func partitionLabels(s string) []int {
	letterPos := make(map[byte][]int)
	res := []int{}

	for i, ch := range s {
		if _, ok := letterPos[byte(ch)]; !ok {
			letterPos[byte(ch)] = []int{i, i}
		} else {
			letterPos[byte(ch)][1] = i
		}
	}

	n := len(s)
	start := 0
	end := 0
	i := start
	for i <= min(end, n-1) {
		lastOccur := letterPos[s[i]][1]
		end = max(end, lastOccur)

		if i == end {
			res = append(res, end - start + 1)
			start, end = i+1, i+1
		}

		i++
	}

	return res
}
