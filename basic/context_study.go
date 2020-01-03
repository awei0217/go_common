package basic

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextWithTimeOutOne() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//开启新的协程
	go func(ctx context.Context) {
		//模拟处理任务
		time.Sleep(10 * time.Second)
		fmt.Println("任务处理完成")
		cancel()
	}(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("context timeout")
	}
}

/**
超时控制，如果父context超时结束，子context超时也结束 ，哪怕父context超时设置1秒，子context超时设置3秒
*/
func ContextWithTimeOut() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	//开启新的协程
	go func(ctx context.Context, wg *sync.WaitGroup) {
		ctxA, _ := context.WithTimeout(ctx, 5*time.Second)
		ch := make(chan int)
		//开起新的协程
		go func(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
			//模拟读取数据
			sumCh := make(chan int)
			go func(sumCh chan int) {
				sum := 10
				time.Sleep(3 * time.Second)
				sumCh <- sum
			}(sumCh)

			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("two done")
			case <-sumCh:
				wg.Done()
			}
		}(ctxA, ch, wg)

		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("one done")
		case <-ch:
			wg.Done()
		}
	}(ctx, wg)
	wg.Wait()
}

/**
手动编码取消，研发人员可控制的取消
*/
func ContextWithCancel() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case c := <-ctx.Done():
			{
				fmt.Println(c)
				wg.Done()
			}
		}
	}(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
}

/**
可延迟一定的时间取消，也可以手动编码取消
*/
func ContextWithDeadLine() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))

	go func(wg *sync.WaitGroup, ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("ctx deadline done")
			wg.Done()
		}
	}(wg, ctx)
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}

/**
上下文key value,根据key获取，当子context 获取不到对应key时，就取父的context获取
*/
func ContextWithValue() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx := context.WithValue(nil, "key", "value")
	go func(ctx context.Context) {
		ctxB := context.WithValue(ctx, "key1", "value1")
		v := ctx.Value("key")
		v1 := ctxB.Value("key1")
		v2 := ctxB.Value("key")
		fmt.Println(v, v1, v2)
		wg.Done()
	}(ctx)
	wg.Wait()
}
