package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)

	var backTrace func(path []int)
	backTrace = func(path []int) {
		// 判断条件
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			res = append(res,tmp)
			return
		}

		for i :=0; i<n; i++ {
			// 筛除掉不符合规则的

			path = append(path, nums[i])
			backTrace(path)
			path = path[:len(path)-1]
		}
	}

	return res
}
