package main

// 组合、子集、
/*
	分析：组合问题建议使用回溯算法，穷举遍历可能的结果
	输入 n,k。也就是要穷举 1 -> n 的可能结果。
	与全排列不同，需要记录访问过的元素，避免重复访问
*/
func combine(n int, k int) [][]int {
	var (
		res = make([][]int, 0)
		backtrace func(start int, path []int)
		path = make([]int, 0, k)
	)
	backtrace = func(start int, path []int) {
		// 终止条件：path中元素的数量等于 k
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n; i++ {
			// 做选择
			path = append(path, i)

			// 回溯
			backtrace(i+1, path)

			// 撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrace(1,path)
	return res
}