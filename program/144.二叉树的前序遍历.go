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
// preorderTraversal_dg 递归的方式遍历二叉树
func preorderTraversal_dg(root *TreeNode) []int {
	// 递归三要素
	// 1. 确定参数和返回值
	// 2. 确定递归终止条件
	// 3. 确定单层递归逻辑

	data :=  make([]int, 0)
	if root == nil {
		return  data
	}
	data = append(data, root.Val)
	data = append(data, preorderTraversal_dg(root.Left)...)
	data = append(data,preorderTraversal_dg(root.Right)...)
	return data
}


// preorderTraversal_stack 用栈模拟递归
func preorderTraversal_stack(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := list.New()
	node := root
	// 注意： node 不为空，或者 stack 不为空，都要处理
	for node != nil || stack.Len() != 0{
		for node != nil {
			// 保存节点值
			res = append(res, node.Val)
			// 将节点入栈
			stack.PushBack(node)
			// 处理左子树
			node = node.Left
		}

		// 处理右子树
		node  = stack.Remove(stack.Back()).(*TreeNode).Right
	}

	return res
}
