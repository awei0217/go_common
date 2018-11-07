package algorithm

import "testing"

func TestPush(t *testing.T) {
	stack := Push(nil, 1)
	stack = Push(stack, 2)
	t.Log(stack.Size)
	Pop(stack)
	t.Log(stack.Size)
}

func TestPeek(t *testing.T) {
	t.Log(Peek(Push(&ArrayStack{}, 1)))
}

func TestIsEmpty(t *testing.T) {
	t.Log(IsEmpty(nil))
	t.Log(IsEmpty(&ArrayStack{}))
	t.Log(IsEmpty(Push(&ArrayStack{}, 1)))
}

func TestNextGreaterElement(t *testing.T) {
	t.Log(NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

func TestCalPoints(t *testing.T) {
	t.Log(CalPoints([]string{"5", "2", "C", "D", "+"}))
}
