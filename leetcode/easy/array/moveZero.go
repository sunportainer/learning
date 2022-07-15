package array

import "fmt"

/*
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。https://leetcode.cn/leetbook/read/top-interview-questions/xmy9jh/
*/
func moveZeroes(nums []int) {
	if nil == nums || len(nums) == 0 {
		return
	}
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] == 0 {
			j--
		}
		for i < j && nums[i] != 0 {
			i++
		}
		if i < j {
			for l := i + 1; l <= j; l++ {
				nums[l-1] = nums[l]
			}
			nums[j] = 0
		}

	}

}

func TestMoveZero() {
	array := []int{0, 1, 0, 3, 12}
	moveZeroes(array)
	fmt.Println(array)
	array = []int{1, 0}
	moveZeroes(array)
	fmt.Println(array)
	array = []int{0, 0, 1}
	moveZeroes(array)
	fmt.Println(array)
}
