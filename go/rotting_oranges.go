// https://leetcode.com/problems/rotting-oranges/
package amazon

type Item struct {
	x, y int
}
type Queue []*Item

func (q *Queue) insert(x *Item) {
	*q = append(*q, x)
}
func (q *Queue) pop() *Item {
	tmp := *q
	item := tmp[0]
	if len(*q) == 1 {
		*q = Queue{}
	} else {
		*q = tmp[1:]
	}
	return item
}
func (q *Queue) size() int { return len(*q) }

func orangesRotting(grid [][]int) int {
	dq := &Queue{}
	fresh := 0
	delta := []Item{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	row, col := len(grid), len(grid[0])
	// fmt.Println(row,col)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				fresh += 1
			}
			if grid[i][j] == 2 {
				dq.insert(&Item{i, j})
			}
		}
	}
	// fmt.Println(dq.size(), fresh)
	time := 0
	for dq.size() > 0 && fresh > 0 {
		time += 1
		// parse all item in queue, ignore newly added in this process
		// n :=
		for i := dq.size(); i > 0; i-- {
			curr := dq.pop()
			// fmt.Println(curr.x,curr.y)
			for _, del := range delta {
				m, n := curr.x+del.x, curr.y+del.y
				if m >= 0 && m < row && n >= 0 && n < col && grid[m][n] == 1 {
					// fmt.Println(m,n)
					grid[m][n] = 2
					fresh -= 1
					dq.insert(&Item{m, n})
				}
			}
		}
		// fmt.Println(time, dq.size(), fresh)
	}
	if fresh > 0 {
		return -1
	} else {
		return time
	}
	// return
}
