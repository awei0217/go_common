package basic

import (
	"fmt"
	"runtime"
)

func StudyRuntime() {

	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("archive:", runtime.GOOS)

	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 2 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}

func init() {
	runtime.GOMAXPROCS(8) //使用单核
}
