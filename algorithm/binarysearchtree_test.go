package algorithm

import (
	"testing"
)

func TestBinarySearchTree_RecursionAdd(t *testing.T) {
	tree := &BinarySearchTree{0, nil, nil}
	for i := 1; i < 10; i++ {
		tree.RecursionAdd(i)
	}
	t.Log(tree)
	t.Log(tree.leftChild)
	t.Log(tree.leftChild.leftChild)
	t.Log(tree.rightChild)
	t.Log(tree.rightChild.rightChild)
	t.Log(tree.Deep())
	t.Log(tree.Height())
}
func TestBinarySearchTree_NotRecursionAdd(t *testing.T) {
	tree := &BinarySearchTree{0, nil, nil}
	for i := 1; i < 10; i++ {
		tree.RecursionAdd(i)
	}
	t.Log(tree.Search(5).value)
	t.Log(tree.Search(5).leftChild.value)
	t.Log(tree.GetMin())
	tree.Remove(1)
	t.Log(tree.Search(1))
	t.Log(tree.Deep())
	t.Log(tree.Search(8).Deep())
	t.Log(tree.Search(8).Height())
}

func TestPreTraversalTree(t *testing.T) {
	PreTraversalTree([]int{0,1,2,3,4,5})
}
func TestPreTraversalRecursionTree(t *testing.T) {
	PreTraversalRecursionTree([]int{0,1,2,3,4,5})
}
func TestInTraversalTree(t *testing.T) {
	InTraversalTree([]int{0,1,2,3,4,5})
}
func TestInTraversalRecursionTree(t *testing.T) {
	InTraversalRecursionTree([]int{0,1,2,3,4,5})
}
func TestPoTraversalTree(t *testing.T) {
	PoTraversalTree([]int{0,1,2,3,4,5})
}
func TestPoTraversalRecursionTree(t *testing.T) {
	PoTraversalRecursionTree([]int{0,1,2,3,4,5})
}
