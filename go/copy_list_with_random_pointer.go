// https://leetcode.com/problems/copy-list-with-random-pointer/
package amazon

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dict := map[*Node]*Node{}
	newHead := *head
	last := &newHead
	dict[head] = last
	last.Next = nil
	for curr := head; curr.Next != nil; curr = curr.Next {
		// dereference & create new copy
		newNode := *(curr.Next)
		last.Next = &newNode
		dict[curr.Next] = &newNode
		last = last.Next
	}
	// traverse new node & update random ptr
	for curr := &newHead; curr != nil; curr = curr.Next {
		curr.Random = dict[curr.Random]
	}
	return &newHead
}
