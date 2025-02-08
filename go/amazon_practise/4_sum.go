// https://leetcode.com/problems/4sum/
package amazon

import "slices"

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	var qtuple []int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					qtuple = []int{nums[i], nums[j], nums[l], nums[r]}
					ans = append(ans, qtuple)
					for l < r && nums[l] == qtuple[2] {
						l += 1
					}
					for l < r && nums[r] == qtuple[3] {
						r -= 1
					}
				} else if sum < target {
					l += 1
				} else {
					r -= 1
				}
			}
		}
	}
	return ans
}
