package array

func topKFrequent(nums []int, k int) []int {
	size := len(nums)
	if size == 0 {
		return nums
	}
	hashmap := make(map[int]int, size)

	for _, value := range nums {
		if _, has := hashmap[value]; has {
			hashmap[value]++
		} else {
			hashmap[value] = 1
		}
	}

	//make a array of size = k
	ansArray := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		ansArray[i] = 0
	}
	//array里面的都是排好序的，所以找到一个比他大的就停下
	for key, v := range hashmap {
		l := k - 1
		for l >= 0 {
			temp := ansArray[l]
			if temp == 0 || v >= hashmap[ansArray[l]] {
				ansArray[l+1] = ansArray[l]
				l--
			} else {
				break
			}

		}
		//如果已经走到这一步，说明ansArray[l] 已经小于 nums[i]，把nums[i]插入它后面
		ansArray[l+1] = key
	}
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, ansArray[i])
	}
	return result

}

//这个做法可以参考https://leetcode.cn/problems/top-k-frequent-elements/solution/bi-guan-fang-ti-jie-geng-kuai-de-fang-fa-lgsc/
//确实很巧妙
func topKFrequentByBuckets(nums []int, k int) []int {
	size := len(nums)
	if size == 0 {
		return nums
	}
	hashmap := make(map[int]int, size)

	for _, value := range nums {
		if _, has := hashmap[value]; has {
			hashmap[value]++
		} else {
			hashmap[value] = 1
		}
	}

	buckets := make([][]int, size+1)

	for key, val := range hashmap {
		//key是数字， val 是出现次数
		if len(buckets[val]) < 0 {
			buckets[val] = []int{}
		}
		buckets[val] = append(buckets[val], key)
	}
	ans := []int{}

	for i := size; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			for b := 0; b < len(buckets[i]); b++ {
				ans = append(ans, buckets[i][b])
				k = k - 1
				if k == 0 {
					return ans
				}
			}
		}
	}

	return ans

}
