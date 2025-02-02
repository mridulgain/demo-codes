// https://leetcode.com/problems/the-kth-factor-of-n/
package amazon

import "slices"

// func kthFactor(n int, k int) int {
//     fc := 0
//     for i := 1; i <= n/2; i++{
//         if n % i == 0 {
//             fc+=1
//         }
//         if fc == k{
//             return i
//         }
//     }
//     // the number itself
//     if fc == k - 1{
//         return n
//     }
//     return -1
// }

func kthFactor(n int, k int) int {
	factors := map[int]bool{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors[i] = true
			factors[n/i] = true
		}
	}
	f_list := []int{}
	for factor := range factors {
		f_list = append(f_list, factor)
	}
	slices.Sort(f_list)
	if len(f_list) < k {
		return -1
	}
	return f_list[k-1]

}
