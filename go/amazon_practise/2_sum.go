// https://leetcode.com/problems/two-sum/
package amazon

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		_, exists := mp[target-n]
		if exists {
			res = append(res, mp[target-n], i)
			break
		}
		mp[n] = i
	}

	return res
}
