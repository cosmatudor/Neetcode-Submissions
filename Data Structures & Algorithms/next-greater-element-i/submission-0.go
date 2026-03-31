func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreats := make(map[int]int)
	s := [][2]int{}

	for i, num := range nums2 {
		for len(s) > 0 && nums2[i] >= s[len(s)-1][0] {
			top := s[len(s)-1]
			s = s[:len(s)-1]

			nextGreats[nums2[top[1]]] = num
		}
		s = append(s, [2]int{num, i})
	}

	var res []int
	for _, x := range nums1 {
		el := -1
		if val, exists := nextGreats[x]; exists {
			el = val
		}
		
		res = append(res, el)
	}

	return res
}
