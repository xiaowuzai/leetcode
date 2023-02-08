package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	result := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		m[v]= struct{}{}
	}

	for _, v := range nums2 {
		if _, has := m[v]; has {
			result = append(result, v)
		}
	}
	return result
}