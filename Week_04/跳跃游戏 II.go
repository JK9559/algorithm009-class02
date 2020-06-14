package main

import "fmt"

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		// 一步步分析：
		// 例子： 2 3 1 1 4 2 1
		// 首先从2 开始 能到达 3 1。这时顺带遍历3 1，3最大到3+1=4位置 1最大到1+2=3位置。贪心来说 4位置更好。
		// 所以选择了3 能到达 1 1 4 因为上一轮第一个1已经比较过了 所以继续往后比较 第二个1 到达3+1=4位置 4最大到4+4=8位置 更好。贪心选择4
		// 所以选择了4 依次往下。
		// len-1 主要是为了防止有刚巧跳到最后节点的，再次起跳已经没意义了
		maxPosition = maxOk(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func maxOk(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 记录跳跃路径--> 一定要看 可以帮助理解上面的代码
func jumpWithPath(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	end, maxPos := 0, 0
	step := 0
	path := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPos {
			path = append(path, i)
			maxPos = i + nums[i]
		}
		if i == end {
			step++
			end = maxPos
		}
	}
	fmt.Println(path)
	return step
}

func main() {
	s := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(s))
	fmt.Println(jumpWithPath(s))
}
