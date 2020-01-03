// 用go思维实现求1-100的质数
package go_thinking

import "fmt"

func generate(ch chan<- int) {
	for i := 2; i < 100; i++ {
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

//1-100的质数
func SieveNew(n int) []int {

	result := make([]int, 0)
	for i := 2; i <= n; i++ {
		if i == 2 {
			result = append(result, i)
			continue
		}
		count := 0
		for j := 1; j <= i; j++ {
			if count > 2 {
				break
			}
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			result = append(result, i)
		}
	}
	return result

}
