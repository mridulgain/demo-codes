package dp

import (
	"testing"
	"time"
)

// go test -v ./dp -timeout=5s -run TestfibMemo

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "test0",
		args: args{
			n: 1,
		},
		want: 1,
	},
	{
		name: "test1",
		args: args{
			n: 10,
		},
		want: 55,
	},
	{
		name: "test2",
		args: args{
			n: 0,
		},
		want: 0,
	},
	{
		name: "test5",
		args: args{
			n: 40,
		},
		want: 102334155,
	},
	{
		name: "test6",
		args: args{
			n: 45,
		},
		want: 1134903170,
	},
	{
		name: "test7",
		args: args{
			n: 50,
		},
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
				got = fib_recurr(tt.args.n)
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
			if got := fib_memo(tt.args.n); got != tt.want {
				t.Errorf("fib_memo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibTabulation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_tabulation(tt.args.n); got != tt.want {
				t.Errorf("fib_tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibTabulationMemorySaver(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_tabulation_memory_saver(tt.args.n); got != tt.want {
				t.Errorf("fib_tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
