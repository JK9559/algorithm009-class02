package main

import "fmt"

/*
转换为树的问题，回溯算法
*/
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	help1(len(nums), &nums, 0, &res)
	return res
}

func help1(n int, nums *[]int, first int, res *[][]int) {
	if n == first {
		t := make([]int, len(*nums))
		copy(t, *nums)
		*res = append(*res, t)
	}
	used := map[int]bool{}
	for i := first; i < n; i++ {
		if _, ok := used[(*nums)[i]]; ok {
			continue
		}
		(*nums)[i], (*nums)[first] = (*nums)[first], (*nums)[i]
		help1(n, nums, first+1, res)
		(*nums)[i], (*nums)[first] = (*nums)[first], (*nums)[i]
		used[(*nums)[i]] = true
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}
