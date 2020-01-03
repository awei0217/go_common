package hystrix

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func init() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		//超时实践
		Timeout: 5000,
		//最大并发数
		MaxConcurrentRequests: 1,
		//错误百分比,错误率达到这个数值开启熔断
		ErrorPercentThreshold: 25,
		//当熔断器被打开后，SleepWindow的时间就是控制过多久后去尝试服务是否可用了(毫秒)
		SleepWindow: 10,
		//最小请求数，只有到达这个数量后才判断是否开启熔断
		RequestVolumeThreshold: 10,
	})
}

//异步调用
func HystrixAsyStudy() {
	result := make(chan string, 1)
	//定义依赖外部系统的函数
	f1 := func() error {
		// 处理业务系统（调用外部服务）
		fmt.Println("处理业务逻辑")
		result <- "处理结果"
		return nil
	}
	//回调函数，只有 err不为空，才会执行回调函数(如果发生了超时，超时之后也会回调)
	fallBack1 := func(err error) error {
		fmt.Println("回调函数")
		return err
	}
	errors := hystrix.Go("my_command", f1, fallBack1)
	select {
	case r := <-result:
		fmt.Println(r)
	case e := <-errors:
		fmt.Println(e)
	}
}

//同步调用
func HystrixSynStudy() {
	result := make(chan string, 1)
	//定义依赖外部系统的函数
	f1 := func() error {
		// 处理业务系统（调用外部服务）
		fmt.Println("处理业务逻辑")
		time.Sleep(3 * time.Second)
		result <- "处理结果"
		return nil
	}
	//回调函数，只有 err不为空，才会执行回调函数(如果发生了超时，超时之后也会回调)
	fallBack1 := func(err error) error {
		fmt.Println("回调函数")
		return err
	}
	errors := hystrix.Do("my_command", f1, fallBack1)

	select {
	case r := <-result:
		fmt.Println(r)
		fmt.Println(errors)

	}
}

func HystrixCurrentStudy() {
	for i := 0; i < 1000; i++ {
		out := make(chan string, 1)
		err := hystrix.Go("my_command", func() error {
			out <- "结果" + strconv.Itoa(i)
			return nil
		}, func(e error) error {
			fmt.Println("回调")
			return e
		})
		select {
		case r := <-out:
			fmt.Println(r)
		case e := <-err:
			fmt.Println(e)
		}
	}
	time.Sleep(100 * time.Second)
}

func HystrixRongDuanStudy() {

	for i := 0; i < 1000; i++ {
		out := make(chan string, 1)
		err := hystrix.Go("my_command", func() error {
			time.Sleep(50 * time.Millisecond)
			if i%2 == 0 {
				return errors.New("错误了")
			}
			out <- "结果" + strconv.Itoa(i)
			return nil
		}, nil)
		select {
		case r := <-out:
			fmt.Println(r)
		case e := <-err:
			fmt.Println(e)
		}
	}
	time.Sleep(100 * time.Second)
}
