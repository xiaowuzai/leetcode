package main

import (
	"math"
)

// minSubArrayLen1
func minSubArrayLen1(target int, nums []int) int {
	result := math.MaxInt
	for i := 0; i < len(nums); i++{
		sum := nums[i]
		if sum >= target {
			return 1
		}

		for j := i+1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				n := j-i + 1
				if n < result {
					result = n
				}
			}
		}
	}


	if result == math.MaxInt {
		return 0
	}
	return result
}


// minSubArrayLen2
// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	result := math.MaxInt
	// i 表示窗口的左侧范围
	i:=0
	sum := 0

	// j表示窗口的右侧
	for j := 0; j < len(nums); j++ {
		// 当某一个值大于等于目标值时，1肯定是最小的值了，不需要再继续查找了
		if nums[j] >=target {
			return 1
		}

		sum += nums[j]
		// 当求和大于目标值时
		for sum >= target {
			// 记录最小的结果
			if result > j-i +1 {
				result = j-i+1
			}

			// 计算好了以后，将窗口的左边界向前挪动一步
			sum -= nums[i]
			i++
		}
	}

	if result == math.MaxInt {
		return 0
	}
	return result
}


