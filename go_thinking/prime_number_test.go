// 用go思维实现求1-100的质数

//质数又称素数。指整数在一个大于1的自然数中，除了1和此整数自身外，没法被其他自然数整除的数
package go_thinking

import (
	"fmt"
	"testing"
)

func TestSieve(t *testing.T) {
	Sieve()
}

func BenchmarkSieve(b *testing.B) {
	for i := 1; i < 100; i++ {
		Sieve()
	}
}

func TestSieveNew(t *testing.T) {
	fmt.Println(SieveNew(100))
}
