package basic

import (
	"fmt"
	"sync"
	"time"
)

func TimeStudy() {
	// 获取当前时间的对象
	t := time.Now()
	// 获取当前的时间戳 秒
	fmt.Println(t.Unix())
	// 获取当前时间戳 纳秒,没有毫秒的方法 1秒 = 1000 毫秒 = 1000 * 1000微妙  = 1000 * 1000 *1000纳秒
	fmt.Println(t.UnixNano())

	fmt.Println(t.Year())       // 年 2018
	fmt.Println(t.Month())      // 月 April
	fmt.Println(t.Day())        // 日  22
	fmt.Println(t.Hour())       // 小时 20
	fmt.Println(t.Minute())     // 分钟 20
	fmt.Println(t.Second())     // 秒 20
	fmt.Println(t.Nanosecond()) // 纳秒 20

	fmt.Println(t.Date())                                                   // 结果 ： 2018 April 22
	fmt.Println(t.Local())                                                  //  本地时间：2018-04-22 20:19:33.0960854 +0800 CST
	fmt.Println(t.Format("2006-01-02 15:04:05"))                            // 必须使用这个固定的字符串格式化日期，时间
	fmt.Println(time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")) // 必须使用这个固定的字符串格式化日期，时间

	var wg sync.WaitGroup
	wg.Add(2)
	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)
	wg.Wait()
}
