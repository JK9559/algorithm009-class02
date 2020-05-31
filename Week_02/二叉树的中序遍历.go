package main

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res, stack := []int{}, []*TreeNode{}
	cur := root
	// 这里应该是 || 不是 &&
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}
