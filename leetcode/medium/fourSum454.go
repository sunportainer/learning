package medium

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	map1 := make(map[int]int)
	for i := range nums1 {
		for k := range nums2 {
			ab := nums1[i] + nums2[k]
			if _, ok := map1[ab]; ok {
				map1[ab]++
				continue
			}
			map1[ab] = 1
		}
	}
	for i := range nums3 {
		for k := range nums4 {
			cd := 0 - (nums3[i] + nums4[k])
			if v, ok := map1[cd]; ok {
				ans = ans + v
			}

		}
	}
	return ans
}
