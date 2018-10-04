package algorithm

/**
top_k 问题，意思就是从一系列数字中选中排名前K的元素
*/

// 求解tok_5
// 采用快排选出 top5,
// 二分后，获取分割值的下标，与top比较
func QuickSortTopK(array []int, top int) []int {
	head, tail := 0, len(array)-1
	index := partition(array, head, tail) // 从大到小排序
	for index != top {                    //分割值的下标与top不相等
		if index > top { // 如果大于，说明index 之前的数据个数大于 top 个
			tail = index - 1
			index = partition(array, head, tail) // 继续分割，把前面的往后面移
		} else { // 说明index 之前的个数小于 top 个
			head = index + 1
			index = partition(array, head, tail) //继续分割把后面的往前移
		}
	}
	return array[0:top]
}
func partition(array []int, head int, tail int) int {
	mid := array[head]
	midIndex := head
	for head <= tail {
		if array[head] > mid {
			array[head], array[midIndex] = mid, array[head]
			head++
			midIndex++
		} else if array[head] < mid {
			array[head], array[tail] = array[tail], array[head]
			tail--
		} else {
			head++
		}
	}
	return midIndex
}
func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	mid := array[0]
	index := 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[index] > mid {
			array[index], array[tail] = array[tail], array[index]
			tail--
		} else {
			array[index], array[head] = array[head], array[index]
			head++
			index++
		}
	}
	quickSort(array[0:head])
	quickSort(array[head+1:])
	return array
}

// 选择排序查找top5
func SelectSortTopK(array []int, top int) []int {
	// 外层只需要循环top次，
	for i := 0; i < top; i++ {
		for j := 1; j < len(array)-i; j++ { //每次都能找出最大的一个
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return array
}

/**
原理是 先构建一个top 小顶堆，然后从下标数组top处，开始和小顶堆的堆头比较
如果大于的话，进行替换。替换完成后，再进行调整堆
*/
func HeadTopK(array []int, top int) []int {
	// 构建一个top堆
	for i := top/2 - 1; i >= 0; i-- {
		createHeap(array, i, top)
	}
	//调整堆
	adjustHeap2(array, top)
	for i := top; i < len(array); i++ {
		if array[i] > array[0] {
			array[i], array[0] = array[0], array[i]
			adjustHeap2(array, top)
		}
	}
	return array[0:top]
}
func adjustHeap2(array []int, top int) {
	for i := top - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		createHeap(array, 0, i)
	}
}

func createHeap(array []int, parentNode, top int) {
	childrenNode := parentNode*2 + 1
	for childrenNode < top {
		if (childrenNode+1) < top && array[childrenNode] < array[childrenNode+1] {
			childrenNode++
		}
		if array[parentNode] < array[childrenNode] {
			array[parentNode], array[childrenNode] = array[childrenNode], array[parentNode]
		}
		parentNode = childrenNode
		childrenNode = parentNode*2 + 1
	}
}
