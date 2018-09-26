package algorithm

import "testing"

func TestQuickSortTopK(t *testing.T) {
	t.Log(QuickSortTopK([]int{15,3,98,2,8,87,19,21,7,65,6,76,12,100,99},5))
}
func TestSelectSortTopK(t *testing.T) {
	t.Log(SelectSortTopK([]int{15,3,5,2,8,0,19,21,7,65,6,76,12,17},5))
}
func TestHeadTopK(t *testing.T) {
	t.Log(HeadTopK([]int{1,3,5,2,18,0,19,21,7,65,6,76,12,17},5))
}
