package algorithm

import (
	"fmt"
	"testing"
)

func Test_MaoPaoSort(t *testing.T) {
	intArray := MaoPaoSort([]int{1, 3, 2, 5, 0, 7})
	t.Log(intArray)
	// Output:
	// [0,1,2,3,5,7]
}

//http://blog.studygolang.com/2017/10/how-to-test-with-go/  golang 单元测试
func ExampleMaoPaoSort() {
	intArray := MaoPaoSort([]int{1, 3, 2, 5, 0, 7})
	fmt.Println(intArray)
	// Output:
	// [0 1 2 3 5 7]
}

func Test_SelectSort(t *testing.T) {
	intArray := SelectSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func TestQuickSort(t *testing.T) {
	t.Log(QuickSort([]int{9, 243, 1, 7, 9, 3, 6, 0}))
}

func Test_InsertSort(t *testing.T) {
	intArray := InsertSort([]int{1, 3, 2, 5})
	t.Log(intArray)
}

func Test_ShellSort(t *testing.T) {
	intArray := ShellSort([]int{0, 8, 5, 1, 3, 2})
	t.Log(intArray)
}
func Test_SmallHeapSort(t *testing.T) {
	intArray := SmallHeapSort([]int{1, 2})
	t.Log(intArray)
}

func Test_BigHeapSort(t *testing.T) {
	intArray := BigHeapSort([]int{-10, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 8, 5, 5, 6, 6, 6, 7, 7, 8, 7, 9, 10, 10, 11})
	t.Log(intArray)
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

func TestBucketSort(t *testing.T) {
	t.Log(BucketSort([]int{1, 4, 2, 5, 9, 2, 4, 5, 7, 8, 1}))
}

func TestStringMergerSort(t *testing.T) {
	d := StringMergerSort([]string{"asd", "ssssw", "s", "mmwdsasdw", "sdas"})
	t.Log(d)
}
