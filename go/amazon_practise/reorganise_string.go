// https://leetcode.com/problems/reorganize-string/
package amazon

import (
	"container/heap"
)

type Item struct {
	char byte
	freq int
}

type MaxHeap []*Item

func (pq MaxHeap) Len() int { return len(pq) }

// max heap
func (pq MaxHeap) Less(i, j int) bool { return pq[i].freq > pq[j].freq }
func (pq MaxHeap) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MaxHeap) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *MaxHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func reorganizeString(s string) string {
	// chars := make([]byte, len(s))
	freqMap := map[byte]int{}
	for i := range s {
		freqMap[s[i]] += 1
	}
	// fmt.Println(freqMap)
	pq := MaxHeap{}
	for k, v := range freqMap {
		item := &Item{
			char: k,
			freq: v,
		}
		pq = append(pq, item)
	}
	// fmt.Println(maxHeap)
	heap.Init(&pq)
	// fmt.Println(maxHeap)
	ans := make([]byte, len(s))
	for i := 0; pq.Len() > 1; {
		item1 := heap.Pop(&pq).(*Item)
		item2 := heap.Pop(&pq).(*Item)
		// fmt.Println(item)
		ans[i] = item1.char
		ans[i+1] = item2.char
		i += 2
		item1.freq -= 1
		item2.freq -= 1
		if item1.freq > 0 {
			heap.Push(&pq, item1)
		}
		if item2.freq > 0 {
			heap.Push(&pq, item2)
		}
	}
	// fmt.Println(ans)
	if pq.Len() == 1 {
		item := heap.Pop(&pq).(*Item)
		if item.freq == 1 {
			ans[len(s)-1] = item.char
		} else {
			return ""
		}
	}

	return string(ans)
}
