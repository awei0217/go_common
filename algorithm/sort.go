package algorithm

/**
（从小到大排序）存在10个不同大小的气泡，由底至上地把较少的气泡逐步地向上升，这样经过遍历一次后，最小的气泡就会被上升到顶（下标为0），然后再从底至上地这样升，循环直至十个气泡大小有序。
在冒泡排序中，最重要的思想是两两比较，将两者较少的升上去
冒泡排序最坏情况的时间复杂度是O(n²)
 */
func MaoPaoSort(intArray []int) []int  {
	for i:=0;i<len(intArray)-1;i++ {
		for j:=0;j<len(intArray)-1-i;j++{
			if intArray[j] < intArray[j+1] {
				intArray[j],intArray[j+1] = intArray[j+1],intArray[j]
			}
		}
	}
	return intArray;
}
/**
选择排序
 */
func SelectSort(intArray []int) []int{
	for i:=0;i<len(intArray);i++{
		var minIndex int = i
		for j:=i;j<len(intArray) ;j++  {
			if(intArray[minIndex]> intArray[j]){
				minIndex = j
			}
		}
		intArray[i],intArray[minIndex] = intArray[minIndex],intArray[i]
	}
	return intArray;
}
// 快速排序
func QuickSort(array []int) []int {
	if len(array) <=1 {
		return array
	}
	medianValue,index := array[0],1
	head,tail := 0,len(array)-1
	for head <tail {
		if array[index] >medianValue {
			array[index],array[tail] = array[tail],array[index]
			tail--
		}else{
			array[index],array[head] = array[head],array[index]
			head++
			index++
		}
	}
	QuickSort(array[:head])
	QuickSort(array[head+1:])
	return array
}

//插入排序
func InsertSort(intArray []int) []int {
	for i:=0;i<len(intArray);i++ {
		preIndex := i - 1;
		current := intArray[i];
		for (preIndex >= 0 && intArray[preIndex] > current) {
			intArray[preIndex + 1] = intArray[preIndex];
			preIndex--;
		}
		intArray[preIndex + 1] = current;
	}
	return intArray

}
//希尔排序
func ShellSort(intArray []int)[]int  {
	d := len(intArray) / 2
	for d>=1{
		for i:=d;i<len(intArray) ;i++  {
			j := i
			for j-d >= 0 && intArray[j] < intArray[j-d] {
				intArray[j], intArray[j-d] = intArray[j-d], intArray[j]
				j -= d
			}
		}
		d = d/2
	}
	return intArray
}


