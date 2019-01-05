package basic

import (
	"container/list"
	"fmt"
)

type TreeNodeNew struct {
	Val   int
	Left  *TreeNodeNew
	Right *TreeNodeNew
}

func QueueStudy() {

	queue := list.New()

	left := &TreeNodeNew{1, nil, nil}
	right := &TreeNodeNew{3, nil, nil}
	root := &TreeNodeNew{2, left, right}

	//往队列末尾添加一个
	queue.PushBack(*root)
	//从队列头部区取一个
	e := queue.Front()
	// 将取出的元素转换为struct类型
	temp := e.Value.(TreeNodeNew)
	//删除一个元素
	queue.Remove(e)
	fmt.Println(temp)

	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(4)
	backE := queue.Back()
	// 将元素backE移动到list的首部
	queue.MoveToFront(backE)
	fmt.Println(queue.Front().Value)

	//func (l *List) MoveAfter(e, mark *Element)  //将元素e移动到元素mark之后，如果元素e或者mark不属于list l，或者e==mark，则list l不改变。
	//func (l *List) MoveBefore(e, mark *Element)//将元素e移动到元素mark之前，如果元素e或者mark不属于list l，或者e==mark，则list l不改变。

	//func (l *List) MoveToBack(e *Element)//将元素e移动到list l的末尾，如果e不属于list l，则list不改变。
	//func (l *List) MoveToFront(e *Element)//将元素e移动到list l的首部，如果e不属于list l，则list不改变。
	//func (l *List) PushBack(v interface{}) *Element//在list l的末尾插入值为v的元素，并返回该元素。
	//func (l *List) PushBackList(other *List)//在list l的尾部插入另外一个list，其中l和other可以相等。
	//func (l *List) PushFront(v interface{}) *Element//在list l的首部插入值为v的元素，并返回该元素。
	//func (l *List) PushFrontList(other *List)//在list l的首部插入另外一个list，其中l和other可以相等。
	//func (l *List) Remove(e *Element) interface{}//如果元素e属于list l，将其从list中删除，并返回元素e的值。
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(*root)
	for queue.Len() > 0 {
		res = append(res, queue.Back().Value.(TreeNode).Val)
		newQueue := list.New()
		for queue.Len() > 0 {
			e := queue.Front()
			temp := e.Value.(TreeNode)
			queue.Remove(e)
			if temp.Left != nil {
				newQueue.PushBack(*temp.Left)
			}
			if temp.Right != nil {
				newQueue.PushBack(*temp.Right)
			}
		}
		queue.PushBackList(newQueue)
	}
	return res
}
