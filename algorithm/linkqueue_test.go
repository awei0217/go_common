package algorithm

import (
	"testing"
	"time"
)



func TestLinkQueue_Add(t *testing.T) {
	queue := &LinkQueue{nil,nil,0}
	start := time.Now().Unix()
	for i:=0;i<100000000;i++{
		queue.Add(i)
	}
	end := time.Now().Unix()
	t.Log(start)
	t.Log(end)
	result := end-start
	t.Log("添加1000万数据耗时：",result," 秒")

	getStart := time.Now().Unix()
	for i:=0;i<100000000;i++{
		queue.Take()
	}
	getEnd := time.Now().Unix()
	t.Log(getStart)
	t.Log(getEnd)
	getResult := getEnd-getStart
	t.Log("取出1000万数据耗时：",getResult," 秒")
}

func BenchmarkLinkQueue_Add(b *testing.B) {
	queue := &LinkQueue{nil,nil,0}
	for i:=0;i<b.N;i++{
		queue.Add(i)
	}

}