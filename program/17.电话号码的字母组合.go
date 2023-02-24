package main

import "strings"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	var numbers = map[string][]string{
		"0": {},
		"1": {},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var (
		res = make([]string, 0)
		path = make([]string, 0, )
		n = len(digits)
	)

	var backtrace func(idx int, path []string)
	backtrace = func(idx int, path []string){
		if idx == n {
			res = append(res, strings.Join(path, ""))
			return
		}

		nums := numbers[string(digits[idx])]
		nl := len(nums)
		for i := 0; i < nl; i++ {
			path = append(path, nums[i])

			backtrace(idx+1, path)

			path = path[:len(path)-1]
		}
	}
	backtrace(0, path)
	return res
}
