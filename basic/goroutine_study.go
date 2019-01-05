package basic

import "fmt"

func GoRoutineStudy() {
	fmt.Println("goroutine test")
	pages1 := make([]int, 50)
	pages2 := make([]int, 50)
	for i := 0; i < cap(pages1); i++ {
		pages1[i] = i * 2
	}
	for i := 0; i < cap(pages2); i++ {
		pages2[i] = i*2 + 1
	}
	pageSlice := make([]interface{}, 2)
	pageSlice[0] = pages1
	pageSlice[1] = pages2
	var channel chan string = make(chan string, 2)
	for i := 0; i < 2; i++ {
		// interface 类型到其他类型的转换 pageSlice[i].([]int)
		go findByPage(pageSlice[i].([]int), channel)
	}
	// 直到协成跑完取到消息
	for {
		message, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(message)
	}
}

func findByPage(pages []int, channel chan string) {
	for i := 0; i < len(pages); i++ {
		fmt.Printf("查询第 %d 页成功\n", pages[i])
	}
	channel <- "查询完毕"
	close(channel)
}
