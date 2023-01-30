package main

// removeElement1 创建新数组
// 时间复杂度 O(N)，空间复杂度 O(N)
func removeElement1(nums []int, val int) int {
	arr := make([]int, 0, len(nums))
	for _, v := range nums {
		if v != val {
			arr = append(arr, v)
		}
	}
	return len(arr)
}

// removeElement2 暴力破解
// 时间复杂度 O(N^2)，空间复杂度 O(1)
func removeElement2(nums []int, val int) int {
	// 记录挪动的次数
	n := 0
	num := len(nums)

	for i := 0; i< num; i++ {
		if  nums[i] != val{
			continue
		}

		n++
		for j := i+1; j< len(nums); j++ {
			nums[j-1]= nums[j]  // 移动数据
		}
		i-- // 还是要从移动过来的数据的位置开始。因为这个数值还没有进行确认
		num--  // 移动一次，总数量要减少
	}

	return len(nums)-n
}

// removeElement3 快慢指针
// 时间复杂度 O(N)，空间复杂度 O(1)
func removeElement3(nums []int, val int) int {
	slow := 0
	for quick := 0; quick < len(nums); quick ++ {
		if nums[quick] == val {
			continue
		}
		nums[slow] = nums[quick]
		slow++
	}
	return slow
}


