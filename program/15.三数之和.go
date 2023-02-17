package main

import (
	"sort"
)

// 像这种时间复杂度可能超过 O(N^2) 的题目。就想着是否可以通过排序来解决问题了
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	//排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := make([][]int, 0)

	for i := 0 ; i < len(nums); i ++ {
		if i > 0  && nums[i] == nums[i-1] {
			continue
		}

		left := i +1
		right := len(nums)-1

		for left < right {
			n := nums[i]+ nums[left]+ nums[right]
			if n >0 {
				right--
			}else if n <0 {
				left++
			}else {
				result = append(result,[]int{nums[i], nums[left], nums[right]})
				for left < right && nums[right]==nums[right-1] {
					right--
				}
				for left < right && nums[left]==nums[left+1] {
					left++
				}

				left++
				right--
			}
		}

	}

	return result
}