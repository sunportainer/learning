package array

import "fmt"

func findDuplicate(nums []int) int {

	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow

}

func TestFindDulicate() {
	arry := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(arry))
	arry = []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(arry))
}
