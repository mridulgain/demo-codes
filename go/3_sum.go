// https://leetcode.com/problems/3sum/
package amazon

import "slices"

func threeSum(arr []int) [][]int {
	slices.Sort(arr)
	res := [][]int{}
	for i := 0; i < len(arr)-2; i += 1 {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		l := i + 1
		r := len(arr) - 1
		for l < r {
			sum := arr[i] + arr[l] + arr[r]
			if sum < 0 {
				l += 1
			} else if sum > 0 {
				r -= 1
			} else {
				res = append(res, []int{arr[i], arr[l], arr[r]})
				for l < r && arr[l] == res[len(res)-1][1] {
					l += 1
				}
				for l < r && arr[r] == res[len(res)-1][2] {
					r -= 1
				}
			}
		}
	}
	return res
}
