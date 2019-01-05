package algorithm

import (
	"fmt"
	"testing"
)

func TestAvlInsert(t *testing.T) {
	avlTree := AvlInsert(nil, 8)
	avlTree = AvlInsert(avlTree, 4)
	avlTree = AvlInsert(avlTree, 15)
	avlTree = AvlInsert(avlTree, 5)
	avlTree = AvlInsert(avlTree, 6)

	fmt.Println(avlTree.high)
	t.Log(getAvlHeight(avlTree))
	//MidOrder(avlTree)
	//PoOrder(avlTree)
}
