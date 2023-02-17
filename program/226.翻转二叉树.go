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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func invertTree_queue(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)

		node.Right, node.Left = node.Left, node.Right
		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return root
}