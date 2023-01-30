package main

// searchRange1 暴力破解法
func searchRange1(nums []int, target int) []int {
	start, end := -1, -1
	for i :=0; i < len(nums); i++ {
		if nums[i] == target {
			if start == -1 {
				start, end = i, i
			}else {
				end=i
			}
		}
	}

	return []int{start, end}
}

// searchRange2 二分法.前闭后闭
func searchRange2(nums []int, target int) []int {
	left, hasLeft := findLeft(nums, target)
	right, hasRight := findRight(nums, target)
	if hasLeft && hasRight {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func findLeft(nums []int, target int) (int, bool){
	left, right := 0, len(nums)-1
	has := false

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			// 找左区间
			right = mid-1
			has = true
		}else if nums[mid] > target {
			right = mid-1
		}else {
			left = mid+1
		}
	}



	return left,has
}

// 找右区间
func findRight(nums []int, target int)(int, bool) {
	left, right := 0, len(nums)-1
	has := false

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			left = mid+1
			has = true
		}else if nums[mid] > target {
			right = mid -1
		}else {
			left = mid +1
		}
	}
	return right, has
}
