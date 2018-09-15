package algorithm

import (
	"testing"

)

func Test_MaoPaoSort1(t *testing.T)  {
	intArray := MaoPaoSort1()
	t.Log(intArray)
}

func Test_MaoPaoSort2(t *testing.T)  {
	intArray := MaoPaoSort2()
	t.Log(intArray)
}

func Test_SelectSort1(t *testing.T)  {
	intArray := SelectSort1();
	t.Log(intArray)
}

func Test_SelectSort2(t *testing.T)  {
	intArray := SelectSort2();
	t.Log(intArray)
}

func Test_InsertSort(t *testing.T)  {
	intArray := InsertSort();
	t.Log(intArray)
}

func Test_ShellSort(t *testing.T){
	intArray := ShellSort()
	t.Log(intArray)
}

func Test_SortIntegers(t *testing.T) {
	var A = []int {0}
	SortIntegers(&A)
}
