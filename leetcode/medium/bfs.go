package medium

import "fmt"

//beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions/x2rkbs/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//marks to remember which word has been put to queue

	df := func(t string) []string {
		length := len(t)
		var ret []string

		for i := 0; i < length; i++ { //挨个替换每个位置
			for ch := 'a'; ch <= 'z'; ch++ {
				replace := t[0:i] + string(ch) + t[i+1:]
				ret = append(ret, replace)
			}
		}
		return ret
	}

	marks := make(map[string]bool)
	for _, w := range wordList {
		marks[w] = true
	}
	queue := []string{beginWord}
	step := 1
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == endWord {
				return step
			}
			//每个位置循环替换成一下一个字母

			alternatives := df(node)
			for _, v := range alternatives {
				if marks[v] {
					queue = append(queue, v)
					marks[v] = false
				}
			}
		}

		step++

	}
	return 0
}

func TestLoadLadder() {
	b := "hit"
	e := "dog"
	lst := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(b, e, lst))

}
