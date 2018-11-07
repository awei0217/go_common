package algorithm

type IntHeap []int

func (h *IntHeap) Len() int {
	return 0
}

func (h *IntHeap) Less(i, j int) bool {
	return false
}

func (h *IntHeap) Swap(i, j int) {

}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
