package main

import "container/list"

// 队列，层级遍历
func maxDepth(root *TreeNode) int {
	level := 0
	if root == nil {
		return level
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		n := queue.Len()
		level++
		for i := 0;i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return level
}

func maxDepth_dfs(root *TreeNode) int {
	max := 0
	if root == nil {
		return max
	}

	var dfs func (node *TreeNode, level int)
	dfs = func (node *TreeNode, level int) {
		if node == nil{
			return
		}
		l := level+1
		if max < l {
			max = l
		}
		dfs(node.Left, l)
		dfs(node.Right, l)
	}

	dfs(root, max)

	return max
}

