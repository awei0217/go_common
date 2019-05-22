package benchmark

import (
	"fmt"
	"runtime"
	"time"
)

func init() {

	runtime.GOMAXPROCS(8)

}

func Fib(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

type I interface {
	f()
}

type S struct {
	a int64
}

var a int

func f() {
	a++
}

func Method() {
	start := time.Now().UnixNano()
	for j := 0; j < 1; j++ {
		f()
	}
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}
