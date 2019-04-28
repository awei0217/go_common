// 用go思维实现求1-100的质数
package go_thinking

import "fmt"

func generate(ch chan<- int) {
	for i := 2; i < 10000; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
	close(ch)
}

func filter(src <-chan int, dst chan<- int) {
	prime, ok := <-src
	if !ok {
		close(dst)
		return
	}
	fmt.Println(prime)
	out := make(chan int)
	go filter(out, dst)

	for num := range src {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out)
}

//筛选素数
func Sieve() {

	origin, wait := make(chan int), make(chan int) // Create a new channel.

	go generate(origin)

	go filter(origin, wait)

	<-wait

}
