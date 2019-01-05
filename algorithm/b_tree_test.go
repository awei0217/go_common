package algorithm

import (
	"fmt"
	"testing"
)

func TestNewBTree(t *testing.T) {
	bt := NewBTree(4)

	for i := 1; i < 10; i++ {
		bt.Insert(i)
	}
	fmt.Println(bt)
}
