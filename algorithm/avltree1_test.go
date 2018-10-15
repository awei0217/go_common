package algorithm

import "testing"

func TestAvlInsert(t *testing.T) {
	avlTree := AvlInsert(nil, 5)
	avlTree = AvlInsert(avlTree, 3)
	avlTree = AvlInsert(avlTree, 7)
	avlTree = AvlInsert(avlTree, 2)
	avlTree = AvlInsert(avlTree, 1)
	PreOrder(avlTree)
}
