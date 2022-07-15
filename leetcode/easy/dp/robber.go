package dp

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int) int {
	size := len(nums)
	if 0 == size {
		return 0
	}
	/*
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}*/
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 0
	}

	for i := 0; i < size; i++ {
		if i == 0 {
			dp[0] = nums[0]
		}
		if i == 1 {
			dp[i] = max(dp[0], nums[1])
		}
		if i > 1 {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
	}
	return dp[size-1]

}

func robcircle(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob(nums[:n-1]), rob(nums[1:]))
}
