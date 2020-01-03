package algorithm

import (
	"container/list"
)

type TreeNodeBfs struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNodeBfs, x int, y int) bool {

	queue := list.New()
	queue.PushFront(root)

	for queue.Len() > 0 {
		tempQueue := list.New()
		f1, f2 := false, false
		var p1, p2 *TreeNodeBfs
		for queue.Len() > 0 {
			element := queue.Front()
			e := element.Value.(*TreeNodeBfs)
			if e.Left != nil {
				if e.Left.Val == x {
					f1 = true
					p1 = e
				}
				if e.Left.Val == y {
					f2 = true
					p2 = e
				}
				tempQueue.PushFront(e.Left)
			}
			if e.Right != nil {
				if e.Left.Val == x {
					f1 = true
					p1 = e
				}
				if e.Right.Val == y {
					f2 = true
					p2 = e
				}
				tempQueue.PushFront(e.Right)
			}
			if f1 && f2 && p1 != p2 {
				return true
			}
			queue.Remove(element)
		}
		queue.PushFrontList(tempQueue)
	}
	return false
}

func orangesRotting(grid [][]int) int {

	result := 0
	falg := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		temp := make([]bool, len(grid[i]))
		falg[i] = temp
	}
	queue := list.New()
	queue.PushFront(&RowCol{0, 0})

	for queue.Len() > 0 {
		tempQueue := list.New()
		for queue.Len() > 0 {
			ele := queue.Front()
			v := ele.Value.(RowCol)
			if v.Row < len(grid)-1 && grid[v.Row+1][v.Col] != 0 {
				tempQueue.PushFront(&RowCol{v.Row + 1, v.Col})
				result++
			}

			if v.Col < len(grid[v.Row])-1 && grid[v.Row][v.Col+1] != 0 {
				tempQueue.PushFront(&RowCol{v.Row, v.Col + 1})
				result++
			}
			queue.Remove(ele)
		}

		queue.PushFrontList(tempQueue)
	}
	return result / 2
}

type RowCol struct {
	Row int
	Col int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	DFS(root, &result, 0)
	return result
}

func DFS(root *TreeNode, result *[][]int, deep int) {
	if root == nil {
		return
	}
	if len(*result) <= deep {
		*result = append(*result, make([]int, 0))
	}
	if deep%2 == 0 {
		t := (*result)[deep]
		t = append(t, root.Val)
		(*result)[deep] = t
	} else {
		t := (*result)[deep]
		temp := make([]int, 0)
		temp = append(temp, root.Val)
		temp = append(temp, t[:]...)
		(*result)[deep] = temp
	}
	DFS(root.Left, result, deep+1)
	DFS(root.Right, result, deep+1)

}

func findMinHeightTrees(n int, edges [][]int) []int {

	count := make([]int, n)

	m := make(map[int]*[]int, 0)
	for i := 0; i < n; i++ {
		temp := make([]int, 0)
		m[i] = &temp
	}
	for i := 0; i < len(edges); i++ {
		count[edges[i][0]]++
		count[edges[i][1]]++
		if _, ok := m[edges[i][0]]; ok {
			t := m[edges[i][0]]
			*t = append(*t, edges[i][1])
		} else {
			t := make([]int, 0)
			t = append(t, edges[i][1])
			m[edges[i][0]] = &t
		}

	}

	return nil
}
