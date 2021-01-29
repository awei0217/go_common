package gout

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
)

func StudyGoutGet() {

	result := ""
	url := "https://www.baidu.com"
	// 发送GET方法
	gout.GET(url).BindBody(&result).Do()

	fmt.Println(result)
}

func StudyGoutGetAndCallBack() {

	result := ""
	url := "https://www.baidu.com"
	// 发送GET方法
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
