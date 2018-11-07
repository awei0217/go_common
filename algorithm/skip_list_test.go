package algorithm

import (
	"testing"
)

func TestNewSkipList(t *testing.T) {
	//创建一个32层的跳跃列表
	skip := NewSkipList(32)
	for i:=0;i<10000000 ;i++  {
		skip.insert(uint64(i),i)
	}
}

func BenchmarkNewSkipList(b *testing.B) {
	//创建一个32层的跳跃列表
	skip := NewSkipList(32)
	for i:=0;i<1000000 ;i++  {
		skip.insert(uint64(i),i)
	}
}
