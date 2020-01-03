package go_thinking

import (
	"fmt"
)

// golang的思想求斐波那契数
// 1，1，2，3，5，8，13，21、、、、
func Fibonacci(n int) {

	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < n; i++ {
			num := <-c //这里读是因为防止select那里会阻塞
			fmt.Println(num)
		}
		//可以停止
		quit <- true
	}()

	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println(flag)
			return
		}
	}
	close(c)
	close(quit)
}

// 1，1，2，3，5，8，13，21、、、、

func FibonacciNew(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	x, y := 1, 1
	for i := 3; i <= n; i++ {
		temp := x
		x = y
		y = y + temp
	}
	return y
}
