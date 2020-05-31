package main

func nthUglyNumber(n int) int {
	if n < 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		n1, n2, n3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = minUnly(n1, minUnly(n2, n3))
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

func minUnly(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
