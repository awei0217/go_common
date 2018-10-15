package algorithm

type NodeQueue struct {
	value interface{}
	prev  *NodeQueue
	next  *NodeQueue
}

/**
无界队列
*/
type LinkQueue struct {
	head *NodeQueue
	tail *NodeQueue
	size int
}

func (queue *LinkQueue) Add(value interface{}) {
	newNode := &NodeQueue{value, nil, nil}
	if queue.head == nil {
		queue.head = newNode
		queue.tail = newNode
	} else {
		currentNode := queue.tail
		queue.tail = newNode
		currentNode.prev = queue.tail
		currentNode = nil
	}
	newNode = nil
	queue.size++
}

func (queue *LinkQueue) Take() interface{} {
	if queue.head == nil {
		panic("queue is empty")
	}
	current := queue.head
	prevNode := queue.head.prev
	queue.head = prevNode
	queue.size--
	prevNode = nil
	return current.value
}
func (queue *LinkQueue) Size() int {
	return queue.size
}
