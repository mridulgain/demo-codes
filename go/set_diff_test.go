// {1,1,2,4,5,6}
// {1,3,3,5,6,8}
// Diff= {2, 3, 4, 8}
// O(m+n)
package main

import (
	"fmt"
	"testing"
)

func findDiff(l1, l2 []int) (res []int) {
	s1 := make(map[int]bool)
	for _, x := range l1 {
		s1[x] = true
	}
	s2 := make(map[int]bool)
	for _, x := range l2 {
		s2[x] = true
	}

	hm := make(map[int]int)
	for k, _ := range s1 {
		hm[k]++
	}
	for k, _ := range s2 {
		hm[k]++
	}

	for k, v := range hm {
		if v > 1 {
			continue
		}
		res = append(res, k)
	}
	return res
}

func Test_diff(t *testing.T) {
	l1 := []int{1, 1, 2, 5, 4, 6}
	l2 := []int{1, 3, 3, 5, 6, 8}
	ans := findDiff(l1, l2)
	fmt.Println(ans)
}
