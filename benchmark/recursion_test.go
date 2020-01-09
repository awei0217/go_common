package benchmark

import "testing"

func TestFib(t *testing.T) {
	Fib(45)
}

func BenchmarkFib(b *testing.B) {
	Fib(45)
}

func TestMethod(t *testing.T) {
	Method()
}

func BenchmarkMethod2(b *testing.B) {
	Method()
}

func BenchmarkMethod(b *testing.B) {
	Method()
}

func TestSort(t *testing.T) {
	Sort()
}
