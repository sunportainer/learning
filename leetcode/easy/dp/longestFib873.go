package dp

func lenLongestFibSubseq(arr []int) int {
	ans := 0
	size := len(arr)
	hashmap := make(map[int]int, size)
	for x, v := range arr {
		hashmap[v] = x
	}
	dp := make([][]int, size)
	for x := range dp {
		dp[x] = make([]int, size)
	}
	for i := range arr {
		//dp[i][0] = 1
		for j := i - 1; j >= 0 && j+2 > ans; j-- {
			if arr[i]-arr[j] >= arr[j] {
				break
			}
			if k, ok := hashmap[arr[i]-arr[j]]; ok {
				dp[i][j] = max(3, dp[j][k]+1)
				ans = max(ans, dp[i][j])
			}

		}
	}

	return ans

}
