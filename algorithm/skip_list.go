package algorithm

import (
	"math/rand"
	"sync/atomic"
)

//节点存储的key 和 value
type Node struct {
	index     uint64 //用来排序
	value     interface{}
	nextNodes []*Node // 这个节点的当前层存储的节点 0 - 31
}

/**
	创建一个节点，创建这个节点时指定这个节点所在层
 */
func newNode(index uint64, value interface{}, level int) *Node {
	return &Node{
		index:     index,
		value:     value,
		nextNodes: make([]*Node, level, level),
	}
}

/**
	返回节点的key
 */
func (n *Node) Index() uint64 {
	return n.index
}
/**
	返回节点的value
 */
func (n *Node) Value() interface{} {
	return n.value
}
/**
	空间换时间的一种数据结构，时间复杂度为log2n
 */
type skipList struct {
	level  int   //跳跃俩表的最大层高
	length int32 //跳跃列表的长度
	head   *Node //跳跃列表头节点
	tail   *Node //跳跃列表尾节点
}

/**
	创建一个跳跃列表，入参为最大层
 */
func NewSkipList(level int) *skipList {
	head := newNode( 0, nil, level)
	var tail *Node
	for i := 0; i < len(head.nextNodes); i++ {
		head.nextNodes[i] = tail
	}

	return &skipList{
		level:  level,
		length: 0,
		head:   head,
		tail:   tail,
	}
}

// searchWithPreviousNode will search given index in skip list.
// The first return value represents the previous nodes need to update when call Insert function.
// The second return value represents the value with given index or the closet value whose index is larger than given index.
func (s *skipList) searchWithPreviousNodes(index uint64) ([]*Node, *Node) {
	// Store all previous value whose index is less than index and whose next value's index is larger than index.
	previousNodes := make([]*Node, s.level)

	// fmt.Printf("start doSearch:%v\n", index)
	currentNode := s.head

	// Iterate from top level to bottom level.
	for l := s.level - 1; l >= 0; l-- {
		// Iterate value util value's index is >= given index.
		// The max iterate count is skip list's length. So the worst O(n) is N.
		for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].index < index {
			currentNode = currentNode.nextNodes[l]
		}

		// When next value's index is >= given index, add current value whose index < given index.
		previousNodes[l] = currentNode
	}

	// Avoid point to tail which will occur panic in Insert and Delete function.
	// When the next value is tail.
	// The index is larger than the maximum index in the skip list or skip list's length is 0. Don't point to tail.
	// When the next value isn't tail.
	// Next value's index must >= given index. Point to it.
	if currentNode.nextNodes[0] != s.tail {
		currentNode = currentNode.nextNodes[0]
	}
	// fmt.Printf("previous value:\n")
	// for _, n := range previousNodes {
	// 	fmt.Printf("%p\t", n)
	// }
	// fmt.Println()
	// fmt.Printf("end doSearch %v\n", index)

	return previousNodes, currentNode
}

// searchWithoutPreviousNodes will return the value whose index is given index.
// If can not find the given index, return nil.
// This function is faster than searchWithPreviousNodes and it used to only searching index.
func (s *skipList) searchValueByKey(index uint64) *Node {
	currentNode := s.head
	// Iterate from top level to bottom level.
	for l := s.level - 1; l >= 0; l-- {
		// Iterate value util value's index is >= given index.
		// The max iterate count is skip list's length. So the worst O(n) is N.
		for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].index < index {
			currentNode = currentNode.nextNodes[l]
		}
	}

	currentNode = currentNode.nextNodes[0]
	if currentNode == s.tail || currentNode.index > index {
		return nil
	} else if currentNode.index == index {
		return currentNode
	} else {
		return nil
	}
}

// insert will insert a value into skip list and update the length.
// If skip has these this index, overwrite the value, otherwise add it.
func (s *skipList) insert(index uint64, value interface{}) {
	// Write lock and unlock.
	previousNodes, currentNode := s.searchWithPreviousNodes(index)

	if currentNode != s.head && currentNode.index == index {
		currentNode.value = value
		return
	}

	// Make a new value.
	newNode := newNode(index, value, s.randomLevel())

	// Adjust pointer. Similar to update linked list.
	for i := len(newNode.nextNodes) - 1; i >= 0; i-- {
		// Firstly, new value point to next value.
		newNode.nextNodes[i] = previousNodes[i].nextNodes[i]

		// Secondly, previous nodes point to new value.
		previousNodes[i].nextNodes[i] = newNode

		// Finally, in order to release the slice, point to nil.
		previousNodes[i] = nil
	}

	atomic.AddInt32(&s.length, 1)

	for i := len(newNode.nextNodes); i < len(previousNodes); i++ {
		previousNodes[i] = nil
	}
}

// delete will find the index is existed or not firstly.
// If existed, delete it and update length, otherwise do nothing.
func (s *skipList) delete(index uint64) {
	// Write lock and unlock.
	previousNodes, currentNode := s.searchWithPreviousNodes(index)

	// If skip list length is 0 or could not find value with the given index.
	if currentNode != s.head && currentNode.index == index {
		// Adjust pointer. Similar to update linked list.
		for i := 0; i < len(currentNode.nextNodes); i++ {
			previousNodes[i].nextNodes[i] = currentNode.nextNodes[i]
			currentNode.nextNodes[i] = nil
			previousNodes[i] = nil
		}

		atomic.AddInt32(&s.length, -1)
	}

	for i := len(currentNode.nextNodes); i < len(previousNodes); i++ {
		previousNodes[i] = nil
	}
}

// snapshot will create a snapshot of the skip list and return a slice of the nodes.
func (s *skipList) snapshot() []*Node {
	result := make([]*Node, s.length)
	i := 0

	currentNode := s.head.nextNodes[0]
	for currentNode != s.tail {
		node := &Node{
			index:     currentNode.index,
			value:     currentNode.value,
			nextNodes: nil,
		}

		result[i] = node
		currentNode = currentNode.nextNodes[0]
		i++
	}

	return result
}

// getLength will return the length of skip list.
func (s *skipList) getLength() int32 {
	return atomic.LoadInt32(&s.length)
}

// randomLevel will generate and random level that level > 0 and level < skip list's level
// This comes from redis's implementation.
func (s *skipList) randomLevel() int {
	level := 1
	for rand.Float64() < 0.25 && level < s.level {
		level++
	}

	return level
}
