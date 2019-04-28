// 用go思维实现求1-100的质数
package go_thinking

import "testing"

func TestSieve(t *testing.T) {
	Sieve()
}

//
func BenchmarkSieve(b *testing.B) {
	for i := 0; i < 1000; i++ {
		Sieve()
	}
}
