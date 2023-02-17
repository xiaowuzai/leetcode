package main

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

	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int{
		res := make([]int, 0)
		if root == nil {
			return res
		}

		res = append(res, dfs(root.Left)...)
		res = append(res, root.Val)
		res = append(res, dfs(root.Right)...)
		return res
	}

	data := dfs(root)
	
	n := len(data)
	mid := n/2
	for i := 0; i < mid; i++ {
		if data[mid-i] != data[mid+i] {
			return false
		}
	}
	return true
}

