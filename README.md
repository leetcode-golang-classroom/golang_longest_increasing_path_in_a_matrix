# golang_longest_increasing_path_in_a_matrix

Given an `m x n` integers `matrix`, return *the length of the longest increasing path in* `matrix`.

From each cell, you can either move in four directions: left, right, up, or down. You **may not** move **diagonally** or move **outside the boundary** (i.e., wrap-around is not allowed).

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg](https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg)

```
Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: The longest increasing path is[1, 2, 6, 9].

```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg](https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg)

```
Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation:The longest increasing path is[3, 4, 5, 6]. Moving diagonally is not allowed.

```

**Example 3:**

```
Input: matrix = [[1]]
Output: 1

```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 200`
- `0 <= matrix[i][j] <= 231 - 1`

## 解析

給定一個 2維整數矩陣 matrix 

每個 entry 都是大於等於 0

要求寫一個演算法找出一個 嚴格遞增的數列

數列串聯的方式只能透過上下左右四個方向去連接

不能透過對角線也不能跨越邊界反向

直覺的作法是

從每個座標去從上下左右方向做DFS 找最長的遞增數列長度

因為每次都有4個方向可以選擇

所以畫成 決策樹如下

![](https://i.imgur.com/wBjkY3j.png)

因為每個座標都要走訪一遍所以是時間複雜度是 O($4^{m*n+k}$) , k = 最大長度

然而可以發現其實 從每個座標出發所能形成的最大長度可以可以猜解成

dfs(row, col, previous) = 1 + max(dfs(row-1, col, previous), dfs(row, col-1, previous), dfs(row, col-1, previous),dfs(row+1, col, previous), dfs(row, col+1, previous))

![](https://i.imgur.com/IIqqOjd.png)

而走過點 透過 hashtable 的方式儲存起來 就可以不必重複走訪

m, n 組合最多有 m * n 種

所以只要時間複雜度 O(m*n )

空間複雜度也是 O(m*n) 

## 程式碼
```go
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
```
## 困難點

1. 需要看出 每個出發位置遞迴關係

## Solve Point

- [x]  建立一個 hashtable cache 用來存每次走過的最大長度
- [x]  每次檢查四個可能方向是否可能可以繼續往下找逐步把可能的結果找出來