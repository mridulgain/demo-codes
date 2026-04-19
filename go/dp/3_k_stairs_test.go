package dp

import "testing"

func Test_k_stairs(t *testing.T) {
	type args struct {
		h []int
		k int
	}

	var testCases = []struct {
		name     string
		input    args
		expected int
	}{
		{
			name: "test1",
			input: args{
				h: []int{15, 4, 1, 14, 15},
				k: 3,
			},
			expected: 2,
		},
		{
			name: "test2",
			input: args{
				h: []int{10, 5, 20, 0, 15},
				k: 2,
			},
			expected: 15,
		},
	}
	t.Run("k staris tabulation", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if got := k_stairs_tabulation(tc.input.h, tc.input.k); got != tc.expected {
					t.Errorf("k_stairs_tabulation() = %v, want %v", got, tc.expected)
				}
			})
		}
	})

	t.Run("k staris memoization", func(t *testing.T) {
		t.Skip("not implemented")
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if got := k_stairs_memoization(tc.input.h, tc.input.k); got != tc.expected {
					t.Errorf("k_stairs_memoization() = %v, want %v", got, tc.expected)
				}
			})
		}
	})

}
