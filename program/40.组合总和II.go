package main

import "sort"

// 输入数组中的元素可以重复，但是不能重复选择。这种情况就需要用索引来标记可选择的位置
func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	var (
		res  = make([][]int,0)
		path = make([]int, 0)
	)

	if n == 0 {
		return res
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i]< candidates[j]
	})

	var backtrace func(path []int, start int)
	backtrace = func(path []int, start int) {
		// path 中的数据等于 target。将结果保存到结果集
		if sumSlice(path,0) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < n; i ++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// 因为无序，所以后面的数据还要遍历
			if sumSlice(path, candidates[i]) > target{
				continue
			}

			// 选择过的数据就要跳过去，所以传入 i+1
			path = append(path, candidates[i])
			backtrace(path, i+1)

			path = path[: len(path)-1]
		}
	}

	backtrace(path, 0)
	return res
}