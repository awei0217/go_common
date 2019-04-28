package benchmark

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
	for j := 0; j < 1000000000; j++ {
		f()
	}
}
