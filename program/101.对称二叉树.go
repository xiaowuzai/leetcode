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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() != 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)

		if left ==nil && right == nil {
			continue
		}else if left == nil || right == nil {
			return false
		}else if left.Val != right.Val {
			return false
		}

		queue.PushBack(left.Left)
		queue.PushBack(right.Right)

		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}

	return true
}



func isSymmetric_1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)

}
