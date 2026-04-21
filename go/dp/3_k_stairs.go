package dp

import "math"

func k_stairs_tabulation(heights []int, k int) int {
	n := len(heights)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		minEffort := math.MaxInt
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				break
			}
			effort := dp[i-j] + abs(heights[i]-heights[i-j])
			minEffort = min(effort, minEffort)
		}
		dp[i] = minEffort
	}
	return dp[n-1]
}

func k_stairs_memoization(heights []int, k int) int {
	return 0
}
