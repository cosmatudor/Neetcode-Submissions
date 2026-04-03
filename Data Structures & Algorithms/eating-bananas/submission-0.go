func minEatingSpeed(piles []int, h int) int {
	biggestPile := 0
	for _, pile := range piles {
		biggestPile = max(biggestPile, pile)
	}

	left, right := 1, biggestPile
	for left < right {
		mid := (left + right) / 2
		
		sum := 0
		for _, pile := range piles {
			sum += pile / mid
			if pile % mid != 0 {
				sum++
			}
		}

		if sum <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
