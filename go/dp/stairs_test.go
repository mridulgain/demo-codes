package dp

import (
	"testing"
)

var testCases = []struct {
	desc     string
	input    []int
	expected int
}{
	{
		desc:     "single stair",
		input:    []int{1},
		expected: 0,
	},
	{
		desc:     "two stairs - same height",
		input:    []int{5, 5},
		expected: 0,
	},
	{
		desc:     "two stairs - ascending",
		input:    []int{1, 10},
		expected: 9,
	},
	{
		desc:     "two stairs - descending",
		input:    []int{10, 1},
		expected: 9,
	},
	{
		desc:     "basic case - 5 stairs",
		input:    []int{2, 1, 3, 5, 4},
		expected: 2,
	},
	{
		desc:     "basic case - 5 stairs variant",
		input:    []int{7, 5, 1, 2, 6},
		expected: 9,
	},
	{
		desc:     "monotonically increasing",
		input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: 9,
	},
	{
		desc:     "monotonically decreasing",
		input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected: 9,
	},
	{
		desc:     "all same height",
		input:    []int{5, 5, 5, 5, 5, 5},
		expected: 0,
	},
	{
		desc:     "alternating high-low",
		input:    []int{1, 100, 1, 100, 1, 100},
		expected: 99,
	},
	{
		desc:     "large jump at start",
		input:    []int{1, 50, 2, 3, 4},
		expected: 3,
	},
	{
		desc:     "large jump at end",
		input:    []int{1, 2, 3, 4, 50},
		expected: 49,
	},
	{
		desc:     "valley pattern",
		input:    []int{10, 5, 1, 5, 10},
		expected: 10,
	},
	{
		desc:     "peak pattern",
		input:    []int{1, 5, 10, 5, 1},
		expected: 8,
	},
	{
		desc:     "random pattern - 10 stairs",
		input:    []int{30, 10, 60, 10, 60, 50},
		expected: 40,
	},
	{
		desc:     "three stairs - optimal two-step",
		input:    []int{10, 20, 10},
		expected: 0,
	},
	{
		desc:     "three stairs - optimal one-step",
		input:    []int{10, 15, 20},
		expected: 10,
	},
	{
		desc:     "larger input - 15 stairs",
		input:    []int{10, 20, 30, 10, 20, 30, 10, 20, 30, 10, 20, 30, 10, 20, 30},
		expected: 100,
	},
}

func Test_stairs_memo(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := stairs_memo(tC.input); got != tC.expected {
				t.Errorf("expected %d, got %d", tC.expected, got)
			}
		})
	}
}

func Test_stairs_tabulation(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := stairs_tabulation(tc.input); got != tc.expected {
				t.Errorf("Expected: %d, got %d", tc.expected, got)
			}
		})
	}
}
