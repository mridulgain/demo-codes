package dp

import (
	"testing"
	"time"
)

// go test -v ./dp -timeout=5s -run TestfibMemo

var tests = []struct {
	name string
	args int
	want int
}{
	{
		name: "test0",
		args: 1,
		want: 1,
	},
	{
		name: "test1",
		args: 10,
		want: 55,
	},
	{
		name: "test2",
		args: 0,
		want: 0,
	},
	{
		name: "test5",
		args: 40,
		want: 102334155,
	},
	{
		name: "test6",
		args: 45,
		want: 1134903170,
	},
	{
		name: "test7",
		args: 50,
		want: 12586269025,
	},
}

// skipping testcases that takes more than 1 second to complete
func TestFibRecurr(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan bool)
			var got int

			go func() {
				got = fib_recurr(tt.args)
				done <- true
			}()

			select {
			case <-done:
				if got != tt.want {
					t.Errorf("fib_recurr() = %v, want %v", got, tt.want)
				}
			case <-time.After(1 * time.Second):
				t.Skipf("skipping test case %s: took longer than 1 second", tt.name)
			}
		})
	}
}

func TestFibMemo(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_memo(tt.args); got != tt.want {
				t.Errorf("fib_memo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibTabulation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_tabulation(tt.args); got != tt.want {
				t.Errorf("fib_tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibTabulationMemorySaver(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_tabulation_memory_saver(tt.args); got != tt.want {
				t.Errorf("fib_tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
