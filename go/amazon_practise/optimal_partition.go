// https://leetcode.com/problems/optimal-partition-of-string/
package amazon

func partitionString(s string) int {
	f := map[rune]bool{}
	count := 1 // atleast one substring would be there
	for _, char := range s {
		if _, exists := f[char]; exists {
			count += 1
			f = map[rune]bool{}
			f[char] = true
		} else {
			f[char] = true
		}
	}
	return count
}
