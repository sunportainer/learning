package dp

func lengthOfLIS(nums []int) int {
	size := len(nums)
	//首先定义一个dp
	dp := make([]int, size)
	dp[0] = 1

	for i := 1; i < size; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j] > max {
					max = dp[j]
				}
			}
		}
		dp[i] = max + 1
	}
	ret := 0
	for _, v := range dp {
		if ret < v {
			ret = v
		}
	}
	return ret

}

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

作者：FennelDumplings
链接：https://leetcode.cn/leetbook/read/dynamic-programming-1-plus/5oqnpg/
*/

func findNumberOfLIS(nums []int) int {
	ans := 0
	max_len := 0
	size := len(nums)
	//首先定义一个dp
	dp := make([]int, size)
	count := make([]int, size)

	for i := 1; i < size; i++ {
		dp[i] = 1
		count[i] = 1
		for k := 0; k < i; k++ {
			if nums[k] < nums[i] {
				if dp[k]+1 == dp[i] {
					count[i] = count[i] + count[k]
				} else if dp[k]+1 > dp[i] {
					dp[i] = dp[k] + 1
					count[i] = count[k]
				}
			}
		}
		if dp[i] > max_len {
			max_len = dp[i]
			ans = count[i]
		} else if dp[i] == max_len {
			ans += count[i]
		}

	}

	return ans
}

/*给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。

*/
