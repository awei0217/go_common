package algorithm

import (
	"strconv"
	"testing"
)

func TestNewHashMap(t *testing.T) {
	hm := NewHashMap()
	hm.Put("spw", "Sun")
	hm.Put("spw", "Sun")
	hm.Put("spw", "Sun")

	t.Log(hm.Size())
	for i := 0; i < 1000000; i++ {
		hm.Put("spw"+strconv.Itoa(i), "Sun")
	}
	t.Log(hm.Size())
	t.Log("ss")
}
