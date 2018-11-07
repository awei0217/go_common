package algorithm

import "fmt"

type nodeLinkList struct {
	value interface{}
	prev  *nodeLinkList
	next  *nodeLinkList
}

/**
环形链表
*/
type LinkList struct {
	head *nodeLinkList
	tail *nodeLinkList
	size int
}

func (list *LinkList) Size() int {

	return list.size
}
func (list *LinkList) AddToFirst(value interface{}) {
	newNode := &nodeLinkList{value, nil, list.head}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.head.prev = newNode
		list.head = newNode
	}
	list.head.prev = list.tail
	list.tail.next = list.head
	list.size++
	newNode = nil
}
func (list *LinkList) GetFirst() *nodeLinkList {
	return list.head
}

func (list *LinkList) AddToTail(value interface{}) {
	newNode := &nodeLinkList{value, list.tail, nil}
	if list.tail == nil {
		list.tail = newNode
		list.head = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
	list.tail.next = list.head
	list.head.prev = list.tail
	list.size++
	newNode = nil
}

func (list *LinkList) GetTail() *nodeLinkList {
	return list.tail
}

func (list *LinkList) Add(value interface{}, index int) {
	if index < 0 || (index > list.size && list.size > 0) {
		panic("index transboundary") //下标越界
	}
	if index == 0 {
		list.AddToFirst(value)
		return
	}
	if index == list.size {
		list.AddToTail(value)
		return
	}
	newNode := &nodeLinkList{value, nil, nil}
	indexNode := list.GetNodeByIndex(index)
	indexNodePre := indexNode.prev
	nextNode := indexNode
	indexNodePre.next = newNode
	newNode.prev = indexNodePre
	newNode.next = nextNode
	nextNode.prev = newNode
	list.size++
	newNode = nil
	indexNode = nil
	indexNodePre = nil
	nextNode = nil
}

func (list *LinkList) GetNodeByIndex(index int) *nodeLinkList {
	if index < 0 || index >= list.size {
		panic("index transboundary") //下标越界
	}
	if list.head == nil {
		panic("empty list")
	}
	currentNode := list.head
	if index+1 < list.size/2 {
		for i := 0; i < index; i++ {
			currentNode = currentNode.next
		}
	} else {
		for i := list.size - 1; i >= index; i-- {
			currentNode = currentNode.prev
		}
	}
	return currentNode
}

func (list *LinkList) Get(index int) interface{} {
	return list.GetNodeByIndex(index).value
}

/**
返回第一出现的下标
*/
func (list *LinkList) IndexOfFirst(value interface{}) int {
	if list.IsEmpty() {
		panic("empty list")
	}
	currentNode := list.head
	tailNode := list.tail
	i := 0
	for {
		if currentNode.value == value {
			return i
		}
		if currentNode == tailNode {
			return -1
		}
		if currentNode.value != value {
			currentNode = currentNode.next
		}
		i++
	}
}

func (list *LinkList) RemoveFirstPop() *nodeLinkList {
	if list.IsEmpty() {
		panic("empty list")
	}
	popNode := list.head
	hn := list.head.next
	list.tail.next = hn
	list.head = hn
	hn.prev = list.tail
	list.size--
	return popNode
}
func (list *LinkList) RemoveTailPop() *nodeLinkList {
	if list.IsEmpty() {
		panic("empty list")
	}
	popNode := list.tail
	tp := list.tail.prev
	list.head.prev = tp
	list.tail = tp
	tp.next = list.head
	list.size--
	return popNode
}
func (list *LinkList) RemoveFirst() {
	if list.IsEmpty() {
		panic("empty list")
	}
	hn := list.head.next
	list.tail.next = hn
	list.head = hn
	hn.prev = list.tail
	list.size--
}
func (list *LinkList) RemoveTail() {
	if list.IsEmpty() {
		panic("empty list")
	}
	tp := list.tail.prev
	list.head.prev = tp
	list.tail = tp
	tp.next = list.head
	list.size--
}

func (list *LinkList) RemoveByIndex(index int) {
	if list.IsEmpty() {
		panic("empty list")
	}
	if index < 0 || index >= list.Size() {
		panic("index transboundary")
	}
	indexNode := list.GetNodeByIndex(index)
	prevNode := indexNode.prev
	nextNode := indexNode.next
	indexNode.prev = nil
	indexNode.next = nil
	indexNode.value = nil
	prevNode.next = nextNode
	nextNode.prev = prevNode
	prevNode = nil
	nextNode = nil
	indexNode = nil
	list.size--
}

func (list *LinkList) IsEmpty() bool {
	return list.Size() == 0
}

/**
双链表反转
*/
func (list *LinkList) Reverse() {
	if list.head == nil || list.tail == nil {
		return
	}
	// 从尾部开始
	p := list.tail
	for p != list.head {
		pre := p.prev  // 记录当前的前节点
		tail := p.next //记录当前的后节点
		//互换
		p.next = p.prev
		p.prev = tail
		//改变当前节点为前节点
		p = pre
	}
	pre := p.next
	p.next = p.prev
	p.prev = pre
	//循环完毕后最后一个节点也就是之前链表的首节点改为尾节点
	list.tail = p
	list.head = p.next
}

func (list *LinkList) Print() {
	head := list.head
	for head.next != list.head {
		fmt.Println(head.value)
		head = head.next
	}
	fmt.Print(head.value)
}
