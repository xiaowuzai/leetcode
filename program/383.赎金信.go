package main

func canConstruct(ransomNote string, magazine string) bool {
	mm := make(map[rune]int)
	for _, c := range magazine {
		mm[c]++
	}

	for _, c := range ransomNote {
		n,has:= mm[c]
		if !has {
			return false
		}
		mm[c]--
		if n == 1 {
			delete(mm,c)
		}
	}
	return true
}