func maxProfit(prices []int) int {
    min := 101
    max := 0
    for i := 0; i <= len(prices) - 1; i++ {
        if prices[i] - min > max {
            max = prices[i] - min
        }

        if prices[i] < min {
            min = prices[i]
        }
    }

    return max
}
