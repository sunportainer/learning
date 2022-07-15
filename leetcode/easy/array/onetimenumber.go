package array

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
*/
func singleNumber(nums []int) int {
	ret := 0
	for _, i := range nums {
		fmt.Println(i)
		ret = ret ^ i
	}
	return ret
}

func TestSingleNumber() {
	array := []int{2, 2, 1}
	r := singleNumber(array)
	fmt.Println(r)
}

/*
func majorityElement(nums []int) int {

}nums = [3,2,3] 3
nums = [2,2,1,1,1,2,2] 2
*/
func majorityElement(nums []int) int {
	size := len(nums)
	hash := make(map[int]int)
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			hash[v] = hash[v] + 1

		} else {
			hash[v] = 1
		}
	}
	for key, val := range hash {
		if val > size/2 {
			return key
		}
	}
	return 0

}

//莫尔投票法
func majorityElement2(nums []int) int {
	maj := nums[0]
	_ = maj
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count = 1
			maj = nums[i]
		} else if maj == nums[i] {
			count++
		} else {
			count--
		}
	}
	return maj

}
func TestmajorityElement() {
	array := []int{2, 2, 1, 1, 1, 2, 2}
	r := majorityElement(array)
	fmt.Println(r)
	r = majorityElement2(array)
	fmt.Println(r)

}

/*
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
//输出：false
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions/xmlwi1/

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true
*/

func TestSearchMatrix() {
	matrix := [][]int{
		{1, 4, 7, 11, 15}, /*  第一行索引为 0 */
		{2, 5, 8, 12, 19}, /*  第二行索引为 1 */
		{3, 6, 9, 16, 22}, /* 第三行索引为 2 */
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	m := len(matrix)
	n := len(matrix[0])

	fmt.Printf("%d = %v\n", m, n)
	find := searchMatrix(matrix, 20)
	fmt.Println(find)

}
func searchMatrix(matrix [][]int, target int) bool {

	col := len(matrix[0]) - 1
	row := 0
	fmt.Println("Searching from right-top")
	ret := false
	for {
		fmt.Println(matrix[row][col])
		if matrix[row][col] == target {
			ret = true
		}
		if matrix[row][col] < target {
			row = row + 1
		} else if matrix[row][col] > target {
			col = col - 1
		}
		if !(row <= len(matrix)-1 && col >= 0) {
			break
		}

	}
	return ret
}
