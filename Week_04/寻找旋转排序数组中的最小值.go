package main

func findMin1(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	if nums[l] < nums[r] {
		return nums[l]
	}
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			return nums[m+1]
		} else if nums[l] < nums[m] {
			l = m + 1
		} else if nums[l] > nums[m] {
			r = m
		} else if nums[l] == nums[m] {
			r = m
		}
	}
	return nums[l]
}
