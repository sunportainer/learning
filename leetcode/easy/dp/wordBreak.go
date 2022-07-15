package dp

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1) //这个size很关键，必须是+1 dp[i] 表示从0-（i-1）的字符串是可以的
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for k := 0; k < i; k++ {
			if dp[k] && wordDictSet[s[k:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

/*
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/word-break/solution/dan-ci-chai-fen-by-leetcode-solution/
*/
