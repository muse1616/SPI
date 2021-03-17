package main

func lastRemaining(n int, m int) (ans int) {
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}
