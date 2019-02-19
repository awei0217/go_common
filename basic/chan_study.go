package basic

import (
	"fmt"
	"sync"
	"time"
)

/**
goroutines are asleep - deadlock 会发生死锁
原因是 主程序在等待一个永远也不可能等待到的值，goroutine 会判断为死锁
*/
func ChanStudy1() {
	cint := make(chan int)
	fmt.Println(<-cint) // 从通道中取出一个值
}

/**
改进：将等待放到一个单独的goroutines中，这样就不会阻塞主程序
*/
func ChanStudy1_1() {
	cint := make(chan int)
	go func() {
		fmt.Println(<-cint)
	}()
	time.Sleep(3 * time.Second)
}

/**
goroutines are asleep - deadlock!
发生死锁的原因是往一个没有缓冲区的通道中放入元素时，会阻塞，因此，主程序被阻塞，golang判断为死锁
*/
func ChanStudy2() {
	cint := make(chan int)
	cint <- 1 //往通道中放入一个数据
}

func ChanStudy2_2() {
	cint := make(chan int)
	go func() {
		cint <- 1
	}()
	time.Sleep(3 * time.Second)
}

func ReadAndWriteChanStudy() {
	//创建一个带有1个缓冲区的chan
	cint := make(chan int, 1)
	//写入
	cint <- 1
	//读取 ok代表是否读取到数据
	temp, ok := <-cint
	if ok {
		fmt.Println(temp, ok) // 1，true
		//关闭通道
		close(cint)
		temp, ok = <-cint
		fmt.Println(temp, ok) //对于关闭的通道，ok返回false  0，false  如果closed后，chan有数据，ok还是true的，直到chan没有数据了才false
		//向已经关闭的通道写入数据会报错
		//cint <- 1  //panic: send on closed channel
	}
}

func SelectStudy() {

	cint := make(chan int, 0)
	for i := 0; i < 10; i++ {
		//采用select 往已经满了通道中写入数据时会丢弃，如下，只会写入0，其它的都会丢弃
		// select 往通道写入时不会阻塞
		select {
		case cint <- i:
		default:
			fmt.Println("丢弃了", i)
		}
	}
}

/**
生产者消费者的例子
*/
func ProductAndConsumer() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cint := make(chan int, 10)

	go func() {
		//product
		for i := 0; i < 1000; i++ {
			cint <- i
		}
		close(cint)
	}()
	go func() {
		defer wg.Done()
		//consumer
		for temp := range cint {
			fmt.Println(temp)
		}
	}()
	wg.Wait()
}
