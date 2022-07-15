package dfs

import "fmt"

//https://leetcode.cn/problems/number-of-islands/solution/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
//有空了把上面4个岛屿的题给做了
func numIslands(grid [][]byte) (ans int) {
	rows := len(grid)
	cols := len(grid[0])
	var dfs func(x, y int)
	dfs = func(i, j int) {
		//has out of the grid
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return
		}
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
		return
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}
func TestNumberIslands1() {
	array := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(array))
}
func numIslands1(grid [][]byte) int {
	var ans int
	var m, n = len(grid), len(grid[0])
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = '0'
		for _, dir := range dirs {
			if x+dir[0] >= 0 && x+dir[0] < m && y+dir[1] >= 0 && y+dir[1] < n && grid[dir[0]+x][dir[1]+y] == '1' {
				dfs(x+dir[0], y+dir[1])
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
