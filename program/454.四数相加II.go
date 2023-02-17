package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var (
		result = 0
		m = make(map[int]int, len(nums1))
	)

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}

	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			v := 0-v3-v4
			if n, has := m[v]; has {
				result +=n
			}
		}
	}

	return result
}
