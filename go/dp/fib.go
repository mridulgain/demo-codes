package dp

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
		if n <= 1 {
			return dp[n]
		}
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
