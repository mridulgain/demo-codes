package oa_test

import (
	"sort"
	"testing"
)

func minOperationsToSortWeights(weights, dist []int32) int64 {
	var i, n int32
	n = int32(len(dist))
	// Pair weights with their indices and distances
	points := make([][3]int32, n)
	for i = 0; i < n; i++ {
		points[i] = [3]int32{weights[i], i, dist[i]}
	}

	// Sort points by weights
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// Calculate the total number of moves
	var totalMoves int64
	for i = 1; i < n; i++ {
		prev_index := points[i-1][1]
		curr_index, step := points[i][1], points[i][2]
		for curr_index <= prev_index {
			curr_index += step
			totalMoves++
		}
		// Update index to target for subsequent calculations
		points[i][1] = curr_index
	}

	return totalMoves
}

func Test_generateParenthesis(t *testing.T) {
	// Example Test Case

	weights := []int32{3, 6, 5, 1}
	dist := []int32{4, 3, 2, 1}
	t.Log(minOperationsToSortWeights(weights, dist)) // Output: 5

	weights = []int32{3, 2, 1}
	dist = []int32{1, 4, 5}
	t.Log(minOperationsToSortWeights(weights, dist)) // output: 7
}
