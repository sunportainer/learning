package array

func longestConsecutive(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	maps := make(map[int]bool, size)
	for _, val := range nums {
		maps[val] = true
	}
	maxseq := 0
	for _, val := range nums { //loop all the nums array
		cursize := 1
		if maps[val-1] { //如果前驱在，退出
			continue
		}
		for maps[val+1] {
			cursize++
			val++
		}
		if maxseq < cursize {
			maxseq = cursize
		}
	}
	return maxseq
}
