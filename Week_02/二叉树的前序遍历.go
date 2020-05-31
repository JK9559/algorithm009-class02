package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res, stack := []int{}, []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return res
}
