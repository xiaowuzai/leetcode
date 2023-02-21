package main

// 组合、子集、
func combine(n int, k int) [][]int {
	res := make([][]int,0)
	var backTrace func(start int, path []int)

	backTrace = func( start int, path []int){
		// 终止条件
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res,temp)
			return
		}

		for i := start; i <= n; i++  {
			// 尝试
			path = append(path, i)

			// 进入下一层。
			backTrace(i+1, path)

			// 回退
			path = path[:len(path)-1]
		}
	}
	backTrace(1,[]int{})
	return res
}