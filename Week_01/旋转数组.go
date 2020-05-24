package main

import "fmt"

/*
189. 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
*/

func rotate(nums []int, k int) {
	k %= len(nums)
	reverseArray(nums, 0, len(nums)-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)
}

func reverseArray(nums []int, start, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}

func rotate2(nums []int, k int) {
	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			tmp := nums[j]
			nums[j] = prev
			prev = tmp
		}
	}
}

func rotate3(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(a); i++ {
		a[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
}

func main() {
	a := []int{-1}
	rotate(a, 3)
	fmt.Println(a)
}
