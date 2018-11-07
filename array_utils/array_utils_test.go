package array_utils

import "testing"

func TestFindMaxSeqSum(t *testing.T) {
	t.Log(FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19}))
}

func TestBinaryFindOrderArray(t *testing.T) {
	t.Log(BinaryFindOrderArray([]int{1, 2, 3, 4, 5, 6, 6, 7}, 0))
}

func TestBinaryFindFirstOrderArray(t *testing.T) {
	t.Log(BinaryFindFirstOrderArray([]int{1, 2, 3, 4, 6, 6, 6, 7}, 6))
}

func TestBinaryFindTailOrderArray(t *testing.T) {
	t.Log(BinaryFindTailOrderArray([]int{1, 2, 3, 4, 6, 6, 8, 8}, 6))
}

func TestMergeTwoArray(t *testing.T) {
	MergeTwoArray([]int{1},1,[]int{},0)
}
