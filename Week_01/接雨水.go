package main

// 在某个位置i处，它能存的水，取决于它左右两边的最大值中较小的一个。
func trap(height []int) int {
	if len(height) <= 0 {
		return 0
	}
	left, right := 0, len(height)-1
	maxL, maxR := 0, 0
	ans := 0
	for left <= right {
		if maxL < maxR {
			ans += maxTrap(0, maxL-height[left])
			maxL = maxTrap(maxL, height[left])
			left++
		} else {
			ans += maxTrap(0, maxR-height[right])
			maxR = maxTrap(maxR, height[right])
			right--
		}
	}
	return ans
}

func maxTrap(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
