func maxProfit(prices []int) int {
    i := 0
    j := i+1
    maxProfit := 0
    for j < len(prices) {
        maxProfit = max(maxProfit, prices[j]-prices[i])
        if prices[j] < prices[i] {
            i = j
        }
        j++
    }
    return maxProfit
}
