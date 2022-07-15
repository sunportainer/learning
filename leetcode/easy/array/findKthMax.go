package array

/*
[3,2,3,1,2,4,5,5,6] 和 k = 4  输出4
[3,2,1,5,6,4] 和 k = 2  输出5
https://leetcode.cn/leetbook/read/top-interview-questions/xal9h6/
*/
func findKthLargest(nums []int, k int) int {
	size := len(nums)
	if nil == nums || size == 0 {
		return -10001
	}
	//make a array of size = k
	ansArray := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		ansArray[i] = -10001
	}
	//now tranverse the nums and find the K
	for i := 0; i < size; i++ {
		l := k - 1
		for l >= 0 {
			if ansArray[l] < nums[i] {
				ansArray[l+1] = ansArray[l]
			} else {
				break
			}
			l--
		}
		ansArray[l+1] = nums[i]
	}
	return ansArray[k-1]

}
