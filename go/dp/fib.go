package dp

// The Fibonacci numbers, commonly denoted F(n) form a sequence,
// called the Fibonacci sequence, such that each number is the sum of the two preceding ones,
// starting from 0 and 1. That is,

// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.

// Given n, calculate F(n).

// refs:
// https://leetcode.com/problems/fibonacci-number/description/
// https://leetcode.com/problems/climbing-stairs/description/

func fib_recurr(n int) int {
	if n <= 1 {
		return n
	}
	return fib_recurr(n-1) + fib_recurr(n-2)
}

func fib_memo(n int) int {
	if n < 1 {
		return n
	}
	dp := make([]int, n+1)
	func() {
		dp[0] = 0
		dp[1] = 1
		for i := 2; i <= n; i++ {
			dp[i] = -1
		}
	}()

	var helper func(int) int
	helper = func(n int) int {
		if dp[n] != -1 {
			return dp[n]
		}
		dp[n] = helper(n-1) + helper(n-2)
		return dp[n]
	}

	return helper(n)
}

func fib_tabulation(n int) int {
	if n < 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib_tabulation_memory_saver(n int) int {
	if n < 1 {
		return n
	}
	e1, e2 := 0, 1
	for i := 2; i <= n; i++ {
		e1, e2 = e2, e1+e2
	}
	return e2
}
