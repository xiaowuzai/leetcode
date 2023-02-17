package main

func isHappy(n int) bool {
	// 存储已经计算过的 n
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10) // 个位、十位、百位、千位依次计算
		n = n / 10
	}
	return sum
}