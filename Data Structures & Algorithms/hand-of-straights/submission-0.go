func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	freq := make(map[int]int)
	for _, x := range hand {
		freq[x]++
	}

	for _, card := range hand {
		if freq[card] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			if freq[card+i] == 0 {
				return false
			}
			freq[card+i]--
		}
	}

	return true
}