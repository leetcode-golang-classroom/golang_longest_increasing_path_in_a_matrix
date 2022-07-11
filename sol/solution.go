package sol

func longestIncreasingPath(matrix [][]int) int {
	ROW := len(matrix)
	COL := len(matrix[0])
	dp := make([][]int, ROW)
	for row := range dp {
		dp[row] = make([]int, COL)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(row, col, previous int) int
	dfs = func(row, col, previous int) int {
		if row < 0 || row == ROW || col < 0 || col == COL || matrix[row][col] <= previous {
			return 0
		}
		if dp[row][col] != 0 { // has computed
			return dp[row][col]
		}
		result := 1
		result = max(result, 1+dfs(row-1, col, matrix[row][col]))
		result = max(result, 1+dfs(row+1, col, matrix[row][col]))
		result = max(result, 1+dfs(row, col-1, matrix[row][col]))
		result = max(result, 1+dfs(row, col+1, matrix[row][col]))
		dp[row][col] = result
		return result
	}
	maxLen := 1
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			result := dfs(row, col, -1)
			if result > maxLen {
				maxLen = result
			}
		}
	}
	return maxLen
}
