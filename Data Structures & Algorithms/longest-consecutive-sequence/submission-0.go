func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    longest := 0

    for num := range numSet {
        if _, hasPrev := numSet[num-1]; !hasPrev {
            currentNum := num
            currentStreak := 1

            for {
                if _, hasNext := numSet[currentNum+1]; hasNext {
                    currentNum++
                    currentStreak++
                } else {
                    break
                }
            }

            if currentStreak > longest {
                longest = currentStreak
            }
        }
    }

    return longest
}