package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// tail[k] 代表长度为k+1的子序列尾部的值
	tail := make([]int, len(nums))
	res := 0
	// 遍历每一个nums的值 找到在tail中的位置 因为需要结果尽可能长，所以需要tail尽可能扁
	for i := 0; i < len(nums); i++ {
		l, r := 0, res
		for l < r {
			m := l + (r-l)/2
			// 找到第一个大于nums[i]的最小值
			// why?? 因为tail[k]存的是长度为k+1的子序列的最小值。
			// 通过二分查找来查找当前元素是不是在tail里
			// 如果在tail的范围里，那么找第一个大于nums[i]的最小值。更新下，为了保证：递增的序列尽量扁平。
			// 那么为啥不更新小于nums的最大值？？因为如果这么样了 序列就不能更扁平了。
			if tail[m] < nums[i] {
				l = m + 1
			} else if tail[m] > nums[i] {
				r = m
			} else if tail[m] == nums[i] {
				r = m
			}
		}
		// 这里找到了tail中第一个大于nums[i]的最小值 放进去
		tail[l] = nums[i]
		if l == res {
			res++
		}
	}
	fmt.Println(tail)
	return res
}
func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 21, 18}))
}
