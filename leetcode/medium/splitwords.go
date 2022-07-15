package medium

import "fmt"

func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}
func isPali(target string, i, j int) bool {
	for i < j {
		if target[i] != target[j] {
			return false
		}
		i++
		j++
	}
	return true
}
func partitionwithBack(s string) [][]string {
	ans := [][]string{}
	size := len(s)
	isPali := func(target string, i, j int) bool {
		for i < j {
			if target[i] != target[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	var dfs func(temp []string, start int)
	dfs = func(temp []string, start int) {
		if start == size {
			//到底了，返回temp里面的内容
			tc := make([]string, len(temp))
			copy(tc, temp)
			ans = append(ans, tc)
			fmt.Println(ans)
			return
		}
		for i := start; i < size; i++ {
			fmt.Printf("the curent string is %s \n", s[start:i+1])
			if isPali(s, start, i) {
				temp = append(temp, s[start:i+1])
				dfs(temp, i+1)
				temp = temp[:len(temp)-1]
			}
		}

	}

	dfs([]string{}, 0)
	return ans
}

func TestpartitionwithBack() {
	partitionwithBack("aab")
}
