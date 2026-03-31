func numDecodings(s string) int {
    return dfs(s, 0, map[int]int{len(s): 1})
}

func dfs(s string, i int, dp map[int]int) int {
    if val, ok := dp[i]; ok {
        return val
    }
    if i == len(s) {
        return 1
    }
    if s[i] == '0' {
        return 0
    }
    res := dfs(s, i+1, dp)
    if i+1 < len(s) && (s[i] == '1' ||
       (s[i] == '2' && s[i+1] <= '6')) {
        res += dfs(s, i+2, dp)
    }
    dp[i] = res
    return res
}