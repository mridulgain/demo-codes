package oa_test

import (
	"log"
	"slices"
	"testing"
)

func calculateSmallestDist(src, dest, total_nodes int) (time int) {
	var d1, d2 int
	if src > dest {
		d1 = src - dest
	} else {
		d1 = dest - src
	}
	d2 = total_nodes - d1
	if d1 < d2 {
		return d1
	} else {
		return d2
	}
}

// clock wise distance
func calculateCwDist(p1, p2 int) (time int) {
	// var d1, d2 int
	if p1 > p2 {
		return p1 - p2
	} else {
		return p2 - p1
	}
}

func solution(total_nodes int, servers []int) (time int) {
	slices.Sort(servers)
	src := servers[0]
	// distance := make([]int, len(servers)-1)
	clockwise_dist := 0
	for i := 0; i < len(servers); i++ {
		if i > 0 {
			clockwise_dist += servers[i] - servers[i-1]
		}
	}
	log.Println(servers)
	dest := servers[1:]
	slices.Reverse(dest)
	s2 := append([]int{src}, dest...)
	anti_clockwise_dist := 0
	for i := 0; i < len(s2); i++ {
		if i > 0 {
			anti_clockwise_dist += -calculateCwDist(s2[i-1], s2[i])
		}
	}
	log.Println(s2)
	if clockwise_dist < anti_clockwise_dist {
		return clockwise_dist
	} else {
		return anti_clockwise_dist
	}
}

func TestSolution(t *testing.T) {
	ans := solution(8, []int{2, 6, 8})
	expectedAns := 4
	if ans != expectedAns {
		t.Errorf("solution() = %d, want %d", ans, expectedAns)
	}

	ans = solution(5, []int{1, 5})
	expectedAns = 1
	if ans != expectedAns {
		t.Errorf("solution() = %d, want %d", ans, expectedAns)
	}
}
