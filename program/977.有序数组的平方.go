package main

import "sort"

//sortedSquares1 循环计算之后再排序
// 时间复杂度 O(NlogN)，空间复杂度 O(1)
func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i]	= nums[i]*nums[i]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]	< nums[j]
	})

	return nums
}

// sortedSquares2 循环计算之后再排序
// 时间复杂度 O(N),空间复杂度 O(N)
func sortedSquares2(nums []int) []int {
	arr := make([]int,len(nums))
	k,j := len(nums)-1, len(nums)-1
	i := 0
	for  i <= j{
		// 谁值大就放谁，从后往前放
		if nums[j]*nums[j] > nums[i]* nums[i] {
			arr[k] = nums[j]*nums[j]
			j--
		}else {
			arr[k] = nums[i]*nums[i]
			i++
		}
		k--
	}

	return arr
}
