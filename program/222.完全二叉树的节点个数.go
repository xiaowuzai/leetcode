package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 广度优先搜索
func countNodes_1(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			count++
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return count
}

// 深度优先遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) +1
}
