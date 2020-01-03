package algorithm

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
		createBigHeap(array, i, top)
	}
	//调整堆
	for i := top - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		createBigHeap(array, 0, i)
	}
	for i := top; i < len(array); i++ {
		if array[i] > array[0] {
			array[i], array[0] = array[0], array[i]
			for i := top - 1; i >= 0; i-- {
				array[0], array[i] = array[i], array[0]
				createBigHeap(array, 0, i)
			}
		}
	}
	return array[0:top]
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

/**
堆排序计算top10
*/
func HeapSortComputeTop10(fileAddress string) ([]int, error) {
	file, err := os.Open(fileAddress)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	data, _ := reader.ReadString('\n')
	array := strArrayConvertIntArray(data)
	//构造小顶堆
	array = CreateSmallHeap(array)
	//计算topN
	computeTopN(array, reader)
	// 调整堆,就是排序的过程，结过为从大到小
	for i := len(array) - 1; i >= 0; i-- {
		array[i], array[0] = array[0], array[i]
		AdjustSmallHeap(array, 0, i)
	}
	return array, nil
}

func computeTopN(array []int, reader *bufio.Reader) []int {
	for {
		data, _ := reader.ReadString('\n')
		//数据已处理到结尾
		if len(data) == 0 {
			break
		}
		temp := strArrayConvertIntArray(data)
		for _, v := range temp {
			//比小顶堆堆顶元素还小
			if v < array[0] {
				continue
			}
			//否则，替换堆顶元素，重新调整堆
			array[0] = v
			array = CreateSmallHeap(array)
		}
	}
	return array
}

//字符串数组转int数组
func strArrayConvertIntArray(s string) []int {
	str := strings.Split(s, ",")
	array := make([]int, len(str))
	for k, v := range str {
		temp, _ := strconv.Atoi(v)
		array[k] = temp
	}
	return array
}

func CreateSmallHeap(intArray []int) []int {
	length := len(intArray)
	//建堆,把最大的放到顶部
	for i := length/2 + 1; i >= 0; i-- {
		adjustSmallHeap(intArray, i, length)
	}
	return intArray
}

func AdjustSmallHeap(intArray []int, i, length int) {
	//左子节点
	childrenNode := i*2 + 1
	for childrenNode < length {
		// 从子节点中找较小的
		if childrenNode+1 < length && intArray[childrenNode] > intArray[childrenNode+1] {
			childrenNode++
		}
		// 如果父节点 < 两个子节点中最小的，则说明父节点最小
		if intArray[i] > intArray[childrenNode] {
			// 把父节点和最小子节点交换
			intArray[i], intArray[childrenNode] = intArray[childrenNode], intArray[i]
		}
		// 父节点变更
		i = childrenNode
		// 父节点的子节点
		childrenNode = i*2 + 1
	}
}
