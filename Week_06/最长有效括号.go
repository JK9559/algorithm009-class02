package main

func longestValidParentheses(s string) int {
	if len(s) <= 0 {
		return 0
	}
	stack := []int{-1}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) <= 0 {
				stack = append(stack, i)
			} else {
				// 可以认为在栈里的左括号都是非法的，所以取栈顶元素的下一个位置的左括号作为合法的第一个，求长度就是：i-validBegin+1
				res = maxVali(res, i-(stack[len(stack)-1]+1)+1)
			}
		}
	}
	return res
}

func maxVali(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
