package main

import "fmt"

func findCircleDfs(M [][]int, visited []bool, i int) {
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && !visited[j] {
			visited[j] = true
			findCircleDfs(M, visited, j)
		}
	}
}

func findCircleNumByDfs(M [][]int) int {
	visited := make([]bool, len(M))
	cnt := 0
	for i := 0; i < len(M); i++ {
		if !visited[i] {
			findCircleDfs(M, visited, i)
			cnt++
		}
	}
	return cnt
}

func findCircleNumByBfs(M [][]int) int {
	visited := make([]bool, len(M))
	cnt := 0
	queue := []int{}
	for i := 0; i < len(M); i++ {
		if !visited[i] {
			queue = append(queue, i)
			for len(queue) != 0 {
				s := queue[0]
				visited[s] = true
				queue = queue[1:]
				for j := 0; j < len(M); j++ {
					if M[s][j] == 1 && !visited[j] {
						queue = append(queue, j)
					}
				}
			}
			cnt++
		}
	}
	return cnt
}

func findCircleNumByUnionFind(M [][]int) int {
	p := make([]int, len(M))
	for i := 0; i < len(p); i++ {
		p[i] = -1
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && i != j {
				union(p, i, j)
			}
		}
	}
	cnt := 0
	for i := 0; i < len(p); i++ {
		if p[i] == -1 {
			cnt++
		}
	}
	return cnt
}

func findParent(p []int, i int) int {
	if p[i] == -1 {
		return i
	}
	return findParent(p, p[i])
}

func union(p []int, i, j int) {
	pi := findParent(p, i)
	pj := findParent(p, j)
	if pi != pj {
		p[pi] = pj
	}
}

func main() {
	M := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNumByDfs(M))
	fmt.Println(findCircleNumByBfs(M))
}
