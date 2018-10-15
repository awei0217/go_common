package algorithm

import (
	"testing"
)

func Test_MaoPaoSort(t *testing.T) {
	intArray := MaoPaoSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func Test_SelectSort(t *testing.T) {
	intArray := SelectSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func TestQuickSort(t *testing.T) {
	t.Log(QuickSort([]int{5,4,3,2,1,6}))
}

func Test_InsertSort(t *testing.T) {
	intArray := InsertSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func Test_ShellSort(t *testing.T) {
	intArray := ShellSort([]int{0,8,5,1,3,2})
	t.Log(intArray)
}
func Test_HeapSort(t *testing.T) {
	intArray := HeapSort([]int{8, 5, 0, 3, 7, 1, 2})
	t.Log(intArray)
}

func TestFindMaxSeqSum(t *testing.T) {
	t.Log(FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19}))
}

func TestMergeSort(t *testing.T) {
	t.Log(MergeSort([]int{1, 3, -9, 6, 8, -19, 20, -20}))
}
func TestMergerSortNotRecursion(t *testing.T) {
	t.Log(MergerSortNotRecursion([]int{1, 3, -9, 6, 8, -19, 20, -20}))
}
func TestQuickSortNotRecursion(t *testing.T) {
	t.Log(QuickSortNotRecursion([]int{1, 3, -9, 6, 8, -19, 20, -20}))
}
