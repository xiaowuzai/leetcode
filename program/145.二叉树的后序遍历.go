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
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

func postorderTraversal_stack(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := list.New()
	node := root

	for node != nil || stack.Len() != 0  {
		if node != nil {
			stack.PushBack(node)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			node = node.Left
		}else {
			res  = append(res,stack.Remove(stack.Back()).(*TreeNode).Val)
		}
	}
	return res
}
