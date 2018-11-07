package algorithm

import "testing"

func TestIntHeap_Len(t *testing.T) {
	heap := IntHeap{1, 3, 4}
	heap.Len()
}
