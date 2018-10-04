package algorithm

import (
	"strconv"
	"testing"
)

func TestLRUArray_Put(t *testing.T) {
	lru := LRUArray{}

	for j :=1;j<=11;j++{
		lru.Put(strconv.Itoa(j),strconv.Itoa(j))
	}
	t.Log(lru.keys)
	t.Log(lru.values)
	t.Log(lru.Get("5"))
	t.Log(lru.keys)
	t.Log(lru.values)
	t.Log(lru.Get("5"))
	t.Log(lru.keys)
	t.Log(lru.values)
	t.Log(lru.maxIndex)
}
