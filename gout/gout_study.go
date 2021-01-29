package gout

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"time"
)

var url = "https://www.baidu.com"

//GET请求发送
func StudyGoutGet() {

	result := ""
	// 发送GET方法
	gout.GET(url).BindBody(&result).Do()

	fmt.Println(result)
}

//请求回调处理
func StudyGoutCallBack() {

	result := ""
	// 发送GET方法，对返回的结果进行回调处理，可以判断http的code码做出不同的处理
	err := gout.GET(url).Callback(func(context *dataflow.Context) error {

		switch context.Code {

		case 200:
			context.BindBody(&result)
		case 404:
			context.BindBody(&result)

		}

		return nil
	}).Do()

	fmt.Println(err, result)
}

//超时设置
func StudyGoutTimeOut() {

	result := ""
	//设置1毫秒的超时，模拟超时
	err := gout.GET(url).SetTimeout(time.Millisecond * 1).BindBody(&result).Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

//debug设置
func StudyGoutGetDebug() {
	result := ""
	// 发送GET方法
	gout.GET(url).Debug(true).BindBody(&result).Do()

	fmt.Println(result)
}

//跟踪请求耗时时间
func StudyGoutGetTraceInfo() {
	result := ""
	// 发送GET方法
	gout.GET(url).Debug(gout.Trace()).BindBody(&result).Do()

	fmt.Println(result)
}

// 指定并发数的情况下持续一段时间的压测
func StudyGoutBenchmarkDuration() {
	err := gout.
		POST(url).                   //压测本机8080端口
		Filter().                    //打开过滤器
		Bench().                     //选择bench功能
		Concurrent(20).              //并发数
		Durations(10 * time.Second). //压测时间
		Do()

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
