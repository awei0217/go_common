package algorithm

import (
	"testing"
)


func TestRedFileSortWriteTempFile(t *testing.T) {
	RedFileSortWriteTempFile("E://number.txt")
}

func TestQuickSort2(t *testing.T) {
	t.Log(QuickSort2([]int{15, 3, 5, 2, 8, 0, 19, 21, 7, 65, 6, 76, 12, 17}))
}

func TestCreateSourceFile(t *testing.T) {
	CreateSourceFile()
}