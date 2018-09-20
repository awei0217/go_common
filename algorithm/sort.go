package algorithm

/**
稳定排序：待排序的记录序列中可能存在两个或两个以上关键字相等的记录。排序前的序列中Ri领先于Rj（即i<j）.若在排序后的序列中Ri仍然领先于Rj，则称所用的方法是稳定的
		  插入排序  ，基数排序  ，归并排序  ，冒泡排序 ，计数排序
不稳定的排序算法有：快速排序，希尔排序，简单选择排序，堆排序
 */


/**
（从小到大排序）存在10个不同大小的气泡，由底至上地把较少的气泡逐步地向上升，这样经过遍历一次后，最小的气泡就会被上升到顶（下标为0），然后再从底至上地这样升，循环直至十个气泡大小有序。
在冒泡排序中，最重要的思想是两两比较，将两者较少的升上去
冒泡排序最坏情况的时间复杂度是O(n²)
*/
func MaoPaoSort(intArray []int) []int {
	for i := 0; i < len(intArray)-1; i++ {
		for j := 0; j < len(intArray)-1-i; j++ {
			if intArray[j] < intArray[j+1] {
				intArray[j], intArray[j+1] = intArray[j+1], intArray[j]
			}
		}
	}
	return intArray
}

/**
选择排序
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是每一次从待排序的数据元素中选出最小（或最大）的一个元素，
存放在序列的起始位置，直到全部待排序的数据元素排完。 选择排序是不稳定的排序方法。
时间复杂度为 O(N^2)
*/
func SelectSort(intArray []int) []int {
	for i := 0; i < len(intArray); i++ {
		minIndex := i
		for j := i; j < len(intArray); j++ {
			if intArray[minIndex] > intArray[j] {
				minIndex = j
			}
		}
		intArray[i], intArray[minIndex] = intArray[minIndex], intArray[i]
	}
	return intArray
}

/**
快速排序:
快速排序由C. A. R. Hoare在1962年提出。它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，
其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，
整个排序过程可以递归进行，以此达到整个数据变成有序序列
快速排序时间复杂度
快速排序的时间复杂度在最坏情况下是O(N2)，平均的时间复杂度是O(N*lgN)。
这句话很好理解：假设被排序的数列中有N个数。遍历一次的时间复杂度是O(N)，需要遍历多少次呢？至少lg(N+1)次，最多N次。
(01) 为什么最少是lg(N+1)次？快速排序是采用的分治法进行遍历的，我们将它看作一棵二叉树，它需要遍历的次数就是二叉树的深度，而根据完全二叉树的定义，它的深度至少是lg(N+1)。因此，快速排序的遍历次数最少是lg(N+1)次。
(02) 为什么最多是N次？这个应该非常简单，还是将快速排序看作一棵二叉树，它的深度最大是N。因此，快读排序的遍历次数最多是N次。
 */
func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	medianValue, index := array[0], 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[index] > medianValue {
			array[index], array[tail] = array[tail], array[index]
			tail--
		} else {
			array[index], array[head] = array[head], array[index]
			head++
			index++
		}
	}
	QuickSort(array[:head])
	QuickSort(array[head+1:])
	return array
}

/**
插入排序
有一个已经有序的数据序列，要求在这个已经排好的数据序列中插入一个数，但要求插入后此数据序列仍然有序，
这个时候就要用到一种新的排序方法——插入排序法,插入排序的基本操作就是将一个数据插入到已经排好序的有序数据中，
从而得到一个新的、个数加一的有序数据，算法适用于少量数据的排序，时间复杂度为O(n^2)。是稳定的排序方法
 */
func InsertSort(intArray []int) []int {
	for i := 0; i < len(intArray); i++ {
		preIndex := i - 1
		current := intArray[i]
		for preIndex >= 0 && intArray[preIndex] > current {
			intArray[preIndex+1] = intArray[preIndex]
			preIndex--
		}
		intArray[preIndex+1] = current
	}
	return intArray

}

//希尔排序
func ShellSort(intArray []int) []int {
	d := len(intArray) / 2
	for d >= 1 {
		for i := d; i < len(intArray); i++ {
			j := i
			for j-d >= 0 && intArray[j] < intArray[j-d] {
				intArray[j], intArray[j-d] = intArray[j-d], intArray[j]
				j -= d
			}
		}
		d = d / 2
	}
	return intArray
}

/**
堆排序
时间复杂度：O(nlgn)
建堆的时间复杂度是 O(n)
调整堆的时间复杂度是 O(lgn)
 */
func HeapSort(intArray []int)[]int{

	length := len(intArray)
	//建堆,把最大的放到顶部
	for i:= length /2-1;i>=0;i--{
		adjustHeap(intArray,i,length)
	}
	// 调整堆
	for i:=length-1;i>=0;i--{
		intArray[i],intArray[0] = intArray[0],intArray[i]
		adjustHeap(intArray,0,i)
	}
	return intArray
}
func adjustHeap(intArray[]int,i,length int){
	//左子节点
	childrenNode := i*2+1
	for childrenNode< length{
		if childrenNode+1 <length && intArray[childrenNode] > intArray[childrenNode+1] { // 从子节点中找较小的
			childrenNode++
		}
		if intArray[i] < intArray[childrenNode]{ // 如果父节点 < 两个子节点中最小的，则说明父节点最小
			break
		}
		// 把父节点和最小子节点交换
		intArray[i],intArray[childrenNode] = intArray[childrenNode],intArray[i]
		// 父节点变更
		i =  childrenNode
		// 父节点的子节点
		childrenNode = i * 2 + 1
	}
}
/**
排序算法	平均时间复杂度	最坏时间复杂度	空间复杂度	是否稳定
冒泡排序	O（n2）	O（n2）	O（1）	是
选择排序	O（n2）	O（n2）	O（1）	不是
直接插入排序	O（n2）	O（n2）	O（1）	是
归并排序	O(nlogn)	O(nlogn)	O（n）	是
快速排序	O(nlogn)	O（n2）	O（logn）	不是
堆排序	O(nlogn)	O(nlogn)	O（1）	不是
希尔排序	O(nlogn)	O（ns）	O（1）	不是
计数排序	O(n+k)	O(n+k)	O(n+k)	是
基数排序	O(N∗M)	O(N∗M)	O(M)	是

 */
