func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}

	var dfs func(i int, curr []int, total int) 
	dfs = func(i int, curr []int, total int) {
		if total == target {
			sol := append([]int(nil), curr...)
			res = append(res, sol)
			return
		}

		if total > target || i >= len(nums) {
			return 
		}

		curr = append(curr, nums[i])
		dfs(i, curr, total + nums[i])
		curr = curr[:len(curr)-1]
		dfs(i+1, curr, total)
	}

	dfs(0, []int{}, 0)
	return res
}
