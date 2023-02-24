package main

/*
	组合问题。回溯算法，记录索引变化
*/
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	end := 9
	path := make([]int, 0,k)

	var backtrace func(path []int, start int)
	backtrace = func(path []int, start int) {
		// 终止条件
		if sumSlice(path, 0)== n && len(path)== k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i <= end; i ++ {
			// 过滤数据
			if sumSlice(path, i) > n {
				return
			}

			// 做选择
			path = append(path, i)

			// 回溯
			backtrace(path,i+1)

			// 撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrace(path, 1)
	return res
}

func sumSlice(path []int, i int) int{
	res := 0
	for _, v := range path {
		res+=v
	}

	res+=i
	return res
}