package algorithm

import "testing"

func TestBitSet_Set(t *testing.T) {
	bitmap := NewBitMap(4)
	bitmap.Set(1)
	bitmap.Set(2)
	bitmap.Set(64)
	t.Log(bitmap.Get(1))
	t.Log(bitmap.Get(2))
	t.Log(bitmap.Get(3))
}
