package algorithm

import (
	"testing"
)

func TestIntHeap_Len(t *testing.T) {
	heap := &IntHeap{}

	heap.Push(1)
	heap.Push(4)
	heap.Push(2)
	heap.Push(5)
	heap.Push(6)
	heap.Push(6)
	heap.Push(5)
	heap.Push(0)
	t.Log(heap)
	t.Log(heap.Pop())
	t.Log(heap)
	t.Log(heap.Pop())
	t.Log(heap)
	t.Log(heap.Pop())
	t.Log(heap)

}
