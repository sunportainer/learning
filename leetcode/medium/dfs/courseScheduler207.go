package dfs

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)
	//这个构造非常关键，有时候不能直接用给的数据算，必须找个中间状态
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	dfs = func(index int) {
		//进入前判断是不是  == 1，如果是，就不进来了
		visited[index] = 1 // 1 表示正在遍历这个节点，但是还没处理完它的相邻节点
		//遍历临近节点
		for _, next := range edges[index] {
			if visited[next] == 1 {
				valid = false
				return
			} else if visited[next] == 0 {
				dfs(next)
				if !valid {
					return
				}
			}
		}
		visited[index] = 2 //marked as visited
		result = append(result, index)
		return
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid

}

/*
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/course-schedule/solution/ke-cheng-biao-by-leetcode-solution/

*/

func canFinish2(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)
	//这个构造非常关键，有时候不能直接用给的数据算，必须找个中间状态
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	dfs = func(index int) {
		//进入前判断是不是  == 1，如果是，就不进来了
		visited[index] = 1 // 1 表示正在遍历这个节点，但是还没处理完它的相邻节点
		//遍历临近节点
		for _, next := range edges[index] {
			if visited[next] == 1 {
				valid = false
				return
			} else if visited[next] == 0 {
				dfs(next)
				if !valid {
					return
				}
			}
		}
		visited[index] = 2 //marked as visited
		result = append(result, index)
		return
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if valid {
		for i := 0; i < len(result)/2; i++ {
			result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
		}
		return result
	}
	return []int{}

}
