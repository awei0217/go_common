package algorithm

import (
	"bytes"
	"testing"
)

// 字符串拼接的集中方法


func BenchmarkAddStringWithOperator(b *testing.B) {
	for i:=0;i<b.N;i++{
		AddStringWithOperator()
	}
}

func BenchmarkAddStringWidthSprintf(b *testing.B) {
	for i:=0;i<b.N;i++{
		AddStringWidthSprintf()
	}
}

func BenchmarkAddStringWidthJoin(b *testing.B) {
	for i:=0;i<b.N;i++{
		AddStringWidthJoin()
	}
}

func BenchmarkAddStringWidthBuffer(b *testing.B) {
	for i:=0;i<b.N;i++{
		AddStringWidthBuffer()
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}
