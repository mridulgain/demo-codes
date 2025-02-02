package neetcode

import (
	"testing"
)

// generate parentheses
type stack []string

func (s *stack) Push(i string) {
	*s = append(*s, i)
}
func (s *stack) Pop() string {
	n := len(*s)
	i := (*s)[n-1]
	*s = (*s)[:n-1]
	return i
}

func NewStack() stack {
	s := []string{}
	return s
}

func generateParenthesis(n int) []string {
	// st := NewStack()
	// st.Push("[")
	// st.Push("]")
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	ans := []string{}
	gen("", n, &ans, 0, 0)

	return ans
}

func gen(s string, n int, arr *[]string, lb, rb int) {
	if rb > lb {
		return
	}
	if len(s) == n {
		if lb == rb {
			*arr = append(*arr, s)
		}
		return
	}
	gen(s+"(", n, arr, lb+1, rb)
	gen(s+")", n, arr, lb, rb+1)

}

func Test_generateParenthesis(t *testing.T) {
	// var ans []string
	for i := 2; i <= 8; i += 2 {
		ans := generateParenthesis(i)
		t.Log(ans)
	}
}
