package algorithm

import (
	"fmt"
	"testing"
)

/**
左左，右右旋转
*/
func TestAvlTree_Insert(t *testing.T) {
	avlTree := &AvlTree{}

	avlTree.Insert(5)
	avlTree.Insert(3)
	avlTree.Insert(7)
	avlTree.Insert(2)
	avlTree.Insert(1)
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(0)
	PrePrint(avlTree.root)
	avlTree.Insert(8)
	fmt.Println()
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(9)
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(10)
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(11)
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(12)
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(-2)
	PrePrint(avlTree.root)
	fmt.Println()
	avlTree.Insert(-3)
	PrePrint(avlTree.root)
}

/**
左右旋转
*/
func TestAvlInsert2(t *testing.T) {
	avlTree := &AvlTree{}
	avlTree.Insert(5)
	avlTree.Insert(4)
	avlTree.Insert(7)
	avlTree.Insert(1)
	avlTree.Insert(2)

	PrePrint(avlTree.root)
}

/**
右左旋转
*/
func TestAvlInsert3(t *testing.T) {
	avlTree := &AvlTree{}
	avlTree.Insert(5)
	avlTree.Insert(4)
	avlTree.Insert(7)
	avlTree.Insert(9)
	avlTree.Insert(8)

	PrePrint(avlTree.root)
}

/**
这个avl的树实现有问题
*/
func TestGetHeight(t *testing.T) {
	avlTree := &AvlTree{}
	avlTree.Insert(8)
	avlTree.Insert(4)
	avlTree.Insert(15)
	avlTree.Insert(5)
	avlTree.Insert(6)

	PrePrint(avlTree.root)
	fmt.Println(GetHeight(avlTree.root))

}
