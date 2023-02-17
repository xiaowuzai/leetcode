package main

//
func twoSum1(nums []int, target int) []int {
	// 双层for循环

	for i := 0 ; i < len(nums)-1; i++  {
		for j := i+1; j < len(nums); j ++  {
			if nums[i]+ nums[j] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}

// 通过 map 来实现
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i,n1 := range nums {
		m[n1]=i
	}

	for j,n2 := range nums {
		t := target-n2
		i, has := m[t]
		if has && i != j{
			return []int{i,j}
		}
	}
	return nil
}
