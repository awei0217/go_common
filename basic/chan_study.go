package basic

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/**
goroutines are asleep - deadlock 会发生死锁
原因是 主程序在等待一个永远也不可能等待到的值，goroutine 会判断为死锁
*/
func ChanStudy1() {
	cint := make(chan int)
	<-cint

}

/**
改进：将等待放到一个单独的goroutines中，这样就不会阻塞主程序
*/
func ChanStudy1_1() {
	cint := make(chan string)
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

func ChanStudy3() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		c1 <- 1
		c2 <- 2
	}()
	<-c2
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
func queryUserById(id int) chan string {
	c := make(chan string)
	go func() {
		c <- "姓名" + strconv.Itoa(id)
	}()
	return c
}

func main() {
	//三个协程同时并发查询，缩小执行时间，
	//本来一次查询需要1秒，顺序执行就得3秒，
	//现在并发执行总共1秒就执行完成
	name1 := queryUserById(1)
	name2 := queryUserById(1)
	name3 := queryUserById(3)
	//从通道中获取执行结果
	<-name1
	<-name2
	<-name3

	c1, c2, c3 := queryUserById(1), queryUserById(2), queryUserById(3)
	c := make(chan string)
	go func() { // 开一个goroutine监视各个信道数据输出并收集数据到信道c
		for {
			select { // 监视c1, c2, c3的流出，并全部流入信道c
			case v1 := <-c1:
				c <- v1
			case v2 := <-c2:
				c <- v2
			case v3 := <-c3:
				c <- v3
			}
		}
	}()
	// 阻塞主线，取出信道c的数据
	for i := 0; i < 3; i++ {
		// 从打印来看我们的数据输出并不是严格的1,2,3顺序
		fmt.Println(<-c)
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
			fmt.Println("当前通道元素个数", len(cint))
		}
	}()
	wg.Wait()
}

func StudyChannel10() {

	var wg sync.WaitGroup
	wg.Add(1)
	//不带缓冲区的channel
	c := make(chan string)
	go func() {
		defer func() {
			wg.Done()
			close(c)
		}()
		for {
			//从通道中取出数据
			temp := <-c
			fmt.Println(temp)
			if temp == "写入数据3" {
				break
			}
		}
	}()
	//主协程循环往通道写入值
	for i := 1; i < 4; i++ {
		c <- "写入数据" + strconv.Itoa(i)
	}
	//等待新的协程运行完毕，程序才退出
	wg.Wait()
}
