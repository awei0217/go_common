package algorithm

/**
常见的时间复杂度有：常数阶O(时间复杂度和空间复杂度)，对数阶O(log2n)，线性阶O(n)，线性对数阶O(nlog2n)，平方阶O(n2)，立方阶O(n3)， k次方阶O(nk)，指数阶O(2n)。
随着问题规模n的不断增大，上述时间复杂度不断增大，算法的执行效率越低。

稳定排序：待排序的记录序列中可能存在两个或两个以上关键字相等的记录。排序前的序列中Ri领先于Rj（即i<j）.若在排序后的序列中Ri仍然领先于Rj，则称所用的方法是稳定的
		  插入排序  ，基数排序  ，归并排序  ，冒泡排序 ，计数排序
不稳定的排序算法有：快速排序，希尔排序，简单选择排序，堆排序

排序算法	平均时间复杂度	最坏时间复杂度	空间复杂度	是否稳定
冒泡排序	O（n2）	        O（n2）	        O（时间复杂度和空间复杂度）	是
选择排序	O（n2）	        O（n2）	        O（时间复杂度和空间复杂度）	不是
插入排序    O（n2）	        O（n2）	        O（时间复杂度和空间复杂度）	是
归并排序	O(nlogn)	    O(nlogn)	    O（n）	是
快速排序	O(nlogn)	    O（n2）	        O（logn）	不是
堆排序	    O(nlogn)	    O(nlogn)	    O（时间复杂度和空间复杂度）	不是
希尔排序	O(nlogn)	    O（ns）	        O（时间复杂度和空间复杂度）	不是
计数排序	O(n+k)	        O(n+k)	        O(n+k)	是
基数排序	O(N∗M)	        O(N∗M)	        O(M)	是
*/

/**
存在10个不同大小的气泡，由底至上地把较少的气泡逐步地向上升，这样经过遍历一次后，最小的气泡就会被上升到顶（下标为0），然后再从底至上地这样升，循环直至十个气泡大小有序。
在冒泡排序中，最重要的思想是两两比较，将两者较少的升上去
冒泡排序最坏时间复杂度是O(n²) 最好时间复杂度是O(n) 平均时间复杂度是 O(n²)
稳定排序
*/
func MaoPaoSort(intArray []int) []int {
	for i := 0; i < len(intArray); i++ {
		flag := true
		for j := 0; j < len(intArray)-1-i; j++ {
			if intArray[j] > intArray[j+1] {
				flag = false
				intArray[j], intArray[j+1] = intArray[j+1], intArray[j]
			}
		}
		if flag {
			break
		}
	}
	return intArray
}

/**
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是每一次从待排序的数据元素中选出最小（或最大）的一个元素，
存放在序列的起始位置，直到全部待排序的数据元素排完。
最好，最坏，平均时间复杂度都是 O(n2)
不稳定排序
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
这句话很好理解：假设被排序的数列中有N个数。遍历一次的时间复杂度是O(N)，需要遍历多少次呢？至少lg(N+时间复杂度和空间复杂度)次，最多N次。
(01) 为什么最少是lg(N+时间复杂度和空间复杂度)次？快速排序是采用的分治法进行遍历的，我们将它看作一棵二叉树，它需要遍历的次数就是二叉树的深度，而根据完全二叉树的定义，它的深度至少是lg(N+时间复杂度和空间复杂度)。因此，快速排序的遍历次数最少是lg(N+时间复杂度和空间复杂度)次。
(02) 为什么最多是N次？这个应该非常简单，还是将快速排序看作一棵二叉树，它的深度最大是N。因此，快读排序的遍历次数最多是N次。
*/
func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	head, tail := 0, len(array)-1
	par := partitionArray(array, head, tail) // 每次执行n次
	// 需要递归logn次 ，最坏时需要递归n次
	QuickSort(array[:par])
	QuickSort(array[par+1:])
	return array
}

/**
快速排序非递归 {1, 3, -9, 6, 8, -19, 20, -20}
*/
func QuickSortNotRecursion(array []int) []int {

	head, tail := 0, len(array)-1
	stack := &ArrayStack{}
	par := partitionArray(array, head, tail)
	if par > head+1 {
		stack = Push(stack, head)
		stack = Push(stack, par-1)
	}
	if par < tail-1 {
		stack = Push(stack, par+1)
		stack = Push(stack, tail)
	}
	for !IsEmpty(stack) {
		tail = Pop(stack).(int)
		head = Pop(stack).(int)
		par = partitionArray2(array, head, tail)
		//par = partitionArray(array, head, tail)
		if par > head+1 {
			stack = Push(stack, head)
			stack = Push(stack, par-1)
		}
		if par < tail-1 {
			stack = Push(stack, par+1)
			stack = Push(stack, tail)
		}
	}
	return array
}

/**
这种分区函数下是稳定排序
*/
func partitionArray(array []int, head int, tail int) int {
	medianValue, index := array[head], head+1
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
	return head
}

/**
这种分区函数下是不稳定排序
*/
func partitionArray2(array []int, head int, tail int) int {
	middleValue := array[head]
	for head < tail {
		for head < tail && array[tail] > middleValue {
			tail--
		}
		array[tail], array[head] = array[head], array[tail]
		for head < tail && array[head] < middleValue {
			head++
		}
		array[head], array[tail] = array[tail], array[head]
	}
	return head
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

/**
希尔排序,最重要的是选取增量序列
*/
func ShellSort(intArray []int) []int {
	d := len(intArray) / 2
	for d > 0 {
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
时间复杂度_空间复杂度：O(nlgn)
建堆的时间复杂度是 O(n)
调整堆的时间复杂度是 O(lgn)
*/
func SmallHeapSort(intArray []int) []int {

	length := len(intArray)
	//建堆,把最大的放到顶部
	for i := length/2 + 1; i >= 0; i-- {
		adjustSmallHeap(intArray, i, length)
	}
	// 调整堆
	for i := length - 1; i >= 0; i-- {
		intArray[i], intArray[0] = intArray[0], intArray[i]
		adjustSmallHeap(intArray, 0, i)
	}
	return intArray
}
func adjustSmallHeap(intArray []int, i, length int) {
	//左子节点
	childrenNode := i*2 + 1
	for childrenNode < length {
		if childrenNode+1 < length && intArray[childrenNode] > intArray[childrenNode+1] { // 从子节点中找较小的
			childrenNode++
		}
		if intArray[i] > intArray[childrenNode] { // 如果父节点 < 两个子节点中最小的，则说明父节点最小
			// 把父节点和最小子节点交换
			intArray[i], intArray[childrenNode] = intArray[childrenNode], intArray[i]
		}
		// 父节点变更
		i = childrenNode
		// 父节点的子节点
		childrenNode = i*2 + 1
	}
}

/**
大顶堆
*/
func BigHeapSort(array []int) []int {
	// 创建堆
	for i := len(array)/2 - 1; i >= 0; i-- {
		createBigHeap(array, i, len(array))
	}
	//调整堆
	for i := len(array) - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		createBigHeap(array, 0, i)

	}
	return array
}
func createBigHeap(array []int, i, length int) {
	childLeft := i*2 + 1
	for childLeft < length {
		if (childLeft+1) < length && array[childLeft] < array[childLeft+1] {
			childLeft++
		}
		if array[childLeft] > array[i] {
			array[i], array[childLeft] = array[childLeft], array[i]
		}
		i = childLeft
		childLeft = i*2 + 1
	}
}

/**
,该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，
得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为二路归并。
二路归并排序 递归
稳定排序
时间复杂度 O(nlogn)
空间复杂度 O(n)
*/
func MergeSort(array []int) []int {
	length := len(array)
	if length <= 1 {
		return array
	}
	num := length / 2
	left := MergeSort(array[:num])
	right := MergeSort(array[num:])
	return merge(left, right)
}

/**
二路归并排序 非递归
*/
func MergerSortNotRecursion(array []int) []int {
	length := len(array)
	if length <= 1 {
		return array
	}
	i := 1 // 子序列大小
	result := make([]int, 0)
	for i < length {
		j := 0
		for j < length {
			if j+2*i > length {
				result = merge(array[j:j+i], array[j+i:length])
				index := j
				for _, v := range result {
					array[index] = v
					index++
				}
			} else {
				result = merge(array[j:j+i], array[j+i:j+2*i])
				index := j
				for _, v := range result {
					array[index] = v
					index++
				}
			}
			j = j + 2*i
		}
		i = i << 1 // 子序列大小翻倍
	}
	return array
}

/**
对两个有序的数组进行合并
*/
func merge(left []int, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

/**
桶排序
适合不重复数组排序，这数组的最大值不宜过大
*/
func BucketSort(array []int) []int {
	var max int
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	max = max + 1
	temp := make([][]int, max, max)
	bucketSizes := make([]int, max, max)
	for i := 0; i < max; i++ {
		temp[i] = make([]int, max, max)
	}
	for _, v := range array {
		mod := v % max
		i := bucketSizes[v%max]
		temp[mod][i] = v
		i++
	}
	index := 0
	for _, v := range temp {
		if v[0] != 0 {
			array[index] = v[0]
			index++
		}
	}
	return array
}

/**
计数排序
入参 是数组 和 数组中的最大值
*/
func CountSort(array []int, max int) []int {
	length := len(array)
	if length == 0 || length == 1 {
		return array
	}
	temp := make([]int, max+1)
	for i := 0; i < length; i++ {
		temp[array[i]]++
	}
	for i := 1; i < max+1; i++ {
		temp[i+1] = temp[i-1] + temp[i]
	}
	sort := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		index := temp[array[i]] - 1
		sort[index] = array[i]
		temp[array[i]]--
	}
	return array
}

/**
按字符串长度，由长到短
字符串合并排序
*/
func StringMergerSort(d []string) []string {
	if len(d) <= 1 {
		return d
	}
	l := len(d) / 2
	left := StringMergerSort(d[:l])
	right := StringMergerSort(d[l:])
	return merger(left, right)
}

func merger(d1, d2 []string) []string {
	l := 0
	r := 0
	result := make([]string, 0, 0)
	for l < len(d1) && r < len(d2) {
		if len(d1[l]) < len(d2[r]) {
			result = append(result, d2[r])
			r++
		} else if len(d1[l]) > len(d2[r]) {
			result = append(result, d1[l])
			l++
		} else {
			result = append(result, d1[l], d2[r])
			l++
			r++
		}
	}
	result = append(result, d1[l:]...)
	result = append(result, d2[r:]...)
	return result
}

func quickSort3(ints []int, i int, i2 int) {
	if i >= i2 {
		return
	}
	mid := partition3(ints, i, i2)
	quickSort3(ints, i, mid-1)
	quickSort3(ints, mid+1, i2)
}
func partition3(ints []int, i int, i2 int) int {
	midValue := ints[i]
	index := i + 1
	for i < i2 {
		if ints[index] > midValue {
			ints[index], ints[i2] = ints[i2], ints[index]
			i2--
		} else {
			ints[index], ints[i] = ints[i], ints[index]
			i++
			index++
		}
	}
	return i
}

//合并区间
func mergeIntervals(intervals [][]int) [][]int {
	//先判断，逻辑严谨性
	if len(intervals) == 0 {
		return nil
	}
	//快速排序
	QuickSort3(intervals)

	//合并区间
	index := 0
	result := make([][]int, 0)
	for index < len(intervals) {
		//获取区间的左右值
		left := intervals[index][0]
		right := intervals[index][1]

		//判断下一个区间的起始值是否在上一个区间的范围
		//index下标小于二位数组长度-1，因为下一个区间是index+1，
		//已经排过序了，下一个区间的起始值只需要小于上一个区间的最大值
		for index < len(intervals)-1 && intervals[index+1][0] <= right {
			//获取区间的最大值
			right = max(intervals[index][1], intervals[index+1][1])
			index++
		}
		result = append(result, []int{left, right})
		index++
	}
	return result
}

func QuickSort3(intervals [][]int) {
	//递归终止条件
	if len(intervals) <= 1 {
		return
	}
	head, tail := 0, len(intervals)-1
	//分区函数排序
	index := partion3(intervals, head, tail)
	//递归左边的数组
	QuickSort3(intervals[:index])
	//递归右边的数组
	QuickSort3(intervals[index+1:])
}

func partion3(intervals [][]int, head, tail int) int {
	//选取一个中间值
	v := intervals[head][0]
	for head < tail {
		//比较末尾值和中间值，末尾值比中间值大，循环倒数第二个
		for head < tail && intervals[tail][0] >= v {
			tail--
		}
		//当末尾值大于中间值时，交换末尾值和中间值
		intervals[head], intervals[tail] = intervals[tail], intervals[head]
		//比较首部值和中间值，首部值比中间值小，继续下一个
		for head < tail && intervals[head][0] <= v {
			head++
		}
		//当首部值大于中间值时，交换中间值和首部值
		intervals[tail], intervals[head] = intervals[head], intervals[tail]
	}
	//返回中间值的下标
	return head

}

func ReorganizeString(S string) string {
	if len(S) == 0 {
		return ""
	}
	m := make(map[rune]int, 0)
	//统计每个字符出现个数，放入kv
	for _, c := range S {
		if _, ok := m[c]; ok {
			m[c]++
			//当有字符出现的个数大于字符串长度一半以上，则可肯定不满足规则
			if (m[c] - len(S)/2) > len(S)%2 {
				return ""
			}
		} else {
			m[c] = 1
		}
	}
	//按照kv组装成结构体数组
	cs := make([]StrCount, 0)
	for k, v := range m {
		cs = append(cs, StrCount{v, k})
	}

	//按照出现的个数排序(插入排序)
	for i, _ := range cs {
		preIndex := i - 1
		current := cs[i]
		for preIndex >= 0 && current.count > cs[preIndex].count {
			cs[preIndex+1] = cs[preIndex]
			preIndex--
		}
		cs[preIndex+1] = current
	}
	result := make([]rune, len(S))
	next := 0
	for _, v := range cs {
		for i := 0; i < v.count; i++ {
			if next >= len(result) {
				next = 1
			}
			result[next] = v.r
			next = next + 2
		}
	}
	return string(result)

}

type StrCount struct {
	count int
	r     rune
}
