package easy

import "fmt"

func TwoSumTest() {
	nums := []int{2, 7, 11, 15}
	ret := twoSum(nums, 9)
	fmt.Println(ret)
}

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		if val, ok := hash[target-nums[i]]; ok {
			return []int{i, val}
		} else {
			hash[nums[i]] = i
		}
	}
	return ret
}
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
