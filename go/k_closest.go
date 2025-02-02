// https://leetcode.com/problems/k-closest-points-to-origin/
package amazon

import "sort"

func kClosest(points [][]int, k int) [][]int {
	distances := make([]int, len(points))
	for i, p := range points {
		distances[i] = p[0]*p[0] + p[1]*p[1]
	}
	indices := make([]int, len(points))
	for i := range indices {
		indices[i] = i
	}
	// Sort the indices based on the values in the numbers slice.
	sort.Slice(indices, func(i, j int) bool {
		return distances[indices[i]] < distances[indices[j]]
	})
	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = points[indices[i]]
	}
	return ans
}
