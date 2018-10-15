package algorithm


/**
	链表实现栈
 */

type NodeStack struct {
	value interface{}
	next  *NodeStack
	prev  *NodeStack
}
type LinkStack struct {
	top  *NodeStack
	size int //栈的大小
}

func (stack *LinkStack) Push(value interface{}) {
	if stack.top == nil {
		newNode := &NodeStack{value, nil, nil}
		stack.top = newNode
	} else {
		newNode := &NodeStack{value, nil, nil}
		current := stack.top
		current.next = newNode
		newNode.prev = current
		stack.top = newNode
	}
	stack.size++
}

func (stack *LinkStack) Pop() interface{} {
	if stack.size == 0 {
		panic("stack is empty")
	}
	current := stack.top
	prevNode := current.prev
	if prevNode == nil {
		stack.size--
		return current.value
	}
	stack.top = prevNode
	prevNode.next = nil
	stack.size--
	return current.value
}

func (stack *LinkStack) Size() int {
	return stack.size
}
