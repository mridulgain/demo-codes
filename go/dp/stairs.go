package dp

// Problem Statement: Given a number of stairs and a frog,
// the frog wants to climb from the 0th stair to the (N-1)th stair.
// At a time the frog can climb either one or two steps.
// A height[N] array is also given. Whenever the frog jumps from a stair i to stair j,
// the energy consumed in the jump is abs(height[i]- height[j]), where abs() means the absolute difference.
// We need to return the minimum energy that can be used by the frog to jump from stair 0 to stair N-1.

// ref: https://takeuforward.org/data-structure/dynamic-programming-frog-jump-dp-3

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func stairs_memo(heights []int) int {
	n := len(heights)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	var f func(int) int
	f = func(n int) int {
		if dp[n] != -1 {
			return dp[n]
		}
		p1 := f(n-1) + abs(heights[n]-heights[n-1])
		if n > 1 {
			p2 := f(n-2) + abs(heights[n]-heights[n-2])
			dp[n] = min(p1, p2)
			return dp[n]
		}
		dp[n] = p1
		return dp[n]
	}
	return f(n - 1)
}

func stairs_tabulation(heights []int) int {
	n := len(heights)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0], dp[1] = 0, abs(heights[1]-heights[0])
	for i := 2; i < n; i++ {
		p1 := dp[i-2] + abs(heights[i]-heights[i-2])
		p2 := dp[i-1] + abs(heights[i]-heights[i-1])
		dp[i] = min(p1, p2)
	}
	return dp[n-1]
}
