package main

// 递归法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 3种情况
	// 以最简单的情况来定 --> 三个节点
	// 如果 root.Left = p && root.Right = q 那么最近公共祖先为 root
	// 如果 root.Left = p && root.Right != q 那么最近公共祖先为 p
	// 如果 root.Left != p && root.Right != q 那么最近公共祖先为 nil
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if right != nil {
		return right
	} else {
		return left
	}
}

// 迭代法 --> 思路是：先dfs遍历一遍，把每个节点的父节点都保存在map里。
// 然后选取一个节点，按照父节点的路径 向根遍历
// 然后在用另一个节点，按照父节点的路径 向根遍历，如果当前节点已经遍历过，那么返回
func lowestCommonAncestorMap(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	parent := make(map[int]*TreeNode)
	lowestDfs(parent, root)
	visited := make(map[int]bool)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func lowestDfs(p map[int]*TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		p[root.Left.Val] = root
		lowestDfs(p, root.Left)
	}
	if root.Right != nil {
		p[root.Right.Val] = root
		lowestDfs(p, root.Right)
	}
}
