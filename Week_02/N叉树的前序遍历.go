package main

// 迭代 注意倒序子节点
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res, stack := []int{}, []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}

// 递归
func preorder1(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	prehelper(root, &res)
	return res
}

func prehelper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for i := 0; i < len(root.Children); i++ {
		prehelper(root.Children[i], res)
	}
}
