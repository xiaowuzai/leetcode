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
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		data := make([]int,0)
		n := queue.Len()

		for i := 0; i <  n; i++{
			node := queue.Remove(queue.Front()).(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			data = append(data, node.Val)
		}
		res = append(res, data)
	}
	return res
}