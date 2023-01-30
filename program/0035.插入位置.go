package main

// searchInsert1 暴力破解
func searchInsert1(nums []int, target int) int {
	for i :=0; i < len(nums); i ++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// searchInsert2 前闭后闭
func searchInsert2(nums []int, target int) int {
	// 前闭后闭
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + ((end - start)/ 2)
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			start = mid+1
		}else {
			end = mid-1
		}
	}

	return end+1
}

// searchInsert3 前闭后开
func searchInsert3(nums []int, target int) int {
	// 前闭后闭
	start, end := 0, len(nums)
	for start < end {
		mid := start + ((end - start)/ 2)
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			start = mid+1
		}else {
			end = mid
		}
	}

	return end
}
