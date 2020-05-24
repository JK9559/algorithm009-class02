package main

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmpMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		tarJ := target - nums[i]
		if j, ok := tmpMap[tarJ]; ok {
			if i < j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
