package main

// 使用map解决问题
func isAnagram(s string, t string) bool {
	m := make(map[int32]int, 0)
	for _, v :=  range s {
		if n, has := m[v]; has {
			n++
			m[v] = n
		}else {
			m[v] = 1
		}
	}

	for _, v := range t {
		if n, has := m[v]; has {
			n--
			if n == 0 {
				delete(m, v)
			}else {
				m[v] = n
			}
		}else {
			return false
		}
	}

	return len(m)==0
}