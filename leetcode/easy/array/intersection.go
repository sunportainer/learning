package array

func intersect(nums1 []int, nums2 []int) []int {
	ans := []int{}
	hash := make(map[int]int)
	for _, val := range nums1 {
		if _, ok := hash[val]; ok {
			hash[val] = hash[val] + 1
		} else {
			hash[val] = 1
		}
	}
	for _, val := range nums2 {
		if _, ok := hash[val]; ok {
			if hash[val] > 0 {
				ans = append(ans, val)
				hash[val] = hash[val] - 1
			}

		}
	}
	return ans

}
