package main

// search1 循环遍历
func search1(nums []int, target int) int {
	for i :=0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// search2 二分查找 前闭后闭
func search2(nums []int, target int) int {

	var (
		start = 0
		end = len(nums)-1
	)

	// 前闭后闭，start 和 end 有交集
	for start <= end {
		index := ((end-start)/2) + start // 防止算术溢出
		if target > nums[index] {
			start = index+1
		}else if  target  < nums[index]  {
			end = index-1
		}else {
			return index
		}

	}
	return -1
}

// search3 二分查找 前闭后开
func search3(nums []int, target int) int {
	var (
		start = 0
		end = len(nums)-1
	)

	// 前闭后开，start 和 end 一定不会有交集
	for start < end {
		index := ((end-start)/2) + start // 防止算术溢出
		if target > nums[index] {
			start = index+1  // 前闭，start有效
		}else if  target  < nums[index]  {
			end = index // 后开, end 不计算在内
		}else {
			return index
		}
	}
	return -1
}
