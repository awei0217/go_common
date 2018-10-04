package algorithm

import "fmt"

type BinarySearchTree struct {
	value      int
	leftChild  *BinarySearchTree
	rightChild *BinarySearchTree
}

//递归添加
func (tree *BinarySearchTree) RecursionAdd(value int) {
	if tree.value == 0 {
		tree.value = value
		tree.leftChild = &BinarySearchTree{}
		tree.rightChild = &BinarySearchTree{}
	} else {
		if value < tree.value {
			tree.leftChild.RecursionAdd(value)
		} else {
			tree.rightChild.RecursionAdd(value)
		}
	}
}

/**
非递归添加
*/
func (tree *BinarySearchTree) NotRecursionAdd(value int) {
	if tree.value == 0 {
		tree.value = value
		tree.leftChild = &BinarySearchTree{}
		tree.rightChild = &BinarySearchTree{}
	} else {
		current := tree
		for {
			if value < current.value {
				current = current.leftChild
			} else if value > current.value {
				current = current.rightChild
			} else {
				return
			}
			if current == nil {
				current = &BinarySearchTree{value, nil, nil}
				break
			}
		}
	}
}

func (tree *BinarySearchTree) Search(value int) *BinarySearchTree {
	current := tree
	if current.value == value {
		return current
	}
	for {
		if current == nil {
			return nil
		}
		if value < current.value {
			current = current.leftChild
		} else if value > current.value {
			current = current.rightChild
		} else {
			return current
		}
	}

}

func (tree *BinarySearchTree) Remove(value int) {
	t := tree.Search(value)
	if t == nil {
		panic("value is not exist")
	}
	if t.leftChild.value == 0 && t.rightChild.value == 0 {
		destroy(t)
	} else if t.leftChild.value != 0 && t.rightChild.value != 0 { // Both LeftChild and RightChild are not empty.
		// Get the min-element in its RightChild.
		min := t.rightChild.GetMin()
		// Replace its value with the min-element's value.
		t.value = min.value
		// Remove the min-element.
		destroy(min)
	} else { // LeftChild or RightChild is empty.
		// Replace it with its child which is not empty.
		if t.leftChild.value > 0 {
			t.replaceWith(t.leftChild)
		} else {
			t.replaceWith(t.rightChild)
		}
	}
}
func (tree *BinarySearchTree) replaceWith(t *BinarySearchTree) {
	tree.value = t.value
	tree.leftChild = t.leftChild
	tree.rightChild = t.rightChild
	destroy(t)
}
func (tree *BinarySearchTree) GetMin() *BinarySearchTree {
	if tree.leftChild.value == 0 {
		return tree
	} else {
		return tree.leftChild.GetMin()
	}
}
func destroy(tree *BinarySearchTree) {
	tree.value = 0
	tree.leftChild = nil
	tree.rightChild = nil
}

/**
获取树的深度
百度百科定义：树的高度或深度：树中节点的最大层次；
*/
func (tree *BinarySearchTree) Deep() int {
	leftDeep := 0
	rightDeep := 0

	if tree.leftChild.value == 0 && tree.rightChild.value == 0 {
		// 空树返回-1
		return -1
	}
	current := tree
	for current.leftChild.value != 0 {
		current = current.leftChild
		leftDeep++
	}
	current = tree
	for current.rightChild.value != 0 {
		current = current.rightChild
		rightDeep++
	}
	if current.leftChild.value != 0 {
		rightDeep++
	}
	if rightDeep > leftDeep {
		return rightDeep + 1
	}
	return leftDeep + 1
}

func (tree *BinarySearchTree) Height() int {

	return tree.Deep()
}

/**
假设该数组满足二叉树
先序非递归遍历的方式打印该数组非递归
时间复杂度O(n) 空间复杂度O(logn)
*/
func PreTraversalTree(array []int) {
	stack := &LinkStack{}
	length := len(array)
	index := 0
	for index < length || stack.Size() > 0 {
		if index < length {
			fmt.Println(array[index])
			stack.Push(index)
			index = index*2 + 1
		} else {
			index = stack.Pop().(int)
			index = index*2 + 2
		}
	}
}

/**
中序遍历非递归
*/
func InTraversalTree(array []int) {
	stack := &LinkStack{}
	length := len(array)
	index := 0
	for index < length || stack.Size() > 0 {
		if index < length {
			stack.Push(index)
			index = index*2 + 1
		} else {
			index = stack.Pop().(int)
			fmt.Println(array[index])
			index = index*2 + 2
		}
	}
}

/**
后续遍历非递归
*/
func PoTraversalTree(array []int) {
	length := len(array)
	stack := &LinkStack{}
	index := 0
	preIndex := -1
	for index < length {
		stack.Push(index)
		index = index*2 + 1
	}
	for stack.Size() > 0 {
		index = stack.Pop().(int)
		if (index*2+1) < length && (index*2+2) < length && (index*2+2) != preIndex {
			stack.Push(index)
			index = index*2 + 2
			for index < length {
				stack.Push(index)
				index = index*2 + 1
			}
		} else {
			fmt.Println(array[index])
			preIndex = index
		}
	}
}

func PreTraversalRecursionTree(array []int) {
	index := 0
	PrintPreArray(index, len(array), array)
}

/**
先序递归遍历
*/
func PrintPreArray(index, length int, array []int) {
	if index < length {
		fmt.Println(array[index])
		PrintPreArray(index*2+1, length, array)
		PrintPreArray(index*2+2, length, array)
	}
}

func InTraversalRecursionTree(array []int) {
	index := 0
	PrintInArray(index, len(array), array)
}

/**
中序递归遍历
*/
func PrintInArray(index, length int, array []int) {
	if index < length {
		PrintInArray(index*2+1, length, array)
		fmt.Println(array[index])
		PrintInArray(index*2+2, length, array)
	}
}

func PoTraversalRecursionTree(array []int) {
	index := 0
	PrintPoArray(index, len(array), array)
}

/**
后序递归遍历
*/
func PrintPoArray(index, length int, array []int) {
	if index < length {
		PrintPoArray(index*2+1, length, array)
		PrintPoArray(index*2+2, length, array)
		fmt.Println(array[index])
	}
}
