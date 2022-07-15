package medium

func minSubArrayLen(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	size := len(nums)
	min := size + 1
	// head == tail == size -1 就结束？
	head, tail := 0, 0

	for !(head == size-1 && tail == size-1) {
		sum := 0
		for i := head; i <= tail; i++ {
			sum = sum + nums[i]
		}
		if sum >= target {
			curMin := tail - head + 1
			if curMin < min {
				min = curMin
			}
			if head < tail {
				head++
			}
		} else { //increase tail
			tail = tail + 1
		}

	}
	if min > size {
		return 0
	}
	return min

}
func minSubArrayLenBetter(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	size := len(nums)
	min := size + 1
	// head == tail == size -1 就结束？
	head := 0
	sum := 0
	for tail := 0; tail < size; tail++ {
		sum = sum + nums[tail]
		for sum >= target {
			curMin := tail - head + 1
			if curMin < min {
				min = curMin
			}

			sum = sum - nums[head]
			head = head + 1
		}
	}
	if min > size {
		return 0
	}
	return min

}
