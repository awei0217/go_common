package basic

import (
	"testing"
)

/**
执行顺序，先声明的后执行
*/
func TestDeferStudy(t *testing.T) {

	DeferStudy()
}

func TestRecoverStudy(t *testing.T) {
	t.Log(RecoverStudy())
}

func TestRecoverA(t *testing.T) {
	l, err := RecoverA()
	t.Log("返回结果", l, err)
}

func TestA(t *testing.T) {
	t.Log(A())
	//t.Log(B())
}

//	BenchmarkNoDefer-8   	13953618	        79.8 ns/op
//	BenchmarkDefer-8   	    13793086	        81.7 ns/op

//  BenchmarkNoDefer-4  　　15759076   　　　   76.8 ns/op
//  BenchmarkDefer-4     　 11046517    　 　   105.4 ns/op
func BenchmarkNoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoDefer()
	}
}
func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Defer()
	}
}
