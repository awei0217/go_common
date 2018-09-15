package algorithm

import (
	"testing"
)

func TestBinarySearchTree_RecursionAdd(t *testing.T) {
	tree := &BinarySearchTree{0,nil,nil}
	for i:=1;i<10;i++{
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
	tree := &BinarySearchTree{0,nil,nil}
	for i:=1;i<10;i++{
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
