package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Sort1() {
	const NUM = 100000000
	start := time.Now().UnixNano()
	for i := 0; i < NUM; i++ {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		BubbleSort(arr)
	}
	fmt.Println(time.Now().UnixNano() - start)
}
func Sort() {
	const NUM = 100000000
	var arr []int
	start := time.Now().UnixNano()
	for i := 0; i < 1; i++ {
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		BubbleSort(arr)
	}
	fmt.Println(time.Now().UnixNano() - start)
}

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return "/usr/home"
	}
	return c.path
}

//func main() {
//
//	var c1 *Config
//	var c2 = &Config{
//		path: "/export",
//	}
//	fmt.Println(c1.Path(), c2.Path())
//}

//const (
//	DATETIME = "2019-04-04 12:12:12"
//)

/*func init() {
	fmt.Println("init A")
}*/
//var result string

/*func main() {\
	hello := "hello"
	var world string = "world"
	result = hello+world
	h.HandleFunc("/",IndexHandler)
	h.ListenAndServe(":8080",nil)
}*/

//func IndexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(DATETIME)
//	w.Write([]byte(result))
//}
//func QueryHelloWorld(str string) (string, error) {
//	if len(result) == 0 {
//		return result, errors.New("result is not null")
//	}
//	return result + str, nil
//}
//
//type I interface {
//	f()
//}
//
//type S struct {
//	a int64
//}
//
//func (this *S) f() {
//	this.a++
//}

//func main() {
//	c := make(chan int, 10)
//	close(c)
//	c <- 1
//	start := time.Now().UnixNano() / 1000 / 1000
//	var s S
//	var i I = &s
//	for j := 0; j < 1000000000; j++ {
//		i.f()
//	}
//	end := time.Now().UnixNano() / 1000 / 1000
//	println(end - start)
//}

//func findByPage(pages []int, channel chan string) {
//	for i := 0; i < len(pages); i++ {
//		fmt.Printf("查询第 %d 页成功\n", pages[i])
//	}
//	channel <- "查询完毕"
//	close(channel)
//}
//
//type student struct {
//	Name string
//	Age  int
//}
//
//func pase_student() map[string]*student {
//	m := make(map[string]*student)
//	stus := make([]student, 0)
//
//	stus = append(stus, student{Name: "sss", Age: 25})
//
//	for _, stu := range stus {
//		stu.Name = "123"
//	}
//	fmt.Println(stus)
//	return m
//}

//func Sort1() {
//	//4242922000
//	const NUM = 100000000
//	start := time.Now().UnixNano()
//	for i := 0; i < NUM; i++ {
//		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//		BubbleSort(arr)
//	}
//	fmt.Println(time.Now().UnixNano() - start)
//}

//排序
func BubbleSort(arr []int) {
	//for j := 0; j < len(arr)-1; j++ {
	//	for k := 0; k < len(arr)-1-j; k++ {
	//		if arr[k] < arr[k+1] {
	//			temp := arr[k]
	//			arr[k] = arr[k+1]
	//			arr[k+1] = temp
	//		}
	//	}
	//}
}

func main() {
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "GoTool"
	//程序的用途描述
	oApp.Usage = "To save the world"
	//程序的版本号
	oApp.Version = "1.0.0"

	//设置多个命令处理函数
	oApp.Commands = []*cli.Command{
		{
			//命令全称
			Name: "lang",
			//命令简写
			Aliases: []string{"l"},
			//命令详细描述
			Usage: "Setting language",
			//命令处理函数
			Action: func(c *cli.Context) error {
				// 通过c.Args().First()获取命令行参数
				fmt.Printf("language=%v  \n", c.Args().First())
				return nil
			},
		},
		{
			Name:    "encode",
			Aliases: []string{"e"},
			Usage:   "Setting encoding",
			Action: func(c *cli.Context) error {
				fmt.Printf("encoding=%v \n", c.Args().First())
				return nil
			},
		},
	}

	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
