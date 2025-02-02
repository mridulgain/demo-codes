package neetcode

type MinStack struct {
	arr []Node
}

type Node struct {
	value, min int
}

func Constructor() MinStack {
	return MinStack{
		arr: []Node{},
	}
}

func (this *MinStack) Push(val int) {
	m := val
	if this.Len() > 0 && m > this.GetMin() {
		m = this.GetMin()
	}
	newNode := Node{
		value: val,
		min:   m,
	}
	(*this).arr = append((*this).arr, newNode)
}

func (this *MinStack) Pop() {
	(*this).arr = (*this).arr[:this.Len()-1]
}

func (this *MinStack) Top() int {
	return (*this).arr[this.Len()-1].value
}

func (this *MinStack) GetMin() int {
	return (*this).arr[this.Len()-1].min
}

func (this *MinStack) Len() int {
	return len((*this).arr)
}
