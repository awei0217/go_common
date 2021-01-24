package charts

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"os"
)

//项目网址
//https://go-echarts.github.io/go-echarts/docs/quickstart

var nameItems = []string{"现状", "目标"}

func randInt() []float64 {
	r := make([]float64, 0)

	for i := 0; i < len(nameItems); i++ {
		if i == 0 {
			r = append(r, float64(500.0*8/3600))
		}
		if i == 1 {
			r = append(r, float64(500.0*3/3600))
		}
	}
	fmt.Println(r)
	return r
}

func BarChartStudy() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("入仓时间", randInt())
	//生成一个html网页
	f, err := os.Create("./bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}
