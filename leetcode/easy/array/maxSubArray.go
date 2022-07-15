package array

import "fmt"

func MaxSubArray(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}

	size := len(nums)
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	ret := nums[0]
	var max int
	_ = max
	//{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 1; i < size; i++ {
		if max < 0 {
			max = nums[i]
		} else {
			max = max + nums[i]
		}
		if ret < max {
			ret = max
		}
	}
	return ret
}

/*

目前有贪心和动态规划2个解法
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。、
示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions/xmi2l7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		fmt.Println(nums1)
	}
	for i := 0; i < n; i++ { //遍历第二个
		k := m - 1
		for k >= 0 {
			if nums1[k] >= nums2[i] {
				nums1[k+1] = nums1[k]
				k = k - 1
				continue
			}
			break
		}
		nums1[k+1] = nums2[i]
		m = m + 1
	}
	fmt.Println(nums1)
}
func TestMergeTwoArrays() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	nums1 = []int{2, 0}
	nums2 = []int{1}
	merge(nums1, 1, nums2, 1)
}
