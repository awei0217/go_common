package algorithm

import "testing"

func TestLinkStack_Push(t *testing.T) {
	stack := &LinkStack{nil, 0}

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for {
		t.Log(stack.Pop())
	}

}
