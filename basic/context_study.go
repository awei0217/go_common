package basic

import (
	"context"
	"fmt"
	"time"
)

func ContextStudy() {

	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go testA(ctx)
	select {}

}

func testA(ctx context.Context) {
	ctxA, _ := context.WithTimeout(ctx, 5*time.Second)
	ch := make(chan int)
	go testB(ctxA, ch)
	select {
	case <-ctx.Done():
		fmt.Println("testA done")
		return
	case i := <-ch:
		fmt.Println(i)
	}
}

func testB(ctx context.Context, ch chan int) {
	//模拟读取数据
	sumCh := make(chan int)
	go func(sumCh chan int) {
		sum := 10
		time.Sleep(15 * time.Second)
		sumCh <- sum
	}(sumCh)

	select {
	case s := <-ctx.Done():
		fmt.Println(s)
		fmt.Println(ctx.Err().Error())
		return
	case i := <-sumCh:
		fmt.Println("send ", i)
		ch <- i
	}
}
