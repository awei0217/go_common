package algorithm

import "testing"

func TestArrayLoopQueue_Add(t *testing.T) {
	alq := &ArrayLoopQueue{}
	for i:=0;i<20;i++{
		t.Log(alq.Add(i))
	}
	for i:=0;i<20;i++{
		t.Log(alq.Take())
	}
}
