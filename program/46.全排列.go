package main

func permute(nums []int) [][]int {
	var (
		res = make([][]int, 0)
		n = len(nums)
		used = make([]bool, n)
	)

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
			if used[i] {
				continue
			}
			// 做选择
			used[i]=true
			path = append(path, nums[i])

			backTrace(path)

			// 撤销选择
			path = path[:len(path)-1]
			used[i]=false
		}
	}

	backTrace([]int{})

	return res
}
