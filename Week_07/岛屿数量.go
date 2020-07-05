package main

// dfs
func numIslandsByDfs(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				res++
				islandDfs(grid, i, j)
			}
		}
	}
	return res
}

func islandDfs(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(dir); i++ {
		islandDfs(grid, row+dir[i][0], col+dir[i][1])
	}
}

// bfs
func numIslandsByBfs(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				res++
				queue := []int{i*col + j}
				for len(queue) > 0 {
					id := queue[0]
					queue = queue[1:]
					// 注意这里的拆分 应该是 /col 或者 %col
					r, c := id/col, id%col
					if r-1 >= 0 && grid[r-1][c] == '1' {
						queue = append(queue, (r-1)*col+c)
						grid[r-1][c] = '0'
					}
					if c-1 >= 0 && grid[r][c-1] == '1' {
						queue = append(queue, r*col+c-1)
						grid[r][c-1] = '0'
					}
					if r+1 < row && grid[r+1][c] == '1' {
						queue = append(queue, (r+1)*col+c)
						grid[r+1][c] = '0'
					}
					if c+1 < col && grid[r][c+1] == '1' {
						queue = append(queue, r*col+c+1)
						grid[r][c+1] = '0'
					}
				}
			}
		}
	}
	return res
}

// 并查集
func numIslandsByUnionFind(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	p, r := make([]int, row*col), make([]int, row*col)
	cnt := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			id := i*col + j
			if grid[i][j] == '1' {
				p[id] = id
				cnt++
			}
			r[id] = 0
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					islandUnion(p, r, i*col+j, (i-1)*col+j, &cnt)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					islandUnion(p, r, i*col+j, i*col+j-1, &cnt)
				}
				if i+1 < row && grid[i+1][j] == '1' {
					islandUnion(p, r, i*col+j, (i+1)*col+j, &cnt)
				}
				if j+1 < col && grid[i][j+1] == '1' {
					islandUnion(p, r, i*col+j, i*col+j+1, &cnt)
				}
			}
		}
	}
	return cnt
}

// r数组可以认为是当前这个父节点有几个孩子
func islandUnion(p, r []int, x, y int, cnt *int) {
	px, py := islandFind(p, x), islandFind(p, y)
	if px != py {
		if r[px] > r[py] {
			p[py] = px
		} else if r[px] < r[py] {
			p[px] = py
		} else {
			p[py] = px
			r[px]++
		}
		*cnt--
	}
}

// 路径压缩
func islandFind(p []int, i int) int {
	if p[i] != i {
		p[i] = islandFind(p, p[i])
	}
	return p[i]
}
