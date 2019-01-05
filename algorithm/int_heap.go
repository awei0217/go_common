package algorithm

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

/**
构造一个小顶堆 ，数组下标为0的元素始终是最小的
*/
func (h *IntHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
	l := len(*h) - 1
	for (l-1)/2 >= 0 && (*h)[l] < (*h)[(l-1)/2] {
		h.Swap(l, (l-1)/2)
		l = (l - 1) / 2
	}
}

/**
返回堆顶的一个元素，并删除,重新调整堆，保证堆定的元素始终是最小的
*/
func (h *IntHeap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}
	x := (*h)[0]
	if h.Len() == 1 {
		h = &IntHeap{}
		return x
	}
	(*h)[0] = (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	adjustHelp(h)
	return x
}

func adjustHelp(h *IntHeap) {
	min := 0
	for i := 0; i < h.Len(); {
		if i*2+1 < h.Len() && !h.Less(i, i*2+1) { //顶节点 > 左节点
			min = i*2 + 1
		}
		if i*2+2 < h.Len() && h.Less(i*2+2, i*2+1) { //右节点 < 左节点
			min = i*2 + 2
		}
		if min == i {
			break
		}
		h.Swap(i, min)
		i = min
	}
}
