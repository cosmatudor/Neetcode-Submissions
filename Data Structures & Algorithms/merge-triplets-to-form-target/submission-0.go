func mergeTriplets(triplets [][]int, target []int) bool {
	start := []int{1,1,1}
	for _, triplet := range triplets {
		shouldContinue := false
		for i := 0; i < 3; i++ {
			if triplet[i] > target[i] {
				shouldContinue = true
			}
		}
		if shouldContinue {
			continue
		}

		cnt := 0
		for i := 0; i < 3; i++ {
			start[i] = max(start[i], triplet[i])
			if start[i] == target[i] {
				cnt++
			}
		} 
		if cnt == 3 {
			return true
		}
	}

	return false
}
