package main

type Node struct {
	Val      int
	Children []*Node
}

// 在后序遍历中，我们会先遍历一个节点的所有子节点，再遍历这个节点本身
func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack := []*Node{root}
	res := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 倒序输出
		res = append([]int{node.Val}, res...)
		for i := 0; i < len(node.Children); i++ {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}

// 递归
func postorder2(root *Node) []int {
	res := []int{}
	postNdfs(root, &res)
	return res
}

func postNdfs(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for i := 0; i < len(root.Children); i++ {
		postNdfs(root.Children[i], res)
	}
	*res = append(*res, root.Val)
}
