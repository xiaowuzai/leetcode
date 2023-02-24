package main

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0{
		return [][]int{}
	}

	var(
		res =make([][]int, 0)
	)

	// 增加位置索引 start 的目的是过滤结果集中的重复元素。
	// 如果没有这个 start 索引，产生的是全排列后的结果。也就是结果集中会出现重复的元素。
	var backtrace func(path []int, start int)
	backtrace = func(path []int, start int) {
		if sumSlice(path, 0) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if sumSlice(path, candidates[i]) > target {
				continue
			}

			path = append(path, candidates[i])
			// 这里是 i 表示可以重复选择当前元素。
			backtrace(path, i)
			path = path[:len(path)-1]
		}
	}

	backtrace([]int{}, 0)

	return res
}

