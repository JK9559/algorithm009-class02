package main

// 迭代方法1
func levelOrderN(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, queue := [][]int{}, []*Node{root}
	for i := 0; len(queue) > 0; i++ {
		queue2 := []*Node{}
		res = append(res, []int{})
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			res[i] = append(res[i], node.Val)
			for j := 0; j < len(node.Children); j++ {
				queue2 = append(queue2, node.Children[j])
			}
		}
		queue = queue2
	}
	return res
}

func levelOrderN2(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, queue, link := [][]int{}, []*Node{root, nil}, []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			link = append(link, node.Val)
			for i := 0; i < len(node.Children); i++ {
				queue = append(queue, node.Children[i])
			}
		} else {
			res = append(res, link)
			link = []int{}
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}
	return res
}

// 递归
func levelOrderN3(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	levelDfsN3(root, 0, &res)
	return res
}

func levelDfsN3(root *Node, index int, res *[][]int) {
	if root == nil {
		return
	}
	if index+1 > len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[index] = append((*res)[index], root.Val)
	for i := 0; i < len(root.Children); i++ {
		levelDfsN3(root.Children[i], index+1, res)
	}
}
